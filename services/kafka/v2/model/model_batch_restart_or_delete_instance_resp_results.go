package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchRestartOrDeleteInstanceRespResults struct {
	Result *string `json:"result,omitempty"`

	Instance *string `json:"instance,omitempty"`
}

func (o BatchRestartOrDeleteInstanceRespResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestartOrDeleteInstanceRespResults struct{}"
	}

	return strings.Join([]string{"BatchRestartOrDeleteInstanceRespResults", string(data)}, " ")
}
