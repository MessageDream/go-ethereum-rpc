package rpc

import (
	"fmt"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

func (client *Client) RequestFilter(method string, params ...interface{}) (*rpctypes.HexString, error) {
	response, err := checkRPCError(client.Call(method, params...))

	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("response returned without error but no transaction found for %v", params)
	}

	return nil, nil
}
