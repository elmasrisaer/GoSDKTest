package handlers

import (
	"errors"
	"fmt"
	"github.com/elmasrisaer/GoSDKTest/internal/clients/rest/httptransport"
	"github.com/elmasrisaer/GoSDKTest/internal/configmanager"
)

type OAuthHandler[T any] struct {
	manager     *configmanager.ConfigManager
	nextHandler Handler[T]
}

func NewOAuthHandler[T any](manager *configmanager.ConfigManager) *OAuthHandler[T] {
	return &OAuthHandler[T]{
		manager:     manager,
		nextHandler: nil,
	}
}

func (h *OAuthHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	if err := h.addToken(request); err != nil {
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	return h.nextHandler.Handle(request)
}

func (h *OAuthHandler[T]) addToken(request httptransport.Request) error {
	if len(request.Scopes) == 0 {
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

func (h *OAuthHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
