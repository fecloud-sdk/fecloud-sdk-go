package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NameServersResp struct {
	Type *string `json:"type,omitempty"`

	Region *string `json:"region,omitempty"`

	NsRecords *[]NsRecords `json:"ns_records,omitempty"`
}

func (o NameServersResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameServersResp struct{}"
	}

	return strings.Join([]string{"NameServersResp", string(data)}, " ")
}
