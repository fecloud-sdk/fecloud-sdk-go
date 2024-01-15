package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointWhiteRequestBody struct {
	Whitelist *[]string `json:"whitelist,omitempty"`

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`
}

func (o UpdateEndpointWhiteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointWhiteRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointWhiteRequestBody", string(data)}, " ")
}
