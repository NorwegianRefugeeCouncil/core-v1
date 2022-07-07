package zanzibar

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"log"
)

func (c *ZanzibarClient) CheckIsGlobalAdmin(ctx context.Context, userId string) (bool, error) {
	r := &pb.CheckPermissionRequest{
		Resource: &pb.ObjectReference{
			ObjectType: c.prefix + "/global",
			ObjectId:   "nrc",
		},
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: c.prefix + "/user",
				ObjectId:   userId,
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

func (c *ZanzibarClient) CheckUserCanViewCountry(ctx context.Context, userId string, countryCode string) (bool, error) {
	r := &pb.CheckPermissionRequest{
		Resource: &pb.ObjectReference{
			ObjectType: c.prefix + "/country",
			ObjectId:   countryCode,
		},
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: c.prefix + "/user",
				ObjectId:   userId,
			},
		},
		Permission: "view_co_individuals",
	}

	resp, err := c.z.CheckPermission(ctx, r)

	if err != nil {
		log.Fatalf("failed to get country view permissions: %s, %s", err, resp)
		return false, err
	}

	canSee := resp.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION

	return canSee, err
}
