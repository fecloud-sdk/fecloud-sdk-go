package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CancelGrantResponse Response Object
type CancelGrantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelGrantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelGrantResponse struct{}"
	}

	return strings.Join([]string{"CancelGrantResponse", string(data)}, " ")
}
