package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowServerLimitsResponse struct {
	Absolute       *ServerLimits `json:"absolute,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowServerLimitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerLimitsResponse struct{}"
	}

	return strings.Join([]string{"ShowServerLimitsResponse", string(data)}, " ")
}
