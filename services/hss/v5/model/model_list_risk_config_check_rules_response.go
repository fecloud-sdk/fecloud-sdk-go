package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRiskConfigCheckRulesResponse struct {
	TotalNum *int64 `json:"total_num,omitempty"`

	DataList       *[]CheckRuleRiskInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListRiskConfigCheckRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigCheckRulesResponse struct{}"
	}

	return strings.Join([]string{"ListRiskConfigCheckRulesResponse", string(data)}, " ")
}
