package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEnhancedConnectionRequestBody struct {
	Hosts []EnhancedConnectionHost `json:"hosts"`
}

func (o UpdateEnhancedConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnhancedConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEnhancedConnectionRequestBody", string(data)}, " ")
}
