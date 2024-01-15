package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListUserChangeHistoriesResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]UserChangeHistoryResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListUserChangeHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserChangeHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListUserChangeHistoriesResponse", string(data)}, " ")
}
