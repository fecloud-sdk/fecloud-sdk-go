package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateKeyRotationIntervalResponse Response Object
type UpdateKeyRotationIntervalResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateKeyRotationIntervalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyRotationIntervalResponse struct{}"
	}

	return strings.Join([]string{"UpdateKeyRotationIntervalResponse", string(data)}, " ")
}
