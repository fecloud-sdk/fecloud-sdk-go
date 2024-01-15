package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeletePublicipTagResponse Response Object
type DeletePublicipTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePublicipTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicipTagResponse struct{}"
	}

	return strings.Join([]string{"DeletePublicipTagResponse", string(data)}, " ")
}
