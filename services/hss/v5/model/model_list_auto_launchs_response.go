package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAutoLaunchsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]AutoLauchResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListAutoLaunchsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchsResponse struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchsResponse", string(data)}, " ")
}
