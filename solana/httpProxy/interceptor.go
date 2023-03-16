package httpProxy

import (
	"fmt"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/jsonrpc"
	"net/http"
	"time"
)

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	// before request
	fmt.Printf("[Logging] Sending request with body: %v\n", req.Body)
	// request sending execution
	resp, e := lrt.Proxied.RoundTrip(req)
	// handle result
	fmt.Printf("[Logging] Receiving responds with body: %v\n", resp.Body)
	return resp, e
}

func GetHttpClient() *http.Client {
	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: LoggingRoundTripper{http.DefaultTransport}}
}

func NewHTTP(
	timeout time.Duration,
) *http.Client {
	return &http.Client{
		Timeout:   timeout,
		Transport: LoggingRoundTripper{Proxied: http.DefaultTransport},
	}
}

func NewRPC(rpcEndpoint string) *rpc.Client {
	var (
		defaultTimeout = 25 * time.Second
	)
	opts := &jsonrpc.RPCClientOpts{
		HTTPClient: NewHTTP(
			defaultTimeout,
		),
	}
	rpcClient := jsonrpc.NewClientWithOpts(rpcEndpoint, opts)
	return rpc.NewWithCustomRPCClient(rpcClient)
}
