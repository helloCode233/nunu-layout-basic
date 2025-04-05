package handler

import (
	"github.com/go-nunu/nunu-layout-basic/pkg/log"
)

type Handler struct {
	logger *log.Logger
}

// @wire:Handler
func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
