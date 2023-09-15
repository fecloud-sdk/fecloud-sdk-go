package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteSingleInstanceResponse Response Object
type DeleteSingleInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSingleInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSingleInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteSingleInstanceResponse", string(data)}, " ")
}
