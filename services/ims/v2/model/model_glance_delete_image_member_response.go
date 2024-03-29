package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GlanceDeleteImageMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceDeleteImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageMemberResponse struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageMemberResponse", string(data)}, " ")
}
