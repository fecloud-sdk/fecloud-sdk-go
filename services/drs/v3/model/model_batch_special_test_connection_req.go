package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchSpecialTestConnectionReq struct {
	Jobs []BatchJobActionReq `json:"jobs"`
}

func (o BatchSpecialTestConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSpecialTestConnectionReq struct{}"
	}

	return strings.Join([]string{"BatchSpecialTestConnectionReq", string(data)}, " ")
}
