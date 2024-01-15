package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSinkTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSinkTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSinkTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteSinkTaskResponse", string(data)}, " ")
}
