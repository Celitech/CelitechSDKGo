package rest

import (
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/handlers"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/hooks"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
)

type RestClient[T any] struct {
	handlers *handlers.HandlerChain[T]
}

func NewRestClient[T any](config celitechconfig.Config, manager *configmanager.ConfigManager) *RestClient[T] {
	retryHandler := handlers.NewRetryHandler[T]()
	oAuthHandler := handlers.NewOAuthHandler[T](manager)
	responseValidationHandler := handlers.NewResponseValidationHandler[T]()
	unmarshalHandler := handlers.NewUnmarshalHandler[T]()
	requestValidationHandler := handlers.NewRequestValidationHandler[T]()
	hookHandler := handlers.NewHookHandler[T](hooks.NewDefaultHook())
	terminatingHandler := handlers.NewTerminatingHandler[T]()

	handlers := handlers.BuildHandlerChain[T]().
		AddHandler(retryHandler).
		AddHandler(oAuthHandler).
		AddHandler(responseValidationHandler).
		AddHandler(unmarshalHandler).
		AddHandler(requestValidationHandler).
		AddHandler(hookHandler).
		AddHandler(terminatingHandler)

	return &RestClient[T]{
		handlers: handlers,
	}
}

func (client *RestClient[T]) Call(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	return client.handlers.CallApi(request)
}
