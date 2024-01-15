package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7policiesResponse struct {
	L7policy       *L7policyResp `json:"l7policy,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateL7policiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7policiesResponse struct{}"
	}

	return strings.Join([]string{"UpdateL7policiesResponse", string(data)}, " ")
}
