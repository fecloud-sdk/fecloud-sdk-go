package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CompareTableInfoWithToken struct {
	TableName string `json:"table_name"`

	MinToken *string `json:"min_token,omitempty"`

	MaxToken *string `json:"max_token,omitempty"`
}

func (o CompareTableInfoWithToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareTableInfoWithToken struct{}"
	}

	return strings.Join([]string{"CompareTableInfoWithToken", string(data)}, " ")
}
