package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddVpcExtendCidrRequest struct {
	VpcId string `json:"vpc_id"`

	Body *AddVpcExtendCidrRequestBody `json:"body,omitempty"`
}

func (o AddVpcExtendCidrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVpcExtendCidrRequest struct{}"
	}

	return strings.Join([]string{"AddVpcExtendCidrRequest", string(data)}, " ")
}
