package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CheckWeakPasswordResponse struct {
	Weak           *bool `json:"weak,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckWeakPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWeakPasswordResponse struct{}"
	}

	return strings.Join([]string{"CheckWeakPasswordResponse", string(data)}, " ")
}
