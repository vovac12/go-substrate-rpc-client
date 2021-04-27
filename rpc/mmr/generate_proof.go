package mmr

import (
	"github.com/snowfork/go-substrate-rpc-client/v2/types"
)

// GenerateProof retrieves a proof and leaf
func (c *MMR) GenerateProof(leafIndex uint64) (types.GenerateMMRProofResponse, error) {
	var res types.GenerateMMRProofResponse
	err := c.client.Call(&res, "mmr_generateProof", leafIndex)
	if err != nil {
		return types.GenerateMMRProofResponse{}, err
	}

	return res, nil
}
