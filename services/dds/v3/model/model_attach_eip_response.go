package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttachEipResponse struct {
	JobId *string `json:"job_id,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	NodeName *string `json:"node_name,omitempty"`

	PublicIpId *string `json:"public_ip_id,omitempty"`

	PublicIp       *string `json:"public_ip,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipResponse struct{}"
	}

	return strings.Join([]string{"AttachEipResponse", string(data)}, " ")
}
