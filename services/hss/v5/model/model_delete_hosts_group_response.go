package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteHostsGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHostsGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostsGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteHostsGroupResponse", string(data)}, " ")
}
