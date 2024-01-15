package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLoadbalancerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerResponse", string(data)}, " ")
}
