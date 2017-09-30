package main

import (
	"context"
	"fmt"
	"time"

	"github.com/osrg/gobgp/packet/bgp"

	api "github.com/osrg/gobgp/api"
	"github.com/osrg/gobgp/table"

	"google.golang.org/grpc"
)

func main() {
	addRoute(context.Background(), "203.0.113.50", 32, "192.0.2.0")
}

func addRoute(ctx context.Context, addr string, prefixLen int, nexthop string) error {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("couldn't connect to gobgpd: %v", err)
	}
	defer conn.Close()
	client := api.NewGobgpApiClient(conn)

	attrs := []bgp.PathAttributeInterface{
		bgp.NewPathAttributeOrigin(0), // IGP
		bgp.NewPathAttributeNextHop(nexthop),
	}
	path := table.NewPath(nil, bgp.NewIPAddrPrefix(uint8(prefixLen), addr), false, attrs, time.Now(), false)
	_, err = client.AddPath(ctx, &api.AddPathRequest{
		Resource: api.Resource_GLOBAL,
		Path:     api.ToPathApi(path),
	})
	if err != nil {
		return fmt.Errorf("coudln't add route: %v", err)
	}
	return nil
}
