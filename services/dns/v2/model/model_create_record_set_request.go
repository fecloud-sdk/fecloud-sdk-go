package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateRecordSetRequest struct {
	ZoneId string `json:"zone_id"`

	Body *CreateRecordSetRequestBody `json:"body,omitempty"`
}

func (o CreateRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordSetRequest", string(data)}, " ")
}
