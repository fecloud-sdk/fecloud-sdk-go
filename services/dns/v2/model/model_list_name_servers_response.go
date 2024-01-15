package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListNameServersResponse struct {
	Nameservers    *[]NameServersResp `json:"nameservers,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListNameServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNameServersResponse struct{}"
	}

	return strings.Join([]string{"ListNameServersResponse", string(data)}, " ")
}
