package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSpecifiedVersionDetailsResponse struct {
	Version        *VersionObject `json:"version,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSpecifiedVersionDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecifiedVersionDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSpecifiedVersionDetailsResponse", string(data)}, " ")
}
