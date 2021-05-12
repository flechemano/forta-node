package main

import (
	"context"

	"OpenZeppelin/fortify-node/config"
	"OpenZeppelin/fortify-node/services"
	"OpenZeppelin/fortify-node/store"
)

func initApi(ctx context.Context, as store.AlertStore, cfg config.Config) (*services.AlertApi, error) {
	return services.NewAlertApi(ctx, as, services.AlertApiConfig{Port: 80})
}

func initListener(ctx context.Context, as store.AlertStore, cfg config.Config) (*services.AlertListener, error) {
	return services.NewAlertListener(ctx, as, services.AlertListenerConfig{Port: 8770})
}

func initPruner(ctx context.Context, as store.AlertStore, cfg config.Config) (*services.DBPruner, error) {
	return services.NewDBPruner(ctx, as)
}

func initServices(ctx context.Context, cfg config.Config) ([]services.Service, error) {
	as, err := store.NewBadgerAlertStore()
	if err != nil {
		return nil, err
	}

	api, err := initApi(ctx, as, cfg)
	if err != nil {
		return nil, err
	}

	listener, err := initListener(ctx, as, cfg)
	if err != nil {
		return nil, err
	}

	pruner, err := initPruner(ctx, as, cfg)
	if err != nil {
		return nil, err
	}

	return []services.Service{
		api,
		listener,
		pruner,
	}, nil
}

func main() {
	services.ContainerMain("query", initServices)
}
