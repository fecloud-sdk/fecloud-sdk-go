package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateRecordSetsRequest struct {
	ZoneId string `json:"zone_id"`

	RecordsetId string `json:"recordset_id"`

	Body *UpdateRecordSetsReq `json:"body,omitempty"`
}

func (o UpdateRecordSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecordSetsRequest", string(data)}, " ")
}
