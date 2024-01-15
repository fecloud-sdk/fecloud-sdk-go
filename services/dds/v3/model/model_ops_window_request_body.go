package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type OpsWindowRequestBody struct {
	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`
}

func (o OpsWindowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpsWindowRequestBody struct{}"
	}

	return strings.Join([]string{"OpsWindowRequestBody", string(data)}, " ")
}
