package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ReinstallServerWithCloudInitResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ReinstallServerWithCloudInitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerWithCloudInitResponse struct{}"
	}

	return strings.Join([]string{"ReinstallServerWithCloudInitResponse", string(data)}, " ")
}
