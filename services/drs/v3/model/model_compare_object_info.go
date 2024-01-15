package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CompareObjectInfo struct {
	DbName string `json:"db_name"`

	TableName *[]string `json:"table_name,omitempty"`
}

func (o CompareObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareObjectInfo struct{}"
	}

	return strings.Join([]string{"CompareObjectInfo", string(data)}, " ")
}
