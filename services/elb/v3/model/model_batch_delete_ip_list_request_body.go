package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteIpListRequestBody struct {
	Ipgroup *BatchDeleteIpListOption `json:"ipgroup,omitempty"`
}

func (o BatchDeleteIpListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpListRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpListRequestBody", string(data)}, " ")
}
