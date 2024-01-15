package v3

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/cts/v3/model"
)

type CreateTrackerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTrackerInvoker) Invoke() (*model.CreateTrackerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTrackerResponse), nil
	}
}

type DeleteTrackerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTrackerInvoker) Invoke() (*model.DeleteTrackerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTrackerResponse), nil
	}
}

type ListTracesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTracesInvoker) Invoke() (*model.ListTracesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTracesResponse), nil
	}
}

type ListTrackersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrackersInvoker) Invoke() (*model.ListTrackersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrackersResponse), nil
	}
}

type UpdateTrackerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTrackerInvoker) Invoke() (*model.UpdateTrackerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTrackerResponse), nil
	}
}
