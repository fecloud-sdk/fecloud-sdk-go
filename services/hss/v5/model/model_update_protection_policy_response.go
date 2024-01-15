package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateProtectionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateProtectionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectionPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateProtectionPolicyResponse", string(data)}, " ")
}
