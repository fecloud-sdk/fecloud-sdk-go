package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteShareResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareResponse struct{}"
	}

	return strings.Join([]string{"DeleteShareResponse", string(data)}, " ")
}
