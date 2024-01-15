package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EnlargeVolumeObject struct {
	Size int32 `json:"size"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o EnlargeVolumeObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeVolumeObject struct{}"
	}

	return strings.Join([]string{"EnlargeVolumeObject", string(data)}, " ")
}
