package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GenerateAuditlogDownloadLinkRequest struct {

	// 审计日志ID列表，限制50条以内。
	Ids []string `json:"ids"`
}

func (o GenerateAuditlogDownloadLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateAuditlogDownloadLinkRequest struct{}"
	}

	return strings.Join([]string{"GenerateAuditlogDownloadLinkRequest", string(data)}, " ")
}
