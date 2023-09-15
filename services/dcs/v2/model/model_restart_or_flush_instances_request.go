package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// RestartOrFlushInstancesRequest Request Object
type RestartOrFlushInstancesRequest struct {
	Body *ChangeInstanceStatusBody `json:"body,omitempty"`
}

func (o RestartOrFlushInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartOrFlushInstancesRequest struct{}"
	}

	return strings.Join([]string{"RestartOrFlushInstancesRequest", string(data)}, " ")
}
