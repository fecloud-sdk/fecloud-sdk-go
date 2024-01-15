package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPtrRecordSetRequest struct {
	Region string `json:"region"`

	FloatingipId string `json:"floatingip_id"`
}

func (o ShowPtrRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPtrRecordSetRequest struct{}"
	}

	return strings.Join([]string{"ShowPtrRecordSetRequest", string(data)}, " ")
}
