package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSqlResultWithJobResponse struct {
	SqlResults     *interface{} `json:"sql_results,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowSqlResultWithJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlResultWithJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlResultWithJobResponse", string(data)}, " ")
}
