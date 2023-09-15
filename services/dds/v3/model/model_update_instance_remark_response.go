package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateInstanceRemarkResponse Response Object
type UpdateInstanceRemarkResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceRemarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRemarkResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRemarkResponse", string(data)}, " ")
}
