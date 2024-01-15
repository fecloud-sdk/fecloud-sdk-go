package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVulHostsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]VulHostInfo `json:"data_list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListVulHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHostsResponse struct{}"
	}

	return strings.Join([]string{"ListVulHostsResponse", string(data)}, " ")
}
