package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type DeleteTrackerRequest struct {
	TrackerName *DeleteTrackerRequestTrackerName `json:"tracker_name,omitempty"`
}

func (o DeleteTrackerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrackerRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrackerRequest", string(data)}, " ")
}

type DeleteTrackerRequestTrackerName struct {
	value string
}

type DeleteTrackerRequestTrackerNameEnum struct {
	SYSTEM DeleteTrackerRequestTrackerName
}

func GetDeleteTrackerRequestTrackerNameEnum() DeleteTrackerRequestTrackerNameEnum {
	return DeleteTrackerRequestTrackerNameEnum{
		SYSTEM: DeleteTrackerRequestTrackerName{
			value: "system",
		},
	}
}

func (c DeleteTrackerRequestTrackerName) Value() string {
	return c.value
}

func (c DeleteTrackerRequestTrackerName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteTrackerRequestTrackerName) UnmarshalJSON(b []byte) error {
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
