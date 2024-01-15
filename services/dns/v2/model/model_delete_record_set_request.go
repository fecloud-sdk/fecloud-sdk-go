package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteRecordSetRequest struct {
	ZoneId string `json:"zone_id"`

	RecordsetId string `json:"recordset_id"`
}

func (o DeleteRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordSetRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordSetRequest", string(data)}, " ")
}
