package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NsRecords struct {
	Hostname *string `json:"hostname,omitempty"`

	Address *string `json:"address,omitempty"`

	Priority *int32 `json:"priority,omitempty"`
}

func (o NsRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NsRecords struct{}"
	}

	return strings.Join([]string{"NsRecords", string(data)}, " ")
}
