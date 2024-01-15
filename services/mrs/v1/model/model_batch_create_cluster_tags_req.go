package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchCreateClusterTagsReq struct {
	Action BatchCreateClusterTagsReqAction `json:"action"`

	Tags []Tag `json:"tags"`
}

func (o BatchCreateClusterTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateClusterTagsReq struct{}"
	}

	return strings.Join([]string{"BatchCreateClusterTagsReq", string(data)}, " ")
}

type BatchCreateClusterTagsReqAction struct {
	value string
}

type BatchCreateClusterTagsReqActionEnum struct {
	CREATE BatchCreateClusterTagsReqAction
}

func GetBatchCreateClusterTagsReqActionEnum() BatchCreateClusterTagsReqActionEnum {
	return BatchCreateClusterTagsReqActionEnum{
		CREATE: BatchCreateClusterTagsReqAction{
			value: "create",
		},
	}
}

func (c BatchCreateClusterTagsReqAction) Value() string {
	return c.value
}

func (c BatchCreateClusterTagsReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateClusterTagsReqAction) UnmarshalJSON(b []byte) error {
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
