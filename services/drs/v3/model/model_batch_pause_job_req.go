package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchPauseJobReq struct {
	Jobs []PauseInfo `json:"jobs"`
}

func (o BatchPauseJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPauseJobReq struct{}"
	}

	return strings.Join([]string{"BatchPauseJobReq", string(data)}, " ")
}
