package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlavorAzObject struct {
	Capacity *string `json:"capacity,omitempty"`

	AzCodes *[]string `json:"az_codes,omitempty"`
}

func (o FlavorAzObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorAzObject struct{}"
	}

	return strings.Join([]string{"FlavorAzObject", string(data)}, " ")
}
