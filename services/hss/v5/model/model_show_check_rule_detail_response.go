package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCheckRuleDetailResponse struct {
	Description *string `json:"description,omitempty"`

	Reference *string `json:"reference,omitempty"`

	Audit *string `json:"audit,omitempty"`

	Remediation *string `json:"remediation,omitempty"`

	CheckInfoList  *[]CheckRuleCheckCaseResponseInfo `json:"check_info_list,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ShowCheckRuleDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckRuleDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCheckRuleDetailResponse", string(data)}, " ")
}
