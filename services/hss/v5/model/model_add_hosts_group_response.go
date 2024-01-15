package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddHostsGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddHostsGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHostsGroupResponse struct{}"
	}

	return strings.Join([]string{"AddHostsGroupResponse", string(data)}, " ")
}
