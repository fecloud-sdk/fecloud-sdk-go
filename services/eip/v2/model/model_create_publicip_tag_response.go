package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePublicipTagResponse Response Object
type CreatePublicipTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePublicipTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicipTagResponse struct{}"
	}

	return strings.Join([]string{"CreatePublicipTagResponse", string(data)}, " ")
}
