package todocli

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sadlil.com/samples/crud/pkg/clients"
	"sadlil.com/samples/crud/pkg/clients/grpctransport"
	"sadlil.com/samples/crud/pkg/clients/httptransport"
)

func newTodoServiceClient(transport, addr string) (clients.TodoServiceClient, error) {
	switch transport {
	case "http":
		return httptransport.NewClient(
			httptransport.WithBaseURL(addr),
		)
	case "grpc":
		return grpctransport.NewClient(
			grpctransport.WithServerAddress(addr),
			grpctransport.WithDialOptions(
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			),
		)
	}
	return nil, fmt.Errorf("unknown transport: %v", transport)
}
