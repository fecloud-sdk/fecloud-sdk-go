package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListGlobalVariablesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListGlobalVariablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalVariablesRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalVariablesRequest", string(data)}, " ")
}
