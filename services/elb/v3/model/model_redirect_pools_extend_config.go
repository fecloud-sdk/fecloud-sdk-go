package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RedirectPoolsExtendConfig struct {
	RewriteUrlEnable *bool `json:"rewrite_url_enable,omitempty"`

	RewriteUrlConfig *RewriteUrlConfig `json:"rewrite_url_config,omitempty"`
}

func (o RedirectPoolsExtendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedirectPoolsExtendConfig struct{}"
	}

	return strings.Join([]string{"RedirectPoolsExtendConfig", string(data)}, " ")
}
