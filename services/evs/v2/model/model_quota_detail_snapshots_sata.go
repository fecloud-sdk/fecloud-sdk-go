package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// QuotaDetailSnapshotsSata SATA云硬盘类型预留快照个数，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）。
type QuotaDetailSnapshotsSata struct {

	// 已使用的数量。
	InUse int32 `json:"in_use"`

	// 最大的数量。
	Limit int32 `json:"limit"`

	// 预留属性。
	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailSnapshotsSata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailSnapshotsSata struct{}"
	}

	return strings.Join([]string{"QuotaDetailSnapshotsSata", string(data)}, " ")
}
