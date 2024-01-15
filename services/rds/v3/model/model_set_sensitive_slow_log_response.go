package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetSensitiveSlowLogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetSensitiveSlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSensitiveSlowLogResponse struct{}"
	}

	return strings.Join([]string{"SetSensitiveSlowLogResponse", string(data)}, " ")
}
