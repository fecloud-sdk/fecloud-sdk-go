package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListRestoreTimesResponse Response Object
type ListRestoreTimesResponse struct {

	// 可恢复时间段列表。
	RestoreTime    *[]ListRestoreTimesResponseBodyRestoreTime `json:"restore_time,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o ListRestoreTimesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreTimesResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreTimesResponse", string(data)}, " ")
}
