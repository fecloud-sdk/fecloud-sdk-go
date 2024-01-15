package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetBackupPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetBackupPolicyResponse", string(data)}, " ")
}
