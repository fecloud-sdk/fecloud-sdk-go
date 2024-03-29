package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MigrateServerRequest struct {
	ServerId string `json:"server_id"`

	Body *MigrateServerRequestBody `json:"body,omitempty"`
}

func (o MigrateServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateServerRequest struct{}"
	}

	return strings.Join([]string{"MigrateServerRequest", string(data)}, " ")
}
