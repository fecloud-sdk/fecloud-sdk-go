package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRiskConfigHostsResponse struct {
	TotalNum *int64 `json:"total_num,omitempty"`

	DataList       *[]SecurityCheckHostInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListRiskConfigHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigHostsResponse struct{}"
	}

	return strings.Join([]string{"ListRiskConfigHostsResponse", string(data)}, " ")
}
