package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaCreateKeypairResponse Response Object
type NovaCreateKeypairResponse struct {
	Keypair        *NovaKeypair `json:"keypair,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o NovaCreateKeypairResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateKeypairResponse struct{}"
	}

	return strings.Join([]string{"NovaCreateKeypairResponse", string(data)}, " ")
}
