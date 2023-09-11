package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteAddonInstanceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAddonInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddonInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteAddonInstanceResponse", string(data)}, " ")
}
