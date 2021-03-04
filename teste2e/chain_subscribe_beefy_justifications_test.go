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
	"context"
	"fmt"
	"testing"
	"time"

	gsrpc "github.com/snowfork/go-substrate-rpc-client/v2"
	"github.com/snowfork/go-substrate-rpc-client/v2/config"
	"github.com/snowfork/go-substrate-rpc-client/v2/types"
	"github.com/stretchr/testify/assert"
)

type Commitment struct {
	Payload        types.H256        // payload
	BlockNumber    types.BlockNumber // block_number
	ValidatorSetID types.U64         // validator_set_id
}

type SignedCommitment struct {
	Commitment Commitment // `json:"commitment"`
	Signatures types.Data // `json:"signatures"`
	// Signatures []OptionBeefySignature // `json:"signatures"`
}

func TestChain_SubscribeBeefyJustifications(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping end-to-end test in short mode.")
	}

	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}

	ch := make(chan types.Bytes)
	// ch := make(chan SignedCommitment)

	t.Log("1")

	sub, err := api.Client.Subscribe(context.Background(), "beefy", "subscribeJustifications", "unsubscribeJustifications", "justifications", ch)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
	t.Log("2")

	timeout := time.After(40 * time.Second)
	received := 0

	for {
		select {
		case msg := <-ch:
			t.Log("3")
			fmt.Printf("%#v\n", msg)

			s := &SignedCommitment{}
			err := types.DecodeFromBytes(msg, s)
			if err != nil {
				panic(err)
			}

			t.Log("4")
			fmt.Printf("%#v\n", s)

			received++

			if received >= 2 {
				return
			}
		case <-timeout:
			t.Log("5")
			assert.FailNow(t, "timeout reached without getting 2 notifications from subscription")
			return
		}
	}
}
