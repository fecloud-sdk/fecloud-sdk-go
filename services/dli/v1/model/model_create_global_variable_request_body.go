package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateGlobalVariableRequestBody struct {
	VarName string `json:"var_name"`

	VarValue string `json:"var_value"`

	IsSensitive *bool `json:"is_sensitive,omitempty"`
}

func (o CreateGlobalVariableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalVariableRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGlobalVariableRequestBody", string(data)}, " ")
}
