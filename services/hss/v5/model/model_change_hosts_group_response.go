package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeHostsGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeHostsGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeHostsGroupResponse struct{}"
	}

	return strings.Join([]string{"ChangeHostsGroupResponse", string(data)}, " ")
}
