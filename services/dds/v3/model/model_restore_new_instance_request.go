package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// RestoreNewInstanceRequest Request Object
type RestoreNewInstanceRequest struct {
	Body *RestoreNewInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreNewInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreNewInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreNewInstanceRequest", string(data)}, " ")
}
