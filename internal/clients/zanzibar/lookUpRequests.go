package zanzibar

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/nrc-no/notcore/internal/utils"
	"log"
)

func (c *ZanzibarClient) CheckGlobalAdminExists(ctx context.Context) (bool, error) {
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
		log.Fatalf("failed to check for existence of global admin: %s, %s", err, resp)
		return false, err
	}

	lookupResp, err := resp.Recv()

	if &lookupResp != nil {
		return true, nil
	}

	return false, err
}

func (c *ZanzibarClient) CheckPermittedLocations(ctx context.Context, locationType LocationType) (bool, error) {
	r := &pb.LookupResourcesRequest{
		ResourceObjectType: c.prefix + "/" + locationType.String(),
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: c.prefix + "/user",
				ObjectId:   utils.GetRequestUser(ctx).ID,
			},
		},
		Permission: "view",
	}

	resp, err := c.z.LookupResources(ctx, r)

	if err != nil {
		log.Fatalf("failed to lookup location permissions: %s, %s", err, resp)
		return false, err
	}

	lookupResp, err := resp.Recv()

	if &lookupResp != nil {
		return true, nil
	}

	return false, err
}
func (c *ZanzibarClient) IsAnyAdmin(ctx context.Context, locationType LocationType, userId string) (bool, error) {
	r := &pb.LookupResourcesRequest{
		ResourceObjectType: c.prefix + "/" + locationType.String(),
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: c.prefix + "/user",
				ObjectId:   userId,
			},
		},
		Permission: "admin",
	}

	resp, err := c.z.LookupResources(ctx, r)

	if err != nil {
		log.Fatalf("failed to lookup location permissions: %s, %s", err, resp)
		return false, err
	}

	lookupResp, err := resp.Recv()

	if &lookupResp != nil {
		return true, nil
	}

	return false, err
}

func (c *ZanzibarClient) GetCountryViewPermissions(ctx context.Context, userId string) (bool, error) {
	r := &pb.LookupResourcesRequest{
		ResourceObjectType: c.prefix + "/country",
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: c.prefix + "/user",
				ObjectId:   userId,
			},
		},
		Permission: "view_co_individuals",
	}

	resp, err := c.z.LookupResources(ctx, r)
	if err != nil {
		log.Fatalf("failed to lookup location permissions: %s, %s", err, resp)
		return false, err
	}

	lookupResp, err := resp.Recv()

	if &lookupResp != nil {
		return true, nil
	}

	return false, err
}

func (c *ZanzibarClient) CheckUserIsStaffAtLocation(ctx context.Context, location LocationType) (bool, error) {
	r := &pb.LookupResourcesRequest{
		ResourceObjectType: c.prefix + "/" + location.String(),
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: c.prefix + "/user",
				ObjectId:   utils.GetRequestUser(ctx).ID,
			},
		},
		Permission: "staff",
	}

	resp, err := c.z.LookupResources(ctx, r)
	if err != nil {
		log.Fatalf("failed to lookup location permissions: %s, %s", err, resp)
		return false, err
	}

	lookupResp, err := resp.Recv()

	if &lookupResp != nil {
		return true, nil
	}

	return false, err
}
