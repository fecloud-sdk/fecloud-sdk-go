package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateJobReq struct {
	Jobs []CreateJobReq `json:"jobs"`
}

func (o BatchCreateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateJobReq struct{}"
	}

	return strings.Join([]string{"BatchCreateJobReq", string(data)}, " ")
}
