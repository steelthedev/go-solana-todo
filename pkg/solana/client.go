package solana

import (
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

var Client *client.Client

func InitClient() {
	Client = client.NewClient(rpc.DevnetRPCEndpoint)
}
