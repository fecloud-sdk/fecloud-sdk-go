package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSharesResponse struct {
	Shares *[]Shares `json:"shares,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSharesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharesResponse struct{}"
	}

	return strings.Join([]string{"ListSharesResponse", string(data)}, " ")
}
