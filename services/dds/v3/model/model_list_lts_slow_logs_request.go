package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLtsSlowLogsRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ListLtsSlowLogsRequestBody `json:"body,omitempty"`
}

func (o ListLtsSlowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsSlowLogsRequest struct{}"
	}

	return strings.Join([]string{"ListLtsSlowLogsRequest", string(data)}, " ")
}
