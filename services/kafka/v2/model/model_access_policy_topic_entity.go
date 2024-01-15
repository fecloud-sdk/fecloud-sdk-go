package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AccessPolicyTopicEntity struct {
	Name string `json:"name"`

	Policies []AccessPolicyEntity `json:"policies"`
}

func (o AccessPolicyTopicEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicyTopicEntity struct{}"
	}

	return strings.Join([]string{"AccessPolicyTopicEntity", string(data)}, " ")
}
