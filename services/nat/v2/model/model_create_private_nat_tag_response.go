package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePrivateNatTagResponse Response Object
type CreatePrivateNatTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePrivateNatTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatTagResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatTagResponse", string(data)}, " ")
}
