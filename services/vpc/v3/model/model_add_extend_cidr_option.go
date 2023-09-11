package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddExtendCidrOption struct {
	ExtendCidrs []string `json:"extend_cidrs"`
}

func (o AddExtendCidrOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddExtendCidrOption struct{}"
	}

	return strings.Join([]string{"AddExtendCidrOption", string(data)}, " ")
}
