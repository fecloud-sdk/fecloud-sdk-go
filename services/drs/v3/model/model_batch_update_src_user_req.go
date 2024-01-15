package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateSrcUserReq struct {
	Jobs []UpdateUserReq `json:"jobs"`
}

func (o BatchUpdateSrcUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateSrcUserReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateSrcUserReq", string(data)}, " ")
}
