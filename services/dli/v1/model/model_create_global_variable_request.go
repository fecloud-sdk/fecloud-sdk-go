package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateGlobalVariableRequest struct {
	Body *CreateGlobalVariableRequestBody `json:"body,omitempty"`
}

func (o CreateGlobalVariableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalVariableRequest struct{}"
	}

	return strings.Join([]string{"CreateGlobalVariableRequest", string(data)}, " ")
}
