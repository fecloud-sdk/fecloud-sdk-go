package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostRaspProtectHistoryInfoResponse struct {
	TotalNum *int64 `json:"total_num,omitempty"`

	DataList       *[]HostRaspProtectHistoryResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListHostRaspProtectHistoryInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostRaspProtectHistoryInfoResponse struct{}"
	}

	return strings.Join([]string{"ListHostRaspProtectHistoryInfoResponse", string(data)}, " ")
}
