package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttachInternalIpResponse struct {
	JobId *string `json:"job_id,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	NewIp          *string `json:"new_ip,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachInternalIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternalIpResponse struct{}"
	}

	return strings.Join([]string{"AttachInternalIpResponse", string(data)}, " ")
}
