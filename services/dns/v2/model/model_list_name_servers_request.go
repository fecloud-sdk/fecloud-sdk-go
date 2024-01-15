package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListNameServersRequest struct {
	Type *string `json:"type,omitempty"`

	Region *string `json:"region,omitempty"`
}

func (o ListNameServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNameServersRequest struct{}"
	}

	return strings.Join([]string{"ListNameServersRequest", string(data)}, " ")
}
