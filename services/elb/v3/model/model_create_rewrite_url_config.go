package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateRewriteUrlConfig struct {
	Host *string `json:"host,omitempty"`

	Path *string `json:"path,omitempty"`

	Query *string `json:"query,omitempty"`
}

func (o CreateRewriteUrlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRewriteUrlConfig struct{}"
	}

	return strings.Join([]string{"CreateRewriteUrlConfig", string(data)}, " ")
}
