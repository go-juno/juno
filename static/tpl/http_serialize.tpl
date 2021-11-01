package serialize

import "juno/internal/endpoint"

type GetGreetingList struct {
	Id uint `json:"id"`
}

func GetGreetingListTransform(e *endpoint.GetGreetingListResponse) (res *List) {
	items := make([]*GetGreetingList, len(e.Items))
	res = &List{
		Total: e.Total,
	}
	for index, s := range e.Items {
		items[index] = &GetGreetingList{
			Id: s.Id,
		}
	}
	res.Items = items
	return
}

type GetGreetingAll struct {
	Id uint `json:"id"`
}

func GetGreetingAllTransform(e []*endpoint.GetGreetingAllResponse) (res []*GetGreetingAll) {
	res = make([]*GetGreetingAll, len(e))
	for index, s := range e {
		res[index] = &GetGreetingAll{
			Id: s.Id,
		}
	}
	return
}

type GetGreetingDetail struct {
	Id uint `json:"id"`
}

func GetGreetingDetailTransform(s *endpoint.GetGreetingDetailResponse) (res *GetGreetingDetail) {
	if s != nil {
		res = &GetGreetingDetail{
			Id: s.Id,
		}
	}
	return
}

type CreateGreeting struct {
}

func CreateGreetingTransform(s *endpoint.CreateGreetingResponse) (res *CreateGreeting) {
	if s != nil {
		res = &CreateGreeting{}
	}
	return
}

type UpdateGreeting struct {
}

func UpdateGreetingTransform(s *endpoint.UpdateGreetingResponse) (res *UpdateGreeting) {
	if s != nil {
		res = &UpdateGreeting{}
	}
	return
}

type DeleteGreeting struct {
}

func DeleteGreetingTransform(s *endpoint.DeleteGreetingResponse) (res *DeleteGreeting) {
	if s != nil {
		res = &DeleteGreeting{}
	}
	return
}
