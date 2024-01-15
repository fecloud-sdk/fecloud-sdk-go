package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSqlResultRequest struct {
	ClusterId string `json:"cluster_id"`

	SqlId string `json:"sql_id"`
}

func (o ShowSqlResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlResultRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlResultRequest", string(data)}, " ")
}
