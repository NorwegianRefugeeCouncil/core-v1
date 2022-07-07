package zanzibar

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/nrc-no/notcore/internal/utils"
	"log"
)

func (c *ZanzibarClient) AddCountry(ctx context.Context, countryCode string) (*pb.WriteRelationshipsResponse, error) {
	r := &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Relationship: &pb.Relationship{
					Relation: "region",
					Resource: &pb.ObjectReference{
						ObjectType: c.prefix + "/global",
						ObjectId:   "nrc",
					},
					Subject: &pb.SubjectReference{
						Object: &pb.ObjectReference{
							ObjectType: c.prefix + "/region",
							ObjectId:   "default_region",
						},
					},
				},
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			},
			{
				Relationship: &pb.Relationship{
					Relation: "country",
					Resource: &pb.ObjectReference{
						ObjectType: c.prefix + "/region",
						ObjectId:   "default_region",
					},
					Subject: &pb.SubjectReference{
						Object: &pb.ObjectReference{
							ObjectType: c.prefix + "/country",
							ObjectId:   countryCode,
						},
					},
				},
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			},
		},
	}

	resp, err := c.z.WriteRelationships(ctx, r)

	if err != nil {
		log.Fatalf("failed to add country to zanzibar graph: %s, %s", err, resp)
		return nil, err
	}
	return resp, nil
}

func (c *ZanzibarClient) AddUserToLocation(ctx context.Context, location LocationType, locationCode string, userId string) (*pb.WriteRelationshipsResponse, error) {
	r := &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Relationship: &pb.Relationship{
					Relation: "staff",
					Resource: &pb.ObjectReference{
						ObjectType: c.prefix + "/" + location.String(),
						ObjectId:   locationCode,
					},
					Subject: &pb.SubjectReference{
						Object: &pb.ObjectReference{
							ObjectType: c.prefix + "/user",
							ObjectId:   userId,
						},
					},
				},
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			},
		},
	}

	resp, err := c.z.WriteRelationships(ctx, r)

	if err != nil {
		log.Fatalf("failed to add user to country staff: %s, %s", err, resp)
		return nil, err
	}
	return resp, nil
}

func (c *ZanzibarClient) AddAdminToLocation(ctx context.Context, location LocationType, locationCode string, userId string) (*pb.WriteRelationshipsResponse, error) {
	r := &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Relationship: &pb.Relationship{
					Relation: "admin",
					Resource: &pb.ObjectReference{
						ObjectType: c.prefix + "/" + location.String(),
						ObjectId:   locationCode,
					},
					Subject: &pb.SubjectReference{
						Object: &pb.ObjectReference{
							ObjectType: c.prefix + "/user",
							ObjectId:   userId,
						},
					},
				},
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			},
		},
	}

	resp, err := c.z.WriteRelationships(ctx, r)

	if err != nil {
		log.Fatalf("failed to add admin to location: %s, %s", err, resp)
		return nil, err
	}
	return resp, nil
}

func (c *ZanzibarClient) AddGlobalAdmin(ctx context.Context) (*pb.WriteRelationshipsResponse, error) {
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
		log.Fatalf("failed to create global admin: %s, %s", err, resp)
		return nil, err
	}
	return resp, nil
}

func (c *ZanzibarClient) AddCountryAdmin(ctx context.Context, countryCode string) (*pb.WriteRelationshipsResponse, error) {
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
		log.Fatalf("failed to create country admin: %s, %s", err, resp)
		return nil, err
	}
	return resp, nil
}
