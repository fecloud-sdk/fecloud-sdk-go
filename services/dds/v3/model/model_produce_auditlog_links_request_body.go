package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ProduceAuditlogLinksRequestBody struct {
	Ids []string `json:"ids"`
}

func (o ProduceAuditlogLinksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProduceAuditlogLinksRequestBody struct{}"
	}

	return strings.Join([]string{"ProduceAuditlogLinksRequestBody", string(data)}, " ")
}
