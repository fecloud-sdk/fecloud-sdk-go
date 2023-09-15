package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CompareConfigurationRequest Request Object
type CompareConfigurationRequest struct {
	Body *DiffConfigurationRequest `json:"body,omitempty"`
}

func (o CompareConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CompareConfigurationRequest", string(data)}, " ")
}
