package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateAlarmRequest struct {
	Body *CreateAlarmRequestBody `json:"body,omitempty"`
}

func (o CreateAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmRequest struct{}"
	}

	return strings.Join([]string{"CreateAlarmRequest", string(data)}, " ")
}
