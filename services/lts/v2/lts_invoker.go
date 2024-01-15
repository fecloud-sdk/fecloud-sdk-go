package v2

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/lts/v2/model"
)

type CreateLogDumpObsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLogDumpObsInvoker) Invoke() (*model.CreateLogDumpObsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogDumpObsResponse), nil
	}
}

type CreateLogGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLogGroupInvoker) Invoke() (*model.CreateLogGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogGroupResponse), nil
	}
}

type CreateLogStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLogStreamInvoker) Invoke() (*model.CreateLogStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogStreamResponse), nil
	}
}

type DeleteLogGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogGroupInvoker) Invoke() (*model.DeleteLogGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogGroupResponse), nil
	}
}

type DeleteLogStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogStreamInvoker) Invoke() (*model.DeleteLogStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogStreamResponse), nil
	}
}

type ListLogGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogGroupsInvoker) Invoke() (*model.ListLogGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogGroupsResponse), nil
	}
}

type ListLogStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogStreamInvoker) Invoke() (*model.ListLogStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogStreamResponse), nil
	}
}

type ListLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogsInvoker) Invoke() (*model.ListLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogsResponse), nil
	}
}
