package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CancelEipResponse struct {
	JobId *string `json:"job_id,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	NodeName       *string `json:"node_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelEipResponse struct{}"
	}

	return strings.Join([]string{"CancelEipResponse", string(data)}, " ")
}
