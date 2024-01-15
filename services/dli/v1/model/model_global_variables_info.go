package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GlobalVariablesInfo struct {
	Id *int64 `json:"id,omitempty"`

	VarName *string `json:"var_name,omitempty"`

	VarValue *string `json:"var_value,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	UserId *string `json:"user_id,omitempty"`

	IsSensitive *bool `json:"is_sensitive,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o GlobalVariablesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalVariablesInfo struct{}"
	}

	return strings.Join([]string{"GlobalVariablesInfo", string(data)}, " ")
}
