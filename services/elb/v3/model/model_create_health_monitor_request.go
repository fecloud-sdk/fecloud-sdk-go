package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateHealthMonitorRequest Request Object
type CreateHealthMonitorRequest struct {
	Body *CreateHealthMonitorRequestBody `json:"body,omitempty"`
}

func (o CreateHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"CreateHealthMonitorRequest", string(data)}, " ")
}
