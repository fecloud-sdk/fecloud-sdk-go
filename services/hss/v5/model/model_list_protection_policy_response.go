package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProtectionPolicyResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]ProtectionPolicyInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListProtectionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListProtectionPolicyResponse", string(data)}, " ")
}
