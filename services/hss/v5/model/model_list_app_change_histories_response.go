package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAppChangeHistoriesResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]AppChangeResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListAppChangeHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppChangeHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAppChangeHistoriesResponse", string(data)}, " ")
}
