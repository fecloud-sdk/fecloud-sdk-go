package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowDiagnosisTaskDetailsResponse struct {
	AbnormalItemSum *int32 `json:"abnormal_item_sum,omitempty"`

	FailedItemSum *int32 `json:"failed_item_sum,omitempty"`

	DiagnosisNodeReportList *[]DiagnosisNodeReport `json:"diagnosis_node_report_list,omitempty"`
	HttpStatusCode          int                    `json:"-"`
}

func (o ShowDiagnosisTaskDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisTaskDetailsResponse", string(data)}, " ")
}
