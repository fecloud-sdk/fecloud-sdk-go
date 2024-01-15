package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAutoLaunchChangeHistoriesResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]AutoLaunchChangeResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListAutoLaunchChangeHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchChangeHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchChangeHistoriesResponse", string(data)}, " ")
}
