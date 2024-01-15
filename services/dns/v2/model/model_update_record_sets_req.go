package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateRecordSetsReq struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Type string `json:"type"`

	Ttl *int32 `json:"ttl,omitempty"`

	Records *[]string `json:"records,omitempty"`

	Weight *int32 `json:"weight,omitempty"`
}

func (o UpdateRecordSetsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordSetsReq struct{}"
	}

	return strings.Join([]string{"UpdateRecordSetsReq", string(data)}, " ")
}
