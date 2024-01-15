package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShrinkInstanceNodeRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ShrinkInstanceNodeRequestBody `json:"body,omitempty"`
}

func (o ShrinkInstanceNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodeRequest struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodeRequest", string(data)}, " ")
}
