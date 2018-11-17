// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package ccid

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// GetID parses signed-proposal header and returns chaincode ID
func GetID(stub shim.ChaincodeStubInterface) (string, error) {
	proposal, err := stub.GetSignedProposal()
	if err != nil {
		return "", err
	}
	header := &peer.ChaincodeHeaderExtension{}
	if err = proto.Unmarshal(proposal.GetProposalBytes(), header); err != nil {
		return "", err
	}
	path := header.GetChaincodeId().GetPath()
	ei := path[7] + 8 // path[7]: id length
	return path[8:ei], nil
}
