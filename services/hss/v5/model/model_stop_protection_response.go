package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StopProtectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopProtectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopProtectionResponse struct{}"
	}

	return strings.Join([]string{"StopProtectionResponse", string(data)}, " ")
}
