package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListEventDetailResponse struct {
	EventName *string `json:"event_name,omitempty"`

	EventType *ListEventDetailResponseEventType `json:"event_type,omitempty"`

	EventUsers *[]string `json:"event_users,omitempty"`

	EventSources *[]string `json:"event_sources,omitempty"`

	EventInfo *[]EventInfoDetail `json:"event_info,omitempty"`

	MetaData       *TotalMetaData `json:"meta_data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEventDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventDetailResponse struct{}"
	}

	return strings.Join([]string{"ListEventDetailResponse", string(data)}, " ")
}

type ListEventDetailResponseEventType struct {
	value string
}

type ListEventDetailResponseEventTypeEnum struct {
	EVENT_SYS    ListEventDetailResponseEventType
	EVENT_CUSTOM ListEventDetailResponseEventType
}

func GetListEventDetailResponseEventTypeEnum() ListEventDetailResponseEventTypeEnum {
	return ListEventDetailResponseEventTypeEnum{
		EVENT_SYS: ListEventDetailResponseEventType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: ListEventDetailResponseEventType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c ListEventDetailResponseEventType) Value() string {
	return c.value
}

func (c ListEventDetailResponseEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventDetailResponseEventType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
