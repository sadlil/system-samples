package todocli

import (
	"fmt"

	"github.com/sadlil/system-samples/crud/pkg/clients"
	"github.com/sadlil/system-samples/crud/pkg/clients/grpctransport"
	"github.com/sadlil/system-samples/crud/pkg/clients/httptransport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
