package zanzibar

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/nrc-no/notcore/internal/utils"
	"log"
)

type RelationClient interface {
	AddGlobalAdmin(ctx context.Context) (*pb.WriteRelationshipsResponse, error)
	AddCountryAdmin(ctx context.Context, countryCode string) (*pb.WriteRelationshipsResponse, error)
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
