package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ApiVersionResponse struct {
	Id string `json:"id"`

	Links []Links `json:"links"`

	Status string `json:"status"`

	Version string `json:"version"`

	MinVersion string `json:"min_version"`

	Updated string `json:"updated"`
}

func (o ApiVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ApiVersionResponse", string(data)}, " ")
}
