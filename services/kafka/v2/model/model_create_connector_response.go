package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateConnectorResponse struct {
	JobId *string `json:"job_id,omitempty"`

	ConnectorId    *string `json:"connector_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateConnectorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectorResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectorResponse", string(data)}, " ")
}
