package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListKeyStoresRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListKeyStoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeyStoresRequest struct{}"
	}

	return strings.Join([]string{"ListKeyStoresRequest", string(data)}, " ")
}
