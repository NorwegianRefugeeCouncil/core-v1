package zanzibar

import (
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"github.com/nrc-no/notcore/cmd/devinit"
	"log"
)

func NewZanzibarClient(c devinit.Config) *ZanzibarClient {
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

	// TODO: use this when connection to local spicedb works
	//request := &pb.WriteSchemaRequest{Schema: schema}
	//_, err = client.WriteSchema(context.Background(), request)
	//if err != nil {
	//	log.Fatalf("failed to write schema: %s", err)
	//}

	return &ZanzibarClient{
		z:      client,
		prefix: c.ZanzibarPrefix,
	}
}

type ZanzibarClient struct {
	z      *authzed.Client
	prefix string
}
