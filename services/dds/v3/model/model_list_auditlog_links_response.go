package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAuditlogLinksResponse struct {
	Links          *[]string `json:"links,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAuditlogLinksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditlogLinksResponse struct{}"
	}

	return strings.Join([]string{"ListAuditlogLinksResponse", string(data)}, " ")
}
