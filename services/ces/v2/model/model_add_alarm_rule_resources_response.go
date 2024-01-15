package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddAlarmRuleResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddAlarmRuleResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAlarmRuleResourcesResponse struct{}"
	}

	return strings.Join([]string{"AddAlarmRuleResourcesResponse", string(data)}, " ")
}
