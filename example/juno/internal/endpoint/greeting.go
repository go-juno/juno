package endpoint

import (
	"context"

	"github.com/go-juno/juno/example/juno/internal/model"
	model2 "github.com/go-juno/juno/example/juno/pkg/model"
	"golang.org/x/xerrors"
)

type GetGreetingListRequest struct {
	PageIndex int
	PageSize  int
}

type GetGreetingList struct {
	Id uint
}

type GetGreetingListResponse struct {
	Items []*GetGreetingList
	Total int64
}

type GetGreetingAllRequest struct {
}

type GetGreetingAllResponse struct {
	Id uint
}

type GetGreetingDetailRequest struct {
	Id uint
}

type GetGreetingDetailResponse struct {
	Id uint
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
	items := make([]*GetGreetingList, len(greetingList))
	for index, item := range greetingList {
		items[index] = &GetGreetingList{
			Id: item.Id,
		}

	}
	response = &GetGreetingListResponse{
		Items: items,
		Total: total,
	}
	return
}

func (e *Endpoints) GetGreetingAllEndpoint(ctx context.Context, request *GetGreetingAllRequest) (response []*GetGreetingAllResponse, err error) {
	greetingList, err := e.greeting.GetAll()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	response = make([]*GetGreetingAllResponse, len(greetingList))
	for index, greeting := range greetingList {
		response[index] = &GetGreetingAllResponse{
			Id: greeting.Id,
		}
	}

	return
}

func (e *Endpoints) GetGreetingDetailEndpoint(ctx context.Context, request *GetGreetingDetailRequest) (response *GetGreetingDetailResponse, err error) {
	greeting, err := e.greeting.GetDetail(request.Id)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if greeting != nil {
		response = &GetGreetingDetailResponse{
			Id: greeting.Id,
		}
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
