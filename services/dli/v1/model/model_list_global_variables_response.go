package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListGlobalVariablesResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	GlobalVars *[]GlobalVariablesInfo `json:"global_vars,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGlobalVariablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalVariablesResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalVariablesResponse", string(data)}, " ")
}
