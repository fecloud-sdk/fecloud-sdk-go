package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateGlobalVariableRequestBody struct {
	VarValue string `json:"var_value"`
}

func (o UpdateGlobalVariableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalVariableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGlobalVariableRequestBody", string(data)}, " ")
}
