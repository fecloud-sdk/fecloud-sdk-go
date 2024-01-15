package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateResourceTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceTagResponse", string(data)}, " ")
}
