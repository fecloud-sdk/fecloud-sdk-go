package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteBackupFileResponse Response Object
type DeleteBackupFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackupFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupFileResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackupFileResponse", string(data)}, " ")
}
