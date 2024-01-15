package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteRecordSetsRequest struct {
	ZoneId string `json:"zone_id"`

	RecordsetId string `json:"recordset_id"`
}

func (o DeleteRecordSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordSetsRequest", string(data)}, " ")
}
