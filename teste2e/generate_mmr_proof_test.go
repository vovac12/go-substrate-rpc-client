// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package teste2e

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	gsrpc "github.com/snowfork/go-substrate-rpc-client/v2"
	"github.com/snowfork/go-substrate-rpc-client/v2/config"
	"github.com/snowfork/go-substrate-rpc-client/v2/types"
)

type LeafProof struct {
	BlockHash types.Hash
	Leaf      []byte
	Proof     []byte
}

func TestGenerateMMRProof(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping end-to-end test in short mode.")
	}

	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}

	blockHashLatest, err := api.RPC.Chain.GetBlockHashLatest()
	if err != nil {
		panic(err)
	}

	block, err := api.RPC.Chain.GetBlock(blockHashLatest)
	if err != nil {
		panic(err)
	}

	blockNumberBytes := make([]byte, 32)
	binary.LittleEndian.PutUint64(blockNumberBytes, uint64(block.Block.Header.Number))
	blockNumberStr := hex.EncodeToString(blockNumberBytes[:])
	blockNumberHexPrefixed := "0x" + blockNumberStr

	leafProof := LeafProof{}
	err = api.Client.Call(&leafProof, "mmr_generateProof", 0, blockNumberHexPrefixed)
	if err != nil {
		panic(err)
	}

	t.Log("Leaf proof:", leafProof)
}
