// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package ccid

import (
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/utils"
)

// GetID parses signed-proposal header and returns chaincode ID
func GetID(stub shim.ChaincodeStubInterface) (string, error) {
	sp, err := stub.GetSignedProposal()
	if err != nil {
		return "", err
	}
	proposal, err := utils.GetProposal(sp.GetProposalBytes())
	if err != nil {
		return "", err
	}
	invocation, err := utils.GetChaincodeInvocationSpec(proposal)
	if err != nil {
		return "", err
	}
	spec := invocation.GetChaincodeSpec()
	if nil == spec {
		return "", errors.New("failed to get chaincode spec")
	}
	id := spec.GetChaincodeId()
	if nil == id {
		return "", errors.New("failed to get chaincode ID")
	}
	return id.GetName(), nil
}
