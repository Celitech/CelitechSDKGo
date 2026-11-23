package api

import (
	json "encoding/json"
	core "github.com/Celitech/CelitechSDKGo/v3/core"
)

// Bad Request
type BadRequestError struct {
	*core.APIError
	Body *ErrorResponse
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body *ErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.StatusCode = 400
	b.Body = body
	return nil
}

func (b *BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

func (b *BadRequestError) Unwrap() error {
	return b.APIError
}

func (b *BadRequestError) GetBody() *ErrorResponse {
	if b == nil {
		var zero *ErrorResponse
		return zero
	}
	return b.Body
}

// Unauthorized
type UnauthorizedError struct {
	*core.APIError
	Body *ErrorResponse
}

func (u *UnauthorizedError) UnmarshalJSON(data []byte) error {
	var body *ErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 401
	u.Body = body
	return nil
}

func (u *UnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnauthorizedError) Unwrap() error {
	return u.APIError
}

func (u *UnauthorizedError) GetBody() *ErrorResponse {
	if u == nil {
		var zero *ErrorResponse
		return zero
	}
	return u.Body
}

// Conflict
type ConflictError struct {
	*core.APIError
	Body *ErrorResponse
}

func (c *ConflictError) UnmarshalJSON(data []byte) error {
	var body *ErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.StatusCode = 409
	c.Body = body
	return nil
}

func (c *ConflictError) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

func (c *ConflictError) Unwrap() error {
	return c.APIError
}

func (c *ConflictError) GetBody() *ErrorResponse {
	if c == nil {
		var zero *ErrorResponse
		return zero
	}
	return c.Body
}

// Unprocessable Entity
type UnprocessableEntityError struct {
	*core.APIError
	Body *ErrorResponse
}

func (u *UnprocessableEntityError) UnmarshalJSON(data []byte) error {
	var body *ErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 422
	u.Body = body
	return nil
}

func (u *UnprocessableEntityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnprocessableEntityError) Unwrap() error {
	return u.APIError
}

func (u *UnprocessableEntityError) GetBody() *ErrorResponse {
	if u == nil {
		var zero *ErrorResponse
		return zero
	}
	return u.Body
}

// Service Unavailable
type ServiceUnavailableError struct {
	*core.APIError
	Body *ErrorResponse
}

func (s *ServiceUnavailableError) UnmarshalJSON(data []byte) error {
	var body *ErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	s.StatusCode = 503
	s.Body = body
	return nil
}

func (s *ServiceUnavailableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Body)
}

func (s *ServiceUnavailableError) Unwrap() error {
	return s.APIError
}

func (s *ServiceUnavailableError) GetBody() *ErrorResponse {
	if s == nil {
		var zero *ErrorResponse
		return zero
	}
	return s.Body
}
