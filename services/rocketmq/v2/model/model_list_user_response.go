package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListUserResponse struct {
	Users *[]User `json:"users,omitempty"`

	Total          float32 `json:"total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserResponse struct{}"
	}

	return strings.Join([]string{"ListUserResponse", string(data)}, " ")
}
