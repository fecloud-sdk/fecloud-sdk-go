package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteGlobalVariableRequest struct {
	VarName string `json:"var_name"`
}

func (o DeleteGlobalVariableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalVariableRequest struct{}"
	}

	return strings.Join([]string{"DeleteGlobalVariableRequest", string(data)}, " ")
}
