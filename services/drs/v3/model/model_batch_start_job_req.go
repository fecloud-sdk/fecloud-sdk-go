package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchStartJobReq struct {
	Jobs []StartInfo `json:"jobs"`
}

func (o BatchStartJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartJobReq struct{}"
	}

	return strings.Join([]string{"BatchStartJobReq", string(data)}, " ")
}
