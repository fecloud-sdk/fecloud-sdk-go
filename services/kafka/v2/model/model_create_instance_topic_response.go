package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceTopicResponse struct {
	Id *string `json:"id,omitempty"`

	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceTopicResponse", string(data)}, " ")
}
