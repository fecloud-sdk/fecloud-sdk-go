package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteServersRequestBody struct {
	DeletePublicip *bool `json:"delete_publicip,omitempty"`

	DeleteVolume *bool `json:"delete_volume,omitempty"`

	Servers []ServerId `json:"servers"`
}

func (o DeleteServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServersRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteServersRequestBody", string(data)}, " ")
}
