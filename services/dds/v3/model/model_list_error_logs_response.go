package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListErrorLogsResponse struct {
	ErrorLogList *[]ErrorlogResult `json:"error_log_list,omitempty"`

	TotalRecord    *int32 `json:"total_record,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListErrorLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorLogsResponse struct{}"
	}

	return strings.Join([]string{"ListErrorLogsResponse", string(data)}, " ")
}
