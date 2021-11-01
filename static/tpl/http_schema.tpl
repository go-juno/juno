package schema

import "juno/internal/endpoint"

type GetGreetingList struct {
	PageIndex int `form:"page_index" json:"page_index" binding:"required,min=1"`
	PageSize  int `form:"page_size" json:"page_size" binding:"required,min=1"`
}

func (s *GetGreetingList) Transform() *endpoint.GetGreetingListRequest {
	req := &endpoint.GetGreetingListRequest{
		PageIndex: s.PageIndex,
		PageSize:  s.PageSize,
	}
	return req
}

type GetGreetingAll struct {
}

func (s *GetGreetingAll) Transform() *endpoint.GetGreetingAllRequest {
	req := &endpoint.GetGreetingAllRequest{}
	return req
}

type GetGreetingDetail struct {
	Id uint `form:"id" json:"id" binding:"required,min=1"`
}

func (s *GetGreetingDetail) Transform() *endpoint.GetGreetingDetailRequest {
	req := &endpoint.GetGreetingDetailRequest{
		Id: s.Id,
	}
	return req
}

type CreateGreeting struct {
}

func (s *CreateGreeting) Transform() *endpoint.CreateGreetingRequest {
	req := &endpoint.CreateGreetingRequest{}
	return req
}

type UpdateGreeting struct {
	Id uint `form:"id" json:"id" binding:"required,min=1"`
}

func (s *UpdateGreeting) Transform() *endpoint.UpdateGreetingRequest {
	req := &endpoint.UpdateGreetingRequest{
		Id: s.Id,
	}
	return req
}

type DeleteGreeting struct {
	Id uint `form:"id" json:"id" binding:"required,min=1"`
}

func (s *DeleteGreeting) Transform() *endpoint.DeleteGreetingRequest {
	req := &endpoint.DeleteGreetingRequest{
		Id: s.Id,
	}
	return req
}
