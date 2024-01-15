package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EnhancedConnectionHost struct {
	Name *string `json:"name,omitempty"`

	Ip *string `json:"ip,omitempty"`
}

func (o EnhancedConnectionHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnhancedConnectionHost struct{}"
	}

	return strings.Join([]string{"EnhancedConnectionHost", string(data)}, " ")
}
