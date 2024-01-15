package v1

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/sfsturbo/v1/model"
)

type BatchAddSharedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddSharedTagsInvoker) Invoke() (*model.BatchAddSharedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddSharedTagsResponse), nil
	}
}

type ChangeSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeSecurityGroupInvoker) Invoke() (*model.ChangeSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeSecurityGroupResponse), nil
	}
}

type ChangeShareNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeShareNameInvoker) Invoke() (*model.ChangeShareNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeShareNameResponse), nil
	}
}

type CreateShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateShareInvoker) Invoke() (*model.CreateShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateShareResponse), nil
	}
}

type CreateSharedTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSharedTagInvoker) Invoke() (*model.CreateSharedTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSharedTagResponse), nil
	}
}

type DeleteShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteShareInvoker) Invoke() (*model.DeleteShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteShareResponse), nil
	}
}

type DeleteSharedTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSharedTagInvoker) Invoke() (*model.DeleteSharedTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSharedTagResponse), nil
	}
}

type ExpandShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandShareInvoker) Invoke() (*model.ExpandShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandShareResponse), nil
	}
}

type ListSharedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSharedTagsInvoker) Invoke() (*model.ListSharedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSharedTagsResponse), nil
	}
}

type ListSharesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSharesInvoker) Invoke() (*model.ListSharesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSharesResponse), nil
	}
}

type ShowShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowShareInvoker) Invoke() (*model.ShowShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowShareResponse), nil
	}
}

type ShowSharedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSharedTagsInvoker) Invoke() (*model.ShowSharedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSharedTagsResponse), nil
	}
}
