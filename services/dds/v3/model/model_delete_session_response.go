package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSessionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSessionResponse struct{}"
	}

	return strings.Join([]string{"DeleteSessionResponse", string(data)}, " ")
}
