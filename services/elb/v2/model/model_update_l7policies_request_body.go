package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7policiesRequestBody struct {
	L7policy *UpdateL7policyReq `json:"l7policy"`
}

func (o UpdateL7policiesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7policiesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateL7policiesRequestBody", string(data)}, " ")
}
