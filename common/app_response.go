package common

type successResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Filter     interface{} `json:"filter,omitempty"`
	Paging     interface{} `json:"paging,omitempty"`
}

func SuccessResponse(data, filter, paging interface{}) *successResponse {
	return &successResponse{StatusCode: 200, Data: data, Filter: filter, Paging: paging}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return SuccessResponse(data, nil, nil)
}
