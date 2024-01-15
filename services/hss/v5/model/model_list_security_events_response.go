package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSecurityEventsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]EventManagementResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListSecurityEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityEventsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityEventsResponse", string(data)}, " ")
}
