package driver

import (
	"context"
	"net/http"

	"shopping-list/internal/article/adapter/decoder"
	"shopping-list/internal/article/adapter/gateway"
	"shopping-list/internal/article/adapter/presenter"
	usecase "shopping-list/internal/article/business/use_case"
	driverHttp "shopping-list/internal/article/driver/http"
	"shopping-list/internal/article/driver/logger"

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

	repository := gateway.NewMysqlArticleRepository(db)

	addArticleHTTPHandler := propre.NewHTTPHandler(
		decoder.NewAddArticleRequestDecoder(),
		usecase.NewAddArticleInteractor(repository),
		presenter.NewAddArticlePresenter(),
	)

	getArticleHTTPHandler := propre.NewHTTPHandler(
		decoder.NewGetArticleRequestDecoder(),
		usecase.NewGetArticleInteractor(repository),
		presenter.NewGetArticlePresenter(),
	)

	listAllArticlesHTTPHandler := propre.NewHTTPHandler(
		decoder.NewListAllArticlesRequestDecoder[any](),
		usecase.NewListAllArticlesInteractor[any](repository),
		presenter.NewListAllArticlesPresenter(),
	)

	updateArticleHTTPHandler := propre.NewHTTPHandler(
		decoder.NewUpdateArticleRequestDecoder(),
		usecase.NewUpdateArticleInteractor(repository),
		presenter.NewUpdateArticlePresenter(),
	)

	srv.Handle("GET /article", applyMiddlewares(ctx, listAllArticlesHTTPHandler))
	srv.Handle("GET /article/{id}", applyMiddlewares(ctx, getArticleHTTPHandler))
	srv.Handle("PUT /article/{id}", applyMiddlewares(ctx, updateArticleHTTPHandler))
	srv.Handle("POST /article", applyMiddlewares(ctx, addArticleHTTPHandler))

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
