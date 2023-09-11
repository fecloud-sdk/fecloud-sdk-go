package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdatePrivateNatResponse Response Object
type UpdatePrivateNatResponse struct {
	Gateway *PrivateNat `json:"gateway,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePrivateNatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateNatResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivateNatResponse", string(data)}, " ")
}
