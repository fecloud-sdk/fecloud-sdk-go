package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CheckPasswordResponse Response Object
type CheckPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPasswordResponse struct{}"
	}

	return strings.Join([]string{"CheckPasswordResponse", string(data)}, " ")
}
