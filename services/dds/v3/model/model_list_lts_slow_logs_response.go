package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLtsSlowLogsResponse struct {
	SlowLogs       *[]SlowLogDetail `json:"slow_logs,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListLtsSlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsSlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLtsSlowLogsResponse", string(data)}, " ")
}
