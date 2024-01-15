package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteListenerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerResponse struct{}"
	}

	return strings.Join([]string{"DeleteListenerResponse", string(data)}, " ")
}
