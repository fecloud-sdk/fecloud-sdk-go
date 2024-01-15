package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateConnectorRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreateConnectorReq `json:"body,omitempty"`
}

func (o CreateConnectorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectorRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectorRequest", string(data)}, " ")
}
