package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAppsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]AppResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsResponse struct{}"
	}

	return strings.Join([]string{"ListAppsResponse", string(data)}, " ")
}
