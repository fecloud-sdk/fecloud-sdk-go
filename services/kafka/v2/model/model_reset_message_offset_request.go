package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetMessageOffsetRequest struct {
	InstanceId string `json:"instance_id"`

	Group string `json:"group"`

	Body *ResetMessageOffsetReq `json:"body,omitempty"`
}

func (o ResetMessageOffsetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetMessageOffsetRequest struct{}"
	}

	return strings.Join([]string{"ResetMessageOffsetRequest", string(data)}, " ")
}
