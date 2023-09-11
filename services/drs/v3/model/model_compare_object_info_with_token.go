package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CompareObjectInfoWithToken
type CompareObjectInfoWithToken struct {

	// 库名。
	DbName string `json:"db_name"`

	// 该库下的表信息列表（带token）。
	TableNameWithToken *[]CompareTableInfoWithToken `json:"table_name_with_token,omitempty"`
}

func (o CompareObjectInfoWithToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareObjectInfoWithToken struct{}"
	}

	return strings.Join([]string{"CompareObjectInfoWithToken", string(data)}, " ")
}
