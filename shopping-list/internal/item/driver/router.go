package driver

import (
	"context"
	"net/http"

	"shopping-list/internal/item/adapter/decoder"
	"shopping-list/internal/item/adapter/gateway"
	"shopping-list/internal/item/adapter/presenter"
	usecase "shopping-list/internal/item/business/use_case"
	driverHttp "shopping-list/internal/item/driver/http"
	"shopping-list/internal/item/driver/logger"

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

	updateItemHTTPHandler := propre.NewHTTPHandler(
		decoder.NewUpdateItemRequestDecoder(),
		usecase.NewUpdateItemInteractor(repository),
		presenter.NewUpdateItemPresenter(),
	)

	srv.Handle("GET /item", applyMiddlewares(ctx, listAllItemsHTTPHandler))
	srv.Handle("GET /item/{id}", applyMiddlewares(ctx, getItemHTTPHandler))
	srv.Handle("PUT /item/{id}", applyMiddlewares(ctx, updateItemHTTPHandler))
	srv.Handle("POST /item", applyMiddlewares(ctx, addItemHTTPHandler))

	return srv
}

type middleware func(http.Handler) http.Handler

func applyMiddlewares(ctx context.Context, next http.Handler) http.Handler {
	next = driverHttp.AcceptRequestHeaderMiddleware()(next)
	next = requestLogMiddleware()(next)
	next = requestContextMiddleware(ctx)(next)

	return next
}

func requestContextMiddleware(ctx context.Context) middleware {
	return func(next http.Handler) http.Handler {
		fn := func(rw http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(rw, req.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func requestLogMiddleware() middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			logger.FromContext(req.Context()).Debug(
				"[http] incoming request",
				"method", req.Method,
				"uri", req.RequestURI,
			)

			next.ServeHTTP(rw, req)
		})
	}
}
