package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteImportedKeyMaterialRequest struct {
	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o DeleteImportedKeyMaterialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImportedKeyMaterialRequest struct{}"
	}

	return strings.Join([]string{"DeleteImportedKeyMaterialRequest", string(data)}, " ")
}
