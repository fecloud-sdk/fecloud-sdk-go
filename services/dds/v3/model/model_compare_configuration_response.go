package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CompareConfigurationResponse struct {
	Differences    *[]DiffDetails `json:"differences,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CompareConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CompareConfigurationResponse", string(data)}, " ")
}
