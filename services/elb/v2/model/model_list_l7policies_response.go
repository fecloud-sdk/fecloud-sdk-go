package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListL7policiesResponse struct {
	L7policies     *[]L7policyResp `json:"l7policies,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListL7policiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7policiesResponse struct{}"
	}

	return strings.Join([]string{"ListL7policiesResponse", string(data)}, " ")
}
