package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddVpcExtendCidrRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	Vpc *AddExtendCidrOption `json:"vpc"`
}

func (o AddVpcExtendCidrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVpcExtendCidrRequestBody struct{}"
	}

	return strings.Join([]string{"AddVpcExtendCidrRequestBody", string(data)}, " ")
}
