package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Request Object
type UpdateKeyDescriptionRequest struct {
	Body *UpdateKeyDescriptionRequestBody `json:"body,omitempty"`
}

func (o UpdateKeyDescriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyDescriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyDescriptionRequest", string(data)}, " ")
}
