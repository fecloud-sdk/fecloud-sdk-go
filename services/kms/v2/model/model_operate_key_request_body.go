package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type OperateKeyRequestBody struct {
	KeyId string `json:"key_id"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o OperateKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateKeyRequestBody struct{}"
	}

	return strings.Join([]string{"OperateKeyRequestBody", string(data)}, " ")
}
