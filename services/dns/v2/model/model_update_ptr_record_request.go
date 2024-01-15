package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePtrRecordRequest struct {
	Region string `json:"region"`

	FloatingipId string `json:"floatingip_id"`

	Body *UpdatePtrReq `json:"body,omitempty"`
}

func (o UpdatePtrRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePtrRecordRequest struct{}"
	}

	return strings.Join([]string{"UpdatePtrRecordRequest", string(data)}, " ")
}
