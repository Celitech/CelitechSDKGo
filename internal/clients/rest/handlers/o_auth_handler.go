package handlers

import (
	"errors"
	"fmt"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"github.com/Celitech/CelitechSDKGo/internal/configmanager"
)

// OAuthHandler manages OAuth token injection into requests requiring OAuth authentication.
// It retrieves tokens from the token manager based on requested scopes and injects them as Bearer tokens.
// T is the response type, E is the error type.
type OAuthHandler[T any, E any] struct {
	manager     *configmanager.ConfigManager
	nextHandler Handler[T, E]
}

// NewOAuthHandler creates a new OAuth handler with the provided configuration manager.
// Returns a handler that will fetch and inject OAuth tokens for requests with scopes.
func NewOAuthHandler[T any, E any](manager *configmanager.ConfigManager) *OAuthHandler[T, E] {
	return &OAuthHandler[T, E]{
		manager:     manager,
		nextHandler: nil,
	}
}

// Handle processes a regular request by injecting an OAuth token if scopes are specified.
// Retrieves the token from the token manager and adds it as a Bearer token to the Authorization header.
func (h *OAuthHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	if err := h.addToken(request); err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	return h.nextHandler.Handle(request)
}

// addToken retrieves an OAuth token for the requested scopes and adds it to the request.
// Fetches the token from the token manager and sets the Authorization header with the Bearer token.
func (h *OAuthHandler[T, E]) addToken(request httptransport.Request) error {
	if request.Scopes == nil {
		return nil
	}

	tokenManager := h.manager.GetTokenManager()

	token, err := tokenManager.GetToken(request.Scopes, request.Config)
	if err != nil {
		return fmt.Errorf("error fetching token: %w", err)
	}

	if token.AccessToken != "" {
		authHeader := fmt.Sprintf("Bearer %s", token.AccessToken)
		request.SetHeader("Authorization", authHeader)
	}

	return nil
}

// HandleStream processes a streaming request by injecting an OAuth token if scopes are specified.
// Returns the stream from the next handler after the OAuth token is applied.
func (h *OAuthHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	if err := h.addToken(request); err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	return h.nextHandler.HandleStream(request)
}

// SetNext sets the next handler in the chain.
// This method is called during chain construction to link handlers together.
func (h *OAuthHandler[T, E]) SetNext(handler Handler[T, E]) {
	h.nextHandler = handler
}
