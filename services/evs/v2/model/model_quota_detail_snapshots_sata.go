package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QuotaDetailSnapshotsSata struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailSnapshotsSata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailSnapshotsSata struct{}"
	}

	return strings.Join([]string{"QuotaDetailSnapshotsSata", string(data)}, " ")
}
