package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateReplSetNameResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateReplSetNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReplSetNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateReplSetNameResponse", string(data)}, " ")
}
