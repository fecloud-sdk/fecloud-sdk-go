package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowL7policyResponse struct {
	L7policy       *L7policyResp `json:"l7policy,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowL7policyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7policyResponse struct{}"
	}

	return strings.Join([]string{"ShowL7policyResponse", string(data)}, " ")
}
