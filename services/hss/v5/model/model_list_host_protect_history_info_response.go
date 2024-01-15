package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostProtectHistoryInfoResponse struct {
	HostName *string `json:"host_name,omitempty"`

	ProtectStatus *string `json:"protect_status,omitempty"`

	TotalNum *int64 `json:"total_num,omitempty"`

	DataList       *[]HostProtectHistoryResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListHostProtectHistoryInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostProtectHistoryInfoResponse struct{}"
	}

	return strings.Join([]string{"ListHostProtectHistoryInfoResponse", string(data)}, " ")
}
