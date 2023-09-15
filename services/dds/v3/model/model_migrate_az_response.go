package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// MigrateAzResponse Response Object
type MigrateAzResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MigrateAzResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateAzResponse struct{}"
	}

	return strings.Join([]string{"MigrateAzResponse", string(data)}, " ")
}
