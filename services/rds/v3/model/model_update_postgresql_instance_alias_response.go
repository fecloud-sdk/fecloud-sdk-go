package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePostgresqlInstanceAliasResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePostgresqlInstanceAliasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostgresqlInstanceAliasResponse struct{}"
	}

	return strings.Join([]string{"UpdatePostgresqlInstanceAliasResponse", string(data)}, " ")
}
