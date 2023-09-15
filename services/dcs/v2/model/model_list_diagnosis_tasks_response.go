package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListDiagnosisTasksResponse Response Object
type ListDiagnosisTasksResponse struct {

	// 诊断报告列表
	DiagnosisReportList *[]DiagnosisReportInfo `json:"diagnosis_report_list,omitempty"`

	// 诊断报告总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDiagnosisTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnosisTasksResponse struct{}"
	}

	return strings.Join([]string{"ListDiagnosisTasksResponse", string(data)}, " ")
}