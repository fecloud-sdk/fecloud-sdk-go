package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVulnerabilitiesResponse struct {
	TotalNum *int64 `json:"total_num,omitempty"`

	DataList       *[]VulInfo `json:"data_list,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListVulnerabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulnerabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ListVulnerabilitiesResponse", string(data)}, " ")
}
