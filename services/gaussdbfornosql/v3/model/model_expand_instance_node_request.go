package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExpandInstanceNodeRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ExpandInstanceNodeRequestBody `json:"body,omitempty"`
}

func (o ExpandInstanceNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodeRequest struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodeRequest", string(data)}, " ")
}
