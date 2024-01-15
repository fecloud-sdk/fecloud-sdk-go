package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeVulStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeVulStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVulStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeVulStatusResponse", string(data)}, " ")
}
