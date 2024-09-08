package models

import "github.com/oalexander6/web-app-template/config"

type Store interface {
	noteStore
	Close()
}

type Models struct {
	config *config.Config
	store  Store
}

func New(store Store, config *config.Config) *Models {
	return &Models{
		config: config,
		store:  store,
	}
}
