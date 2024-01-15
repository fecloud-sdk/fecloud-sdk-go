package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVersionDetailsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`
}

func (o ListVersionDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListVersionDetailsRequest", string(data)}, " ")
}
