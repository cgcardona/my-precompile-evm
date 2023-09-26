// (c) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"fmt"

	"github.com/ava-labs/avalanchego/version"
	"github.com/ava-labs/subnet-evm/plugin/evm"
	"github.com/ava-labs/subnet-evm/plugin/runner"

	// Each precompile generated by the precompilegen tool has a self-registering init function
	// that registers the precompile with the subnet-evm. Importing the precompile package here
	// will cause the precompile to be registered with the subnet-evm.
	_ "github.com/ava-labs/precompile-evm/calculatorplus"
	_ "github.com/ava-labs/precompile-evm/md5"
	_ "github.com/ava-labs/precompile-evm/sha256"
	_ "github.com/ava-labs/precompile-evm/stringstore"
)

const precompileEVMVersion = "v0.0.0"

func main() {
	versionString := fmt.Sprintf("Precompile-EVM/%s Subnet-EVM/%s [AvalancheGo=%s, rpcchainvm=%d]", precompileEVMVersion, evm.Version, version.Current, version.RPCChainVMProtocol)
	runner.Run(versionString)
}
