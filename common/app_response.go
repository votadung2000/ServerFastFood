package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Filter interface{} `json:"filter,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
}

func SuccessResponse(data, filter, paging interface{}) *successResponse {
	return &successResponse{Data: data, Filter: filter, Paging: paging}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return SuccessResponse(data, nil, nil)
}
