package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateIpGroupResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupResponse", string(data)}, " ")
}
