package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// SetReadOnlySwitchResponse Response Object
type SetReadOnlySwitchResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetReadOnlySwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetReadOnlySwitchResponse struct{}"
	}

	return strings.Join([]string{"SetReadOnlySwitchResponse", string(data)}, " ")
}
