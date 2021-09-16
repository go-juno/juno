package endpoint

import (
	"context"

	"juno/internal/model"
	model2 "juno/pkg/model"

	"golang.org/x/xerrors"
)

type GetGreetingListRequest struct {
	PageIndex int
	PageSize  int
}

type GetGreetingListResponse struct {
	GreetingList []*model.Greeting
	Total        int64
}

type GetGreetingAllRequest struct {
}

type GetGreetingAllResponse struct {
	GreetingList []*model.Greeting
}

type GetGreetingDetailRequest struct {
	Id uint
}

type GetGreetingDetailResponse struct {
	Greeting *model.Greeting
}

type CreateGreetingRequest struct {
}

type CreateGreetingResponse struct {
}

type UpdateGreetingRequest struct {
	Id uint
}

type UpdateGreetingResponse struct {
}

type DeleteGreetingRequest struct {
	Id uint
}

type DeleteGreetingResponse struct {
}

func (e *Endpoints) GetGreetingListEndpoint(ctx context.Context, request *GetGreetingListRequest) (response *GetGreetingListResponse, err error) {
	greetingList, total, err := e.greeting.GetList(request.PageIndex, request.PageSize)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	response = &GetGreetingListResponse{
		GreetingList: greetingList,
		Total:        total,
	}
	return
}

func (e *Endpoints) GetGreetingAllEndpoint(ctx context.Context, request *GetGreetingAllRequest) (response *GetGreetingAllResponse, err error) {
	greetingList, err := e.greeting.GetAll()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	response = &GetGreetingAllResponse{
		GreetingList: greetingList,
	}
	return
}

func (e *Endpoints) GetGreetingDetailEndpoint(ctx context.Context, request *GetGreetingDetailRequest) (response *GetGreetingDetailResponse, err error) {
	greeting, err := e.greeting.GetDetail(request.Id)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	response = &GetGreetingDetailResponse{
		Greeting: greeting,
	}
	return
}

func (e *Endpoints) CreateGreetingEndpoint(ctx context.Context, request *CreateGreetingRequest) (response *CreateGreetingResponse, err error) {
	greeting := &model.Greeting{}

	err = e.greeting.Create(greeting)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	response = &CreateGreetingResponse{}
	return
}

func (e *Endpoints) UpdateGreetingEndpoint(ctx context.Context, request *UpdateGreetingRequest) (response *UpdateGreetingResponse, err error) {

	greeting := &model.Greeting{
		Model: model2.Model{
			Id: request.Id,
		},
	}
	err = e.greeting.Update(greeting)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	response = &UpdateGreetingResponse{}
	return
}

func (e *Endpoints) DeleteGreetingEndpoint(ctx context.Context, request *DeleteGreetingRequest) (response *DeleteGreetingResponse, err error) {

	err = e.greeting.Delete(request.Id)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	response = &DeleteGreetingResponse{}
	return
}
