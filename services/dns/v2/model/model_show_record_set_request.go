package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowRecordSetRequest struct {
	ZoneId string `json:"zone_id"`

	RecordsetId string `json:"recordset_id"`
}

func (o ShowRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordSetRequest", string(data)}, " ")
}
