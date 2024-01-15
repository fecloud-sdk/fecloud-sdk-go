package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateRecordSetReq struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Type string `json:"type"`

	Ttl *int32 `json:"ttl,omitempty"`

	Records *[]string `json:"records,omitempty"`
}

func (o UpdateRecordSetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordSetReq struct{}"
	}

	return strings.Join([]string{"UpdateRecordSetReq", string(data)}, " ")
}
