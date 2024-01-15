package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteL7policyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteL7policyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteL7policyResponse struct{}"
	}

	return strings.Join([]string{"DeleteL7policyResponse", string(data)}, " ")
}
