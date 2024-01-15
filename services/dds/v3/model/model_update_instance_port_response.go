package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstancePortResponse struct {
	JobId *string `json:"job_id,omitempty"`

	Port           *int32 `json:"port,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateInstancePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstancePortResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstancePortResponse", string(data)}, " ")
}
