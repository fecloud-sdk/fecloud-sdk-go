package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchAttachSharableVolumesRequestBody struct {
	Serverinfo []BatchAttachSharableVolumesOption `json:"serverinfo"`
}

func (o BatchAttachSharableVolumesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachSharableVolumesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAttachSharableVolumesRequestBody", string(data)}, " ")
}
