package v1

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/mrs/v1/model"
)

type BatchCreateClusterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateClusterTagsInvoker) Invoke() (*model.BatchCreateClusterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateClusterTagsResponse), nil
	}
}

type BatchDeleteClusterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteClusterTagsInvoker) Invoke() (*model.BatchDeleteClusterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteClusterTagsResponse), nil
	}
}

type CreateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterInvoker) Invoke() (*model.CreateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterResponse), nil
	}
}

type CreateClusterTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterTagInvoker) Invoke() (*model.CreateClusterTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterTagResponse), nil
	}
}

type CreateScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScalingPolicyInvoker) Invoke() (*model.CreateScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScalingPolicyResponse), nil
	}
}

type DeleteClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterInvoker) Invoke() (*model.DeleteClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterResponse), nil
	}
}

type DeleteClusterTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterTagInvoker) Invoke() (*model.DeleteClusterTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterTagResponse), nil
	}
}

type ListAllTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllTagsInvoker) Invoke() (*model.ListAllTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllTagsResponse), nil
	}
}

type ListClusterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterTagsInvoker) Invoke() (*model.ListClusterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterTagsResponse), nil
	}
}

type ListClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClustersInvoker) Invoke() (*model.ListClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClustersResponse), nil
	}
}

type ListClustersByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClustersByTagsInvoker) Invoke() (*model.ListClustersByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClustersByTagsResponse), nil
	}
}

type ListHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostsInvoker) Invoke() (*model.ListHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostsResponse), nil
	}
}

type ShowClusterDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterDetailsInvoker) Invoke() (*model.ShowClusterDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterDetailsResponse), nil
	}
}

type UpdateClusterScalingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterScalingInvoker) Invoke() (*model.UpdateClusterScalingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterScalingResponse), nil
	}
}
