package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RewriteUrlConfig struct {
	Host *string `json:"host,omitempty"`

	Path *string `json:"path,omitempty"`

	Query *string `json:"query,omitempty"`
}

func (o RewriteUrlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RewriteUrlConfig struct{}"
	}

	return strings.Join([]string{"RewriteUrlConfig", string(data)}, " ")
}
