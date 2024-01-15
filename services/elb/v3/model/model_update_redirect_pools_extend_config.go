package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateRedirectPoolsExtendConfig struct {
	RewriteUrlEnable *bool `json:"rewrite_url_enable,omitempty"`

	RewriteUrlConfig *CreateRewriteUrlConfig `json:"rewrite_url_config,omitempty"`
}

func (o UpdateRedirectPoolsExtendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRedirectPoolsExtendConfig struct{}"
	}

	return strings.Join([]string{"UpdateRedirectPoolsExtendConfig", string(data)}, " ")
}
