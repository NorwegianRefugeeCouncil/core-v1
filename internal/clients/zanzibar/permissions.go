package zanzibar

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/nrc-no/notcore/internal/utils"
	"log"
)

func (c *ZanzibarClient) CheckIsGlobalAdmin(ctx context.Context) (bool, error) {
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

func (c *ZanzibarClient) GetCountryViewPermissions(ctx context.Context) (*pb.ExpandPermissionTreeResponse, error) {
	r := &pb.ExpandPermissionTreeRequest{
		Resource: &pb.ObjectReference{
			ObjectType: c.prefix + "/user",
			ObjectId:   utils.GetRequestUser(ctx).ID,
		},
		Permission: "view_co_individuals",
	}

	resp, err := c.z.ExpandPermissionTree(ctx, r)

	if err != nil {
		log.Fatalf("failed to get country view permissions: %s, %s", err, resp)
		return resp, err
	}

	tree := resp.TreeRoot
	log.Fatalf(tree.String())

	return nil, err
}

func (c *ZanzibarClient) CheckforUser(ctx context.Context, location LocationType) (bool, error) {

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
