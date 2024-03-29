package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QuotaDetailVolumesGpssd struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailVolumesGpssd) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailVolumesGpssd struct{}"
	}

	return strings.Join([]string{"QuotaDetailVolumesGpssd", string(data)}, " ")
}
