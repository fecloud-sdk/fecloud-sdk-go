package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DecryptDataRequest struct {
	Body *DecryptDataRequestBody `json:"body,omitempty"`
}

func (o DecryptDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDataRequest struct{}"
	}

	return strings.Join([]string{"DecryptDataRequest", string(data)}, " ")
}
