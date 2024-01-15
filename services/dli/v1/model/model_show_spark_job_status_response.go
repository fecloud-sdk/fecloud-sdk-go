package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSparkJobStatusResponse struct {
	Id *string `json:"id,omitempty"`

	State          *string `json:"state,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSparkJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobStatusResponse", string(data)}, " ")
}
