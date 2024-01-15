package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateGlobalVariableRequest struct {
	VarName string `json:"var_name"`

	Body *UpdateGlobalVariableRequestBody `json:"body,omitempty"`
}

func (o UpdateGlobalVariableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalVariableRequest struct{}"
	}

	return strings.Join([]string{"UpdateGlobalVariableRequest", string(data)}, " ")
}
