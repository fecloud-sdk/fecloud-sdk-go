package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostGroupsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]HostGroupItem `json:"data_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListHostGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListHostGroupsResponse", string(data)}, " ")
}
