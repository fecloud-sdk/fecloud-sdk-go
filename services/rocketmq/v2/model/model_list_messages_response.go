package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListMessagesResponse struct {
	Messages *[]Message `json:"messages,omitempty"`

	Total          float32 `json:"total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessagesResponse struct{}"
	}

	return strings.Join([]string{"ListMessagesResponse", string(data)}, " ")
}
