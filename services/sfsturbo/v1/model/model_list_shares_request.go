package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSharesRequest struct {
	ContentType string `json:"Content-Type"`

	Limit *int64 `json:"limit,omitempty"`

	Offset *int64 `json:"offset,omitempty"`
}

func (o ListSharesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharesRequest struct{}"
	}

	return strings.Join([]string{"ListSharesRequest", string(data)}, " ")
}
