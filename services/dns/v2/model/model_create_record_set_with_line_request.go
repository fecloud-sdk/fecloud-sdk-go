package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateRecordSetWithLineRequest struct {
	ZoneId string `json:"zone_id"`

	Body *CreateRecordSetWithLineRequestBody `json:"body,omitempty"`
}

func (o CreateRecordSetWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithLineRequest", string(data)}, " ")
}
