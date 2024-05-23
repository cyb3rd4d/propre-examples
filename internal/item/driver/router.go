package driver

import (
	"context"
	"net/http"

	"github.com/cyb3rd4d/poc-propre/internal/item/adapter/decoder"
	"github.com/cyb3rd4d/poc-propre/internal/item/adapter/gateway"
	"github.com/cyb3rd4d/poc-propre/internal/item/adapter/presenter"
	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
	"github.com/cyb3rd4d/propre"
	"github.com/spf13/viper"
)

func NewRouter(ctx context.Context) *http.ServeMux {
	srv := http.NewServeMux()

	db := NewMysqlConnection(ctx, MysqlOpts{
		User:     viper.GetString("db_user"),
		Password: viper.GetString("db_password"),
		Host:     viper.GetString("db_host"),
		Port:     viper.GetInt("db_port"),
		DBName:   viper.GetString("db_name"),
	})

	repository := gateway.NewMysqlItemRepository(db)

	addItemHTTPHandler := propre.NewHTTPHandler(
		decoder.NewAddItemRequestDecoder(),
		usecase.NewAddItemInteractor(repository),
		presenter.NewAddItemPresenter(),
	)

	getItemHTTPHandler := propre.NewHTTPHandler(
		decoder.NewGetItemRequestDecoder(),
		usecase.NewGetItemInteractor(repository),
		presenter.NewGetItemPresenter(),
	)

	listAllItemsHTTPHandler := propre.NewHTTPHandler(
		decoder.NewListAllItemsRequestDecoder[any](),
		usecase.NewListAllItemsInteractor[any](repository),
		presenter.NewListAllItemsPresenter(),
	)

	srv.Handle("GET /item", listAllItemsHTTPHandler)
	srv.Handle("GET /item/{id}", getItemHTTPHandler)
	srv.Handle("POST /item", addItemHTTPHandler)

	return srv
}
