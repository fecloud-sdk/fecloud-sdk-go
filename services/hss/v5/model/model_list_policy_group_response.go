package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPolicyGroupResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]PolicyGroupResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyGroupResponse", string(data)}, " ")
}
