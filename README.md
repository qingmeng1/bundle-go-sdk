# Bundle Go SDK


## Disclaimer
**The software and related documentation are under active development, all subject to potential future change without
notification and not ready for production use. The code and security audit have not been fully completed and not ready
for any bug bounty. We advise you to be careful and experiment on the network at your own risk. Stay safe out there.**

## Instruction

The BUNDLE-GO-SDK provides enhanced transaction privacy and atomicity for the BNB Smart Chain (BSC) network. By implementing the BEP322 standard, the following capabilities are provided:
1. Privacy. All transactions sent through this API will not be propagated on the P2P network, hence, they won't be detected by any third parties. This effectively prevents transactions from being targeted by sandwich attacks.
2. Batch transaction. Multiple transactions can be consolidated into a single 'bundle', which can then be transmitted through just one API call. The sequence of transactions within a block, as well as the order within a bundle, can be assured to maintain impeccable consistency.
3. Atomicity. Transactions within a bundle either all get included on the chain, or none at all. There's no such scenario where only a portion of the transactions are included on chain.
4. Gas protection. If a single transaction within a bundle fails, the entire bundle is guaranteed not to be packaged onto the blockchain. This mechanism safeguards users from unnecessary gas expenditure.

### Requirement

Go version above 1.21

## Getting started
To get started working with the SDK setup your project for Go modules, and retrieve the SDK dependencies with `go get`.
This example shows how you can use the bundle go SDK to interact with the bundle apis on bsc,

### Initialize Project

```sh
$ mkdir ~/hello_bundle
$ cd ~/hello_bundle
$ go mod init hello_bundle
```

### Add SDK Dependencies

```sh
$ go get github.com/node-real/bundle-go-sdk
```


```go
package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	endpointUrl := "https://bsc-testnet.nodereal.io/v1/{{your api key}}"

	rpcCli, err := rpc.Dial(endpointUrl)
	if err != nil {
		panic(err)
	}

	bundleCli := bundleclient.New(rpcCli)
}

```

###  Quick Start Examples

The examples directory provides a wealth of examples to guide users in using the SDK's various features

#### Config Examples

You need to modify the ENV variables to use the example:

APIKEY: your meganode apikey with growth tier and bundle-package subscribed refer [here](https://nodereal.io/api-marketplace/bsc-bundle-service-api) to subscribe
Address: your wallet address which used as toAddress of transaction
PrivateKey: your privateKey which used to sign transactions

#### Run Examples
The steps to run example are as follows
```
make examples
cd example
./example 
```

## Reference

[BSC MEV](https://docs.bnbchain.org/docs/mev/overview/)

[Meganode API Marketplace](https://nodereal.io/api-marketplace/bsc-bundle-service-api): subscribe those Apis and explore more packages.

[Api introduction](https://docs.nodereal.io/reference/bsc-bundle-service-api): to get details about those Apis