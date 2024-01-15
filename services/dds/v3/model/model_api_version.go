package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ApiVersion struct {
	Id string `json:"id"`

	Links []Links `json:"links"`

	Status string `json:"status"`

	Version string `json:"version"`

	MinVersion string `json:"min_version"`

	Updated string `json:"updated"`
}

func (o ApiVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersion struct{}"
	}

	return strings.Join([]string{"ApiVersion", string(data)}, " ")
}
