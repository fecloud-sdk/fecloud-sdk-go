package v3

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/vpc/v3/model"
)

type AddVpcExtendCidrInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddVpcExtendCidrInvoker) Invoke() (*model.AddVpcExtendCidrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddVpcExtendCidrResponse), nil
	}
}

type CreateVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpcInvoker) Invoke() (*model.CreateVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpcResponse), nil
	}
}

type DeleteVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpcInvoker) Invoke() (*model.DeleteVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpcResponse), nil
	}
}

type ListVpcsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcsInvoker) Invoke() (*model.ListVpcsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcsResponse), nil
	}
}

type RemoveVpcExtendCidrInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveVpcExtendCidrInvoker) Invoke() (*model.RemoveVpcExtendCidrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveVpcExtendCidrResponse), nil
	}
}

type ShowVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpcInvoker) Invoke() (*model.ShowVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpcResponse), nil
	}
}

type UpdateVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpcInvoker) Invoke() (*model.UpdateVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpcResponse), nil
	}
}
