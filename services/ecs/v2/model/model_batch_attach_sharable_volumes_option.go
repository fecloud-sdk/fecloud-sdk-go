package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchAttachSharableVolumesOption struct {
	ServerId string `json:"server_id"`

	Device *string `json:"device,omitempty"`
}

func (o BatchAttachSharableVolumesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachSharableVolumesOption struct{}"
	}

	return strings.Join([]string{"BatchAttachSharableVolumesOption", string(data)}, " ")
}
