package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SignResponse struct {
	KeyId *string `json:"key_id,omitempty"`

	Signature      *string `json:"signature,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SignResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignResponse struct{}"
	}

	return strings.Join([]string{"SignResponse", string(data)}, " ")
}
