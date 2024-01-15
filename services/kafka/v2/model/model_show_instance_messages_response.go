package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceMessagesResponse struct {
	Messages *[]MessagesEntity `json:"messages,omitempty"`

	Total *int64 `json:"total,omitempty"`

	Size           *int64 `json:"size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMessagesResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceMessagesResponse", string(data)}, " ")
}
