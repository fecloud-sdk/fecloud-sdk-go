package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateRecordSetRequest struct {
	ZoneId string `json:"zone_id"`

	RecordsetId string `json:"recordset_id"`

	Body *UpdateRecordSetReq `json:"body,omitempty"`
}

func (o UpdateRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordSetRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecordSetRequest", string(data)}, " ")
}
