package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateBackupPolicyInfoResponse struct {
	ErrorCode *int32 `json:"error_code,omitempty"`

	ErrorDescription *string `json:"error_description,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o UpdateBackupPolicyInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupPolicyInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateBackupPolicyInfoResponse", string(data)}, " ")
}
