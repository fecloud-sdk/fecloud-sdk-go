package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Request Object
type CreateDatakeyRequest struct {
	Body *CreateDatakeyRequestBody `json:"body,omitempty"`
}

func (o CreateDatakeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatakeyRequest struct{}"
	}

	return strings.Join([]string{"CreateDatakeyRequest", string(data)}, " ")
}
