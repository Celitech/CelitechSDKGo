package api

import (
	core "github.com/Celitech/CelitechSDKGo/v3/core"
	internal "github.com/Celitech/CelitechSDKGo/v3/internal"
)

var ErrorCodes internal.ErrorCodes = internal.ErrorCodes{
	400: func(apiError *core.APIError) error {
		return &BadRequestError{
			APIError: apiError,
		}
	},
	401: func(apiError *core.APIError) error {
		return &UnauthorizedError{
			APIError: apiError,
		}
	},
	409: func(apiError *core.APIError) error {
		return &ConflictError{
			APIError: apiError,
		}
	},
	422: func(apiError *core.APIError) error {
		return &UnprocessableEntityError{
			APIError: apiError,
		}
	},
	503: func(apiError *core.APIError) error {
		return &ServiceUnavailableError{
			APIError: apiError,
		}
	},
}
