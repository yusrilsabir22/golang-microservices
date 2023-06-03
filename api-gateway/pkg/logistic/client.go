package logistic

import (
	"fmt"

	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/config"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/logistic/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.LogisticServiceClient
}

func InitServiceClient(c *config.Config) pb.LogisticServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.LogisticSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewLogisticServiceClient(cc)
}
