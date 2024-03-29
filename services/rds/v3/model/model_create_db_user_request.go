package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDbUserRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *UserForCreation `json:"body,omitempty"`
}

func (o CreateDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbUserRequest struct{}"
	}

	return strings.Join([]string{"CreateDbUserRequest", string(data)}, " ")
}
