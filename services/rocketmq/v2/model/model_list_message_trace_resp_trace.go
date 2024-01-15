package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListMessageTraceRespTrace struct {
	Success *bool `json:"success,omitempty"`

	TraceType *ListMessageTraceRespTraceTraceType `json:"trace_type,omitempty"`

	Timestamp float32 `json:"timestamp,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	CostTime float32 `json:"cost_time,omitempty"`

	RequestId *string `json:"request_id,omitempty"`

	ConsumeStatus float32 `json:"consume_status,omitempty"`

	Topic *string `json:"topic,omitempty"`

	MsgId *string `json:"msg_id,omitempty"`

	OffsetMsgId *string `json:"offset_msg_id,omitempty"`

	Tags *string `json:"tags,omitempty"`

	Keys *string `json:"keys,omitempty"`

	StoreHost *string `json:"store_host,omitempty"`

	ClientHost *string `json:"client_host,omitempty"`

	RetryTimes *string `json:"retry_times,omitempty"`

	BodyLength float32 `json:"body_length,omitempty"`

	MsgType *ListMessageTraceRespTraceMsgType `json:"msg_type,omitempty"`

	TransactionState *ListMessageTraceRespTraceTransactionState `json:"transaction_state,omitempty"`

	TransactionId *string `json:"transaction_id,omitempty"`

	FromTransactionCheck *bool `json:"from_transaction_check,omitempty"`
}

func (o ListMessageTraceRespTrace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTraceRespTrace struct{}"
	}

	return strings.Join([]string{"ListMessageTraceRespTrace", string(data)}, " ")
}

type ListMessageTraceRespTraceTraceType struct {
	value string
}

type ListMessageTraceRespTraceTraceTypeEnum struct {
	PUB             ListMessageTraceRespTraceTraceType
	SUB             ListMessageTraceRespTraceTraceType
	END_TRANSACTION ListMessageTraceRespTraceTraceType
}

func GetListMessageTraceRespTraceTraceTypeEnum() ListMessageTraceRespTraceTraceTypeEnum {
	return ListMessageTraceRespTraceTraceTypeEnum{
		PUB: ListMessageTraceRespTraceTraceType{
			value: "Pub",
		},
		SUB: ListMessageTraceRespTraceTraceType{
			value: "Sub",
		},
		END_TRANSACTION: ListMessageTraceRespTraceTraceType{
			value: "EndTransaction",
		},
	}
}

func (c ListMessageTraceRespTraceTraceType) Value() string {
	return c.value
}

func (c ListMessageTraceRespTraceTraceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessageTraceRespTraceTraceType) UnmarshalJSON(b []byte) error {
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

type ListMessageTraceRespTraceMsgType struct {
	value string
}

type ListMessageTraceRespTraceMsgTypeEnum struct {
	NORMAL_MSG       ListMessageTraceRespTraceMsgType
	TRANS_MSG_HALF   ListMessageTraceRespTraceMsgType
	TRANS_MSG_COMMIT ListMessageTraceRespTraceMsgType
	DELAY_MSG        ListMessageTraceRespTraceMsgType
}

func GetListMessageTraceRespTraceMsgTypeEnum() ListMessageTraceRespTraceMsgTypeEnum {
	return ListMessageTraceRespTraceMsgTypeEnum{
		NORMAL_MSG: ListMessageTraceRespTraceMsgType{
			value: "Normal_Msg",
		},
		TRANS_MSG_HALF: ListMessageTraceRespTraceMsgType{
			value: "Trans_Msg_Half",
		},
		TRANS_MSG_COMMIT: ListMessageTraceRespTraceMsgType{
			value: "Trans_msg_Commit",
		},
		DELAY_MSG: ListMessageTraceRespTraceMsgType{
			value: "Delay_Msg",
		},
	}
}

func (c ListMessageTraceRespTraceMsgType) Value() string {
	return c.value
}

func (c ListMessageTraceRespTraceMsgType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessageTraceRespTraceMsgType) UnmarshalJSON(b []byte) error {
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

type ListMessageTraceRespTraceTransactionState struct {
	value string
}

type ListMessageTraceRespTraceTransactionStateEnum struct {
	COMMIT_MESSAGE   ListMessageTraceRespTraceTransactionState
	ROLLBACK_MESSAGE ListMessageTraceRespTraceTransactionState
	UNKNOW           ListMessageTraceRespTraceTransactionState
}

func GetListMessageTraceRespTraceTransactionStateEnum() ListMessageTraceRespTraceTransactionStateEnum {
	return ListMessageTraceRespTraceTransactionStateEnum{
		COMMIT_MESSAGE: ListMessageTraceRespTraceTransactionState{
			value: "COMMIT_MESSAGE",
		},
		ROLLBACK_MESSAGE: ListMessageTraceRespTraceTransactionState{
			value: "ROLLBACK_MESSAGE",
		},
		UNKNOW: ListMessageTraceRespTraceTransactionState{
			value: "UNKNOW",
		},
	}
}

func (c ListMessageTraceRespTraceTransactionState) Value() string {
	return c.value
}

func (c ListMessageTraceRespTraceTransactionState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessageTraceRespTraceTransactionState) UnmarshalJSON(b []byte) error {
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
