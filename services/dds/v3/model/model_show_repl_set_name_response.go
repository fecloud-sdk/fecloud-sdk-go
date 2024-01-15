package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowReplSetNameResponse struct {
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowReplSetNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplSetNameResponse struct{}"
	}

	return strings.Join([]string{"ShowReplSetNameResponse", string(data)}, " ")
}
