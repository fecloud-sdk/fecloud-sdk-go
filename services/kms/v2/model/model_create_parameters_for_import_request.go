package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateParametersForImportRequest struct {
	Body *GetParametersForImportRequestBody `json:"body,omitempty"`
}

func (o CreateParametersForImportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateParametersForImportRequest struct{}"
	}

	return strings.Join([]string{"CreateParametersForImportRequest", string(data)}, " ")
}
