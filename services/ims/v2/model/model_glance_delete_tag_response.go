package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// GlanceDeleteTagResponse Response Object
type GlanceDeleteTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceDeleteTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteTagResponse struct{}"
	}

	return strings.Join([]string{"GlanceDeleteTagResponse", string(data)}, " ")
}
