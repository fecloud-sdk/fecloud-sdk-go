package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteImportedKeyMaterialResponse Response Object
type DeleteImportedKeyMaterialResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteImportedKeyMaterialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImportedKeyMaterialResponse struct{}"
	}

	return strings.Join([]string{"DeleteImportedKeyMaterialResponse", string(data)}, " ")
}
