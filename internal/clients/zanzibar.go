package zanzibar

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"github.com/nrc-no/notcore/cmd/devinit"
	"github.com/nrc-no/notcore/internal/utils"
	"log"
)

const schema = `
definition notcore/user {}

definition notcore/global { // org = nrc?
  relation admin: notcore/user
  permission write_all_individuals = admin
  permission view_all_individuals = admin + write_all_individuals
}

definition notcore/region {
  relation global: notcore/global
  relation ro_admin: notcore/user
  permission write_ro_individuals = ro_admin + global->write_all_individuals
  permission view_ro_individuals = ro_admin + global->view_all_individuals + write_ro_individuals
}

definition notcore/country {
  relation region: notcore/region
  relation co_admin: notcore/user
  permission write_co_individuals = co_admin + region->write_ro_individuals
  permission view_co_individuals = co_admin + region->view_ro_individuals + write_co_individuals
}

definition notcore/individual {
  relation ind_org: notcore/global
  relation ind_region: notcore/region
  relation ind_country: notcore/country

  relation writer: notcore/user
  relation reader: notcore/user

  permission view = reader + writer + ind_country->view_co_individuals
}
`

type Client interface {
	WriteCountryAdmin(ctx context.Context, countryCode string) (*pb.WriteRelationshipsResponse, error)
	WriteGlobalAdmin(ctx context.Context) (*pb.WriteRelationshipsResponse, error)
	CheckGlobalAdmin(ctx context.Context) (bool, error)
	CheckAnyGlobalAdmin(ctx context.Context) (bool, error)
}

func NewZanzibarClient(c devinit.Config) *zanzibarClient {
	client, err := authzed.NewClient(
		//"localhost:50051",
		//grpcutil.WithInsecureBearerToken(c.SpiceDBToken),
		"grpc.authzed.com:443",
		grpcutil.WithBearerToken(c.ZanzibarToken),
		grpcutil.WithSystemCerts(grpcutil.VerifyCA),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	request := &pb.WriteSchemaRequest{Schema: schema}
	_, err = client.WriteSchema(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
	}

	return &zanzibarClient{
		z:      client,
		prefix: c.ZanzibarPrefix,
	}
}

type zanzibarClient struct {
	z      *authzed.Client
	prefix string
}

func (c *zanzibarClient) WriteGlobalAdmin(ctx context.Context) (*pb.WriteRelationshipsResponse, error) {
	r := &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Relationship: &pb.Relationship{
					Relation: "admin",
					Resource: &pb.ObjectReference{
						ObjectType: c.prefix + "/global",
						ObjectId:   "nrc",
					},
					Subject: &pb.SubjectReference{
						Object: &pb.ObjectReference{
							ObjectType: c.prefix + "/user",
							ObjectId:   utils.GetRequestUser(ctx).ID,
						},
					},
				},
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			},
		},
	}

	resp, err := c.z.WriteRelationships(ctx, r)

	if err != nil {
		log.Fatalf("failed to create relationship between database and creator: %s, %s", err, resp)
		return nil, err
	}
	return resp, nil
}

func (c *zanzibarClient) WriteCountryAdmin(ctx context.Context, countryCode string) (*pb.WriteRelationshipsResponse, error) {
	r := &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Relationship: &pb.Relationship{
					Relation: "co_admin",
					Resource: &pb.ObjectReference{
						ObjectType: c.prefix + "/country",
						ObjectId:   countryCode,
					},
					Subject: &pb.SubjectReference{
						Object: &pb.ObjectReference{
							ObjectType: c.prefix + "/user",
							ObjectId:   utils.GetRequestUser(ctx).ID,
						},
					},
				},
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			},
		},
	}

	resp, err := c.z.WriteRelationships(ctx, r)

	if err != nil {
		log.Fatalf("failed to create relationship between database and creator: %s, %s", err, resp)
		return nil, err
	}
	return resp, nil
}

func (c *zanzibarClient) CheckGlobalAdmin(ctx context.Context) (bool, error) {
	r := &pb.CheckPermissionRequest{
		Resource: &pb.ObjectReference{
			ObjectType: c.prefix + "/global",
			ObjectId:   "nrc",
		},
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: c.prefix + "/user",
				ObjectId:   utils.GetRequestUser(ctx).ID,
			},
		},
		Permission: "admin",
	}

	resp, err := c.z.CheckPermission(ctx, r)

	if err != nil {
		log.Fatalf("failed to check for global admin: %s, %s", err, resp)
		return false, err
	}

	if resp.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
		return true, nil
	}

	return false, err
}

func (c *zanzibarClient) CheckAnyGlobalAdmin(ctx context.Context) (bool, error) {
	r := &pb.LookupResourcesRequest{
		ResourceObjectType: c.prefix + "/user",
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: c.prefix + "/global",
				ObjectId:   "nrc",
			},
		},
		Permission: "is_admin_of",
	}

	resp, err := c.z.LookupResources(ctx, r)

	if err != nil {
		log.Fatalf("failed to look up global admin: %s, %s", err, resp)
		return false, err
	}

	lookupResp, err := resp.Recv()

	if &lookupResp != nil {
		//if &lookupResp.ResourceObjectId != nil {
		return true, nil
		//}
		//return false, err
	}

	return false, err
}
