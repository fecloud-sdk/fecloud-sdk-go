package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateIpListRequest struct {
	IpgroupId string `json:"ipgroup_id"`

	Body *UpdateIpListRequestBody `json:"body,omitempty"`
}

func (o UpdateIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpListRequest", string(data)}, " ")
}
