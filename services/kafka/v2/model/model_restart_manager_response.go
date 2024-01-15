package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestartManagerResponse struct {
	Result *string `json:"result,omitempty"`

	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartManagerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartManagerResponse struct{}"
	}

	return strings.Join([]string{"RestartManagerResponse", string(data)}, " ")
}
