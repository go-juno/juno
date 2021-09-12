package serialize

import (
	"juno/internal/endpoint"
)

type GetGreetingList struct {
	Id uint `json:"id"`
}

func GetGreetingListTransform(e *endpoint.GetGreetingListResponse) (list *List) {
	items := make([]*GetGreetingList, len(e.GreetingList))
	list = &List{
		Total: e.Total,
	}
	for index, greeting := range e.GreetingList {
		items[index] = &GetGreetingList{
			Id: greeting.Id,
		}
	}
	list.Items = items
	return
}

type GetGreetingAll struct {
	Id uint `json:"id"`
}

func GetGreetingAllTransform(e *endpoint.GetGreetingAllResponse) (res []*GetGreetingAll) {
	res = make([]*GetGreetingAll, len(e.GreetingList))
	for index, greeting := range e.GreetingList {
		res[index] = &GetGreetingAll{
			Id: greeting.Id,
		}
	}
	return
}

type GetGreetingDetail struct {
	Id uint `json:"id"`
}

func GetGreetingDetailTransform(e *endpoint.GetGreetingDetailResponse) (res *GetGreetingDetail) {
	if e.Greeting != nil {
		res = &GetGreetingDetail{
			Id: e.Greeting.Id,
		}
	}

	return
}

type CreateGreeting struct {
}

func CreateGreetingTransform(e *endpoint.CreateGreetingResponse) (res *CreateGreeting) {
	res = &CreateGreeting{}
	return
}

type UpdateGreeting struct {
}

func UpdateGreetingTransform(e *endpoint.UpdateGreetingResponse) (res *UpdateGreeting) {
	res = &UpdateGreeting{}
	return
}

type DeleteGreeting struct {
}

func DeleteGreetingTransform(e *endpoint.DeleteGreetingResponse) (res *DeleteGreeting) {
	res = &DeleteGreeting{}
	return
}
