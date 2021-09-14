package service

import (
	"context"
	"log"

	"github.com/go-juno/juno/example/juno/api/grpc/protos"
	"github.com/go-juno/juno/example/juno/internal/endpoint"
)

type GreetingServer struct {
	endpoints *endpoint.Endpoints
}

func (r *GreetingServer) GetList(ctx context.Context, in *protos.GetGreetingListParam) (out *protos.GetGreetingListReply, err error) {
	request := &endpoint.GetGreetingListRequest{}
	res, err := r.endpoints.GetGreetingListEndpoint(ctx, request)
	if err != nil {
		log.Printf("err:%+v", err)
		return
	}
	item := make([]*protos.GetGreetingListReply_List, len(res.GreetingList))
	for index := range res.GreetingList {
		item[index] = &protos.GetGreetingListReply_List{}
	}
	out = &protos.GetGreetingListReply{
		Item:  item,
		Total: res.Total,
	}
	return
}

func (r *GreetingServer) GetAll(ctx context.Context, in *protos.GetGreetingAllParam) (out *protos.GetGreetingAllReply, err error) {
	request := &endpoint.GetGreetingAllRequest{}
	res, err := r.endpoints.GetGreetingAllEndpoint(ctx, request)
	if err != nil {
		log.Printf("err:%+v", err)
		return
	}
	item := make([]*protos.GetGreetingAllReply_List, len(res))
	for index := range res {
		item[index] = &protos.GetGreetingAllReply_List{}
	}
	out = &protos.GetGreetingAllReply{
		Item: item,
	}
	return
}

func (r *GreetingServer) GetDetail(ctx context.Context, in *protos.GetGreetingDetailParam) (out *protos.GetGreetingDetailReply, err error) {
	request := &endpoint.GetGreetingDetailRequest{
		Id: uint(in.Id),
	}
	_, err = r.endpoints.GetGreetingDetailEndpoint(ctx, request)
	if err != nil {
		log.Printf("err:%+v", err)
		return
	}

	out = &protos.GetGreetingDetailReply{}
	return
}

func (r *GreetingServer) Create(ctx context.Context, in *protos.CreateGreetingParam) (out *protos.CreateGreetingReply, err error) {
	request := &endpoint.CreateGreetingRequest{}
	_, err = r.endpoints.CreateGreetingEndpoint(ctx, request)
	if err != nil {
		log.Printf("err:%+v", err)
		return
	}

	out = &protos.CreateGreetingReply{}
	return
}

func (r *GreetingServer) Update(ctx context.Context, in *protos.UpdateGreetingParam) (out *protos.UpdateGreetingReply, err error) {
	request := &endpoint.UpdateGreetingRequest{}
	_, err = r.endpoints.UpdateGreetingEndpoint(ctx, request)
	if err != nil {
		log.Printf("err:%+v", err)
		return
	}

	out = &protos.UpdateGreetingReply{}
	return
}

func (r *GreetingServer) Delete(ctx context.Context, in *protos.DeleteGreetingParam) (out *protos.DeleteGreetingReply, err error) {
	request := &endpoint.DeleteGreetingRequest{}
	_, err = r.endpoints.DeleteGreetingEndpoint(ctx, request)
	if err != nil {
		log.Printf("err:%+v", err)
		return
	}

	out = &protos.DeleteGreetingReply{}
	return
}

func NewGreetingServer(endpoints *endpoint.Endpoints) *GreetingServer {
	return &GreetingServer{
		endpoints: endpoints,
	}
}
