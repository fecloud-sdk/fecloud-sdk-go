package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRestoreTimesResponseBodyRestoreTime struct {
	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`
}

func (o ListRestoreTimesResponseBodyRestoreTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreTimesResponseBodyRestoreTime struct{}"
	}

	return strings.Join([]string{"ListRestoreTimesResponseBodyRestoreTime", string(data)}, " ")
}
