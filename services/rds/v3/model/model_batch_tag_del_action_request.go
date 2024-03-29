package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchTagDelActionRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *BatchTagActionDelRequestBody `json:"body,omitempty"`
}

func (o BatchTagDelActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagDelActionRequest struct{}"
	}

	return strings.Join([]string{"BatchTagDelActionRequest", string(data)}, " ")
}
