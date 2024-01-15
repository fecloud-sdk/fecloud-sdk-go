package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSlowlogForLtsResponse struct {
	SlowLogList *[]MysqlSlowLogDetailsItem `json:"slow_log_list,omitempty"`

	LongQueryTime  *string `json:"long_query_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSlowlogForLtsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowlogForLtsResponse struct{}"
	}

	return strings.Join([]string{"ListSlowlogForLtsResponse", string(data)}, " ")
}
