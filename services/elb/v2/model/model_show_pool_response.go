package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowPoolResponse Response Object
type ShowPoolResponse struct {
	Pool           *PoolResp `json:"pool,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolResponse struct{}"
	}

	return strings.Join([]string{"ShowPoolResponse", string(data)}, " ")
}
