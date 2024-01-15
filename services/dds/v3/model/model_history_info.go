package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HistoryInfo struct {
	ParameterName string `json:"parameter_name"`

	OldValue string `json:"old_value"`

	NewValue string `json:"new_value"`

	UpdatedAt string `json:"updated_at"`
}

func (o HistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryInfo struct{}"
	}

	return strings.Join([]string{"HistoryInfo", string(data)}, " ")
}
