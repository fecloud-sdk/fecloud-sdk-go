package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateDiagnosisTaskResponse Response Object
type CreateDiagnosisTaskResponse struct {

	// 报告ID
	ReportId       *string `json:"report_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDiagnosisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnosisTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateDiagnosisTaskResponse", string(data)}, " ")
}
