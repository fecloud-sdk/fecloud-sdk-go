package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateServersNameResponse struct {
	Response       *[]ServerId `json:"response,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o BatchUpdateServersNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateServersNameResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateServersNameResponse", string(data)}, " ")
}
