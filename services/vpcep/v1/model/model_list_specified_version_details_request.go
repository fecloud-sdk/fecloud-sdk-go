package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSpecifiedVersionDetailsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	Version string `json:"version"`
}

func (o ListSpecifiedVersionDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecifiedVersionDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSpecifiedVersionDetailsRequest", string(data)}, " ")
}
