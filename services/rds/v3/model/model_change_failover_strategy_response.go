package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeFailoverStrategyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeFailoverStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFailoverStrategyResponse struct{}"
	}

	return strings.Join([]string{"ChangeFailoverStrategyResponse", string(data)}, " ")
}
