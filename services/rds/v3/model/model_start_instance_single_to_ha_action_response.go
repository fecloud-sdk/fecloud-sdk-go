package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StartInstanceSingleToHaActionResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceSingleToHaActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceSingleToHaActionResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceSingleToHaActionResponse", string(data)}, " ")
}
