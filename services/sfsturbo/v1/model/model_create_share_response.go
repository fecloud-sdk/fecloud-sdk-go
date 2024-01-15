package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateShareResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareResponse struct{}"
	}

	return strings.Join([]string{"CreateShareResponse", string(data)}, " ")
}
