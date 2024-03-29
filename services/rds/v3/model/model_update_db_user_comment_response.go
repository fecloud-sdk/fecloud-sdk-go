package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateDbUserCommentResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDbUserCommentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbUserCommentResponse struct{}"
	}

	return strings.Join([]string{"UpdateDbUserCommentResponse", string(data)}, " ")
}
