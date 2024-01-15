package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRiskConfigsResponse struct {
	TotalNum *int64 `json:"total_num,omitempty"`

	DataList       *[]SecurityCheckInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListRiskConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListRiskConfigsResponse", string(data)}, " ")
}
