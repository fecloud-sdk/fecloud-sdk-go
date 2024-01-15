package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLogDumpObsResponse struct {
	LogDumpObsId   *string `json:"log_dump_obs_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogDumpObsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogDumpObsResponse struct{}"
	}

	return strings.Join([]string{"CreateLogDumpObsResponse", string(data)}, " ")
}
