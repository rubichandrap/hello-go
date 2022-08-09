package dtos

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponsePayload struct {
	Message string
	Data    interface{}
}

func (r *Response) Send(statusCode int, p ResponsePayload) *Response {
	switch statusCode {
	case 200:
		return r.OK(p)
	case 201:
		return r.Created(p)
	case 401:
		return r.Unauthorized(p)
	case 404:
		return r.NotFound(p)
	case 500:
		return r.InternalServerError(p)
	default:
		return r.BadRequest(p)
	}
}

func (r *Response) OK(p ResponsePayload) *Response {
	return &Response{
		Success: true,
		Message: "Success",
		Data:    p.Data,
	}
}

func (r *Response) Created(p ResponsePayload) *Response {
	return &Response{
		Success: true,
		Message: "Created",
		Data:    p.Data,
	}
}

func (r *Response) BadRequest(p ResponsePayload) *Response {
	return &Response{
		Success: false,
		Message: p.Message,
		Data:    nil,
	}
}

func (r *Response) Unauthorized(p ResponsePayload) *Response {
	return &Response{
		Success: false,
		Message: p.Message,
		Data:    nil,
	}
}

func (r *Response) NotFound(p ResponsePayload) *Response {
	return &Response{
		Success: false,
		Message: p.Message,
		Data:    nil,
	}
}

func (r *Response) InternalServerError(p ResponsePayload) *Response {
	return &Response{
		Success: false,
		Message: p.Message,
		Data:    nil,
	}
}
