package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StartProtectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartProtectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartProtectionResponse struct{}"
	}

	return strings.Join([]string{"StartProtectionResponse", string(data)}, " ")
}
