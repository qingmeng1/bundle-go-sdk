package bundleclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	types2 "github.com/node-real/bundle-go-sdk/pkg/types"
)

type Client interface {
	SendBundle(context.Context, types.SendBundleArgs) (common.Hash, error)
	QueryBundle(context.Context, common.Hash) (*types2.Bundle, error)
	BundlePrice(context.Context) (*big.Int, error)
	Builders(context.Context) ([]common.Address, error)
	Validators(context.Context) ([]common.Address, error)
}

type client struct {
	c *rpc.Client
}

func New(rpcCli *rpc.Client) Client {
	return &client{rpcCli}
}

func (c *client) SendBundle(ctx context.Context, args types.SendBundleArgs) (common.Hash, error) {
	var hash common.Hash
	err := c.c.CallContext(ctx, &hash, "eth_sendBundle", args)
	if err != nil {
		return common.Hash{}, err
	}
	return hash, nil
}

func (c *client) QueryBundle(ctx context.Context, bundleHash common.Hash) (*types2.Bundle, error) {
	var bundle types2.Bundle
	err := c.c.CallContext(ctx, &bundle, "eth_queryBundle", bundleHash)
	if err != nil {
		return nil, err
	}
	return &bundle, nil
}

func (c *client) BundlePrice(ctx context.Context) (*big.Int, error) {
	var price *big.Int
	err := c.c.CallContext(ctx, &price, "eth_bundlePrice")
	if err != nil {
		return nil, err
	}
	return price, nil
}

func (c *client) Builders(ctx context.Context) ([]common.Address, error) {
	var builders []common.Address
	err := c.c.CallContext(ctx, &builders, "eth_builders")
	if err != nil {
		return nil, err
	}
	return builders, nil
}

func (c *client) Validators(ctx context.Context) ([]common.Address, error) {
	var validators []common.Address
	err := c.c.CallContext(ctx, &validators, "eth_validators")
	if err != nil {
		return nil, err
	}
	return validators, nil
}
