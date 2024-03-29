package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteL7PolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteL7PolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteL7PolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteL7PolicyResponse", string(data)}, " ")
}
