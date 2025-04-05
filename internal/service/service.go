package service

import "github.com/go-nunu/nunu-layout-basic/pkg/log"

type Service struct {
	logger *log.Logger
}

// @wire:Service
func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}
