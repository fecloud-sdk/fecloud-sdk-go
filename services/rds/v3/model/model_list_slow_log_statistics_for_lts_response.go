package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListSlowLogStatisticsForLtsResponse Response Object
type ListSlowLogStatisticsForLtsResponse struct {

	// 数据集合。
	SlowLogList *[]MysqlSlowLogStatisticsItem `json:"slow_log_list,omitempty"`

	// 总条数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSlowLogStatisticsForLtsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogStatisticsForLtsResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogStatisticsForLtsResponse", string(data)}, " ")
}
