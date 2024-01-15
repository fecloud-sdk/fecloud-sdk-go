package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceByEngineResponse struct {
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceByEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceByEngineResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceByEngineResponse", string(data)}, " ")
}
