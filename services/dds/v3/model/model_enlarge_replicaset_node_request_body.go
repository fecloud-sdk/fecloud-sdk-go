package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EnlargeReplicasetNodeRequestBody struct {
	Num int32 `json:"num"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o EnlargeReplicasetNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeReplicasetNodeRequestBody struct{}"
	}

	return strings.Join([]string{"EnlargeReplicasetNodeRequestBody", string(data)}, " ")
}
