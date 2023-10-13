// Package korecontract Right f first refusal details
package korecontract

import (
	"kore_chaincode/core/status"
)

// IsExists bool
func (data RightOfFirstRefusal) IsExists() (bool, error) {
	if data.Exists {
		return true, nil
	}
	return false, status.ErrBadRequest.WithMessage("Right of first refusal does not hold for this koresecurites.")
}

// IsNormalization bool
func (data RightOfFirstRefusal) IsNormalization() bool {
	return data.AllocationBasis == 2
}
