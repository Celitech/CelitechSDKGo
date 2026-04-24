package celitech

import (
	"github.com/Celitech/CelitechSDKGo/celitechconfig"
	"github.com/Celitech/CelitechSDKGo/internal/oauthtokenmanager"
	"github.com/Celitech/CelitechSDKGo/param"
)

// The type aliases below let consumers use a single import path for the entire SDK.
// Internally the concrete types live in celitechconfig and param.

// Config holds all configuration parameters for the SDK client.
type Config = celitechconfig.Config

// RequestOption is a function that configures a single request.
type RequestOption = celitechconfig.RequestOption

// Environment defines the available API base URLs.
type Environment = celitechconfig.Environment

// RetryConfig holds all runtime-configurable retry parameters.
type RetryConfig = celitechconfig.RetryConfig

// NewConfig creates a Config with spec-derived defaults.
var NewConfig = celitechconfig.NewConfig

// NewRetryConfig returns a RetryConfig initialized with spec-derived defaults.
var NewRetryConfig = celitechconfig.NewRetryConfig

// WithBaseURL returns a RequestOption that overrides BaseURL for a single request.
var WithBaseURL = celitechconfig.WithBaseURL

// WithTimeout returns a RequestOption that overrides Timeout for a single request.
var WithTimeout = celitechconfig.WithTimeout

// WithClientID returns a RequestOption that overrides ClientID for a single request.
var WithClientID = celitechconfig.WithClientID

// WithClientSecret returns a RequestOption that overrides ClientSecret for a single request.
var WithClientSecret = celitechconfig.WithClientSecret

// WithOAuthBaseURL returns a RequestOption that overrides OAuthBaseURL for a single request.
var WithOAuthBaseURL = celitechconfig.WithOAuthBaseURL

// WithRetryConfig returns a RequestOption that overrides the RetryConfig for a single request.
var WithRetryConfig = celitechconfig.WithRetryConfig

// Nullable returns a *param.Nullable[T] set to v — use for nullable fields with a value.
func Nullable[T any](v T) *param.Nullable[T] { return &param.Nullable[T]{Value: v} }

// Null returns a *param.Nullable[T] with IsNull set to true, signalling an explicit JSON null.
func Null[T any]() *param.Nullable[T] { return param.Null[T]() }

// Ptr returns a pointer to v — use when no type-specific helper exists.
func Ptr[T any](v T) *T { return param.Ptr(v) }

// Environment constants for the available API base URLs.
const (
	DefaultEnvironment Environment = celitechconfig.DefaultEnvironment
)

// TokenCache is the interface for persisting OAuth access tokens across process restarts.
// Implement this interface and pass it to SetTokenCache to enable token persistence.
type TokenCache = oauthtokenmanager.TokenCache

// OAuthToken represents a cached OAuth access token with scopes and optional expiry.
type OAuthToken = oauthtokenmanager.OAuthToken

// NewOAuthToken creates a new OAuthToken — use when implementing a custom TokenCache.
var NewOAuthToken = oauthtokenmanager.NewOAuthToken
