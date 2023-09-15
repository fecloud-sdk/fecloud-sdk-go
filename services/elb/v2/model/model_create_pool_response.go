package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePoolResponse Response Object
type CreatePoolResponse struct {
	Pool           *PoolResp `json:"pool,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreatePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolResponse struct{}"
	}

	return strings.Join([]string{"CreatePoolResponse", string(data)}, " ")
}
