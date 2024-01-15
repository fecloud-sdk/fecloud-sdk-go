package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteJobReq struct {
	Jobs []DeleteJobReq `json:"jobs"`
}

func (o BatchDeleteJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobReq", string(data)}, " ")
}
