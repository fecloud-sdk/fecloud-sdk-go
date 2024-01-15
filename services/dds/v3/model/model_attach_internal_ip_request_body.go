package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttachInternalIpRequestBody struct {
	NodeId string `json:"node_id"`

	NewIp string `json:"new_ip"`
}

func (o AttachInternalIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternalIpRequestBody struct{}"
	}

	return strings.Join([]string{"AttachInternalIpRequestBody", string(data)}, " ")
}
