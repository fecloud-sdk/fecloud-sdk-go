package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DecryptDatakeyRequest Request Object
type DecryptDatakeyRequest struct {
	Body *DecryptDatakeyRequestBody `json:"body,omitempty"`
}

func (o DecryptDatakeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDatakeyRequest struct{}"
	}

	return strings.Join([]string{"DecryptDatakeyRequest", string(data)}, " ")
}
