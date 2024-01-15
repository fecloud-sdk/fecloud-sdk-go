package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CancelSparkJobResponse struct {
	Msg            *string `json:"msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelSparkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSparkJobResponse struct{}"
	}

	return strings.Join([]string{"CancelSparkJobResponse", string(data)}, " ")
}
