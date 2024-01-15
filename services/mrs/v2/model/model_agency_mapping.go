package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AgencyMapping struct {
	Agency string `json:"agency"`

	IdentifierType string `json:"identifier_type"`

	Identifiers []string `json:"identifiers"`

	AgencyId string `json:"agency_id"`
}

func (o AgencyMapping) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyMapping struct{}"
	}

	return strings.Join([]string{"AgencyMapping", string(data)}, " ")
}
