package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicAccessPolicyReq struct {
	Topics []AccessPolicyTopicEntity `json:"topics"`
}

func (o UpdateTopicAccessPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyReq struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyReq", string(data)}, " ")
}
