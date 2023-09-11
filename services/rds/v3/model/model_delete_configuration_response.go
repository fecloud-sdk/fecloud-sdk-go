package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteConfigurationResponse Response Object
type DeleteConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigurationResponse struct{}"
	}

	return strings.Join([]string{"DeleteConfigurationResponse", string(data)}, " ")
}
