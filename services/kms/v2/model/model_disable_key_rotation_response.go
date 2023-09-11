package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DisableKeyRotationResponse Response Object
type DisableKeyRotationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableKeyRotationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableKeyRotationResponse struct{}"
	}

	return strings.Join([]string{"DisableKeyRotationResponse", string(data)}, " ")
}
