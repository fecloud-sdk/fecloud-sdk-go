package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointPolicyRequestBody struct {
	PolicyStatement []PolicyStatement `json:"policy_statement"`
}

func (o UpdateEndpointPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointPolicyRequestBody", string(data)}, " ")
}
