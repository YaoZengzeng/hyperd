package main

import (
	"fmt"
	"os"

	"github.com/hyperhq/hyperd/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	server = "127.0.0.1:22318"
)

func main() {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Connect server error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := types.NewPublicAPIClient(conn)
/*	request := types.PodInfoRequest{
		PodID: "pod-UDPChNpcCq",
	}
	response, err := client.PodInfo(context.Background(), &request)
	if err != nil {
		fmt.Printf("Get PodInfo error: %v", err)
		os.Exit(1)
	}*/
/*	request := types.PodListRequest{
		PodID:	"",
		VmID:	"",
	}
	response, err := client.PodList(context.Background(), &request)
	if err != nil {
		fmt.Printf("Get PodList error: %v", err)
		os.Exit(1)
	}*/

	request := types.PodRemoveRequest{
		PodID: "pod-UDPChNpcCq",
	}
	response, err := client.PodRemove(context.Background(), &request)
	if err != nil {
		fmt.Printf("Get PodRemove error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Got response: %v", response)
}
