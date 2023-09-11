package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListDrRelationsRequest Request Object
type ListDrRelationsRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListDrRelationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrRelationsRequest struct{}"
	}

	return strings.Join([]string{"ListDrRelationsRequest", string(data)}, " ")
}
