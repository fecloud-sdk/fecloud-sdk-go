package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteAlarmRuleResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAlarmRuleResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRuleResourcesResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRuleResourcesResponse", string(data)}, " ")
}
