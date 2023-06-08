package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Request Object
type CreateDatakeyWithoutPlaintextRequest struct {
	Body *CreateDatakeyRequestBody `json:"body,omitempty"`
}

func (o CreateDatakeyWithoutPlaintextRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatakeyWithoutPlaintextRequest struct{}"
	}

	return strings.Join([]string{"CreateDatakeyWithoutPlaintextRequest", string(data)}, " ")
}
