package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpgradeDbVersionResponse Response Object
type UpgradeDbVersionResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeDbVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDbVersionResponse struct{}"
	}

	return strings.Join([]string{"UpgradeDbVersionResponse", string(data)}, " ")
}
