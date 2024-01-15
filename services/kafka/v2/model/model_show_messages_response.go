package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowMessagesResponse struct {
	Messages *[]ShowMessagesRespMessages `json:"messages,omitempty"`

	MessagesCount *int32 `json:"messages_count,omitempty"`

	OffsetsCount *int32 `json:"offsets_count,omitempty"`

	Offset         *int32 `json:"offset,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessagesResponse struct{}"
	}

	return strings.Join([]string{"ShowMessagesResponse", string(data)}, " ")
}
