package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePrivateNatResponse Response Object
type CreatePrivateNatResponse struct {
	Gateway *PrivateNat `json:"gateway,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrivateNatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatResponse", string(data)}, " ")
}
