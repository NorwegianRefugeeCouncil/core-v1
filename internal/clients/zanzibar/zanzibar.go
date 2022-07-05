package zanzibar

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"github.com/nrc-no/notcore/cmd/devinit"
	"log"
)

type Client struct {
	Relation    RelationClient
	Nodes       NodeClient
	Permissions PermissionClient
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
