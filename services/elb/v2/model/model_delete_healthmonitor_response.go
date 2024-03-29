package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteHealthmonitorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHealthmonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthmonitorResponse struct{}"
	}

	return strings.Join([]string{"DeleteHealthmonitorResponse", string(data)}, " ")
}
