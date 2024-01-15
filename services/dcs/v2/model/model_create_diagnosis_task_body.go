package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDiagnosisTaskBody struct {
	BeginTime string `json:"begin_time"`

	EndTime string `json:"end_time"`

	NodeIpList *[]string `json:"node_ip_list,omitempty"`
}

func (o CreateDiagnosisTaskBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnosisTaskBody struct{}"
	}

	return strings.Join([]string{"CreateDiagnosisTaskBody", string(data)}, " ")
}
