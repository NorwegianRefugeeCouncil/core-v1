package zanzibar

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/nrc-no/notcore/internal/utils"
	"log"
)

type NodeClient interface {
	AddCountry(ctx context.Context, countryCode string) (*pb.WriteRelationshipsResponse, error)
	AddUserToCountry(ctx context.Context, countryCode string) (*pb.WriteRelationshipsResponse, error)
}

func (c *zanzibarClient) AddCountry(ctx context.Context, countryCode string) (*pb.WriteRelationshipsResponse, error) {
	r := &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Relationship: &pb.Relationship{
					Relation: "region",
					Resource: &pb.ObjectReference{
						ObjectType: c.prefix + "/global",
						ObjectId:   countryCode,
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
		log.Fatalf("failed to create relationship between database and creator: %s, %s", err, resp)
		return nil, err
	}
	return resp, nil
}

func (c *zanzibarClient) AddUserToCountry(ctx context.Context, countryCode string) (*pb.WriteRelationshipsResponse, error) {
	r := &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Relationship: &pb.Relationship{
					Relation: "staff",
					Resource: &pb.ObjectReference{
						ObjectType: c.prefix + "/" + countryCode,
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
