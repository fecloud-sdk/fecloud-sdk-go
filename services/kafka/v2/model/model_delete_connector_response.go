package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteConnectorResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteConnectorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectorResponse struct{}"
	}

	return strings.Join([]string{"DeleteConnectorResponse", string(data)}, " ")
}
