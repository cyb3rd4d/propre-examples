package driver

import (
	"context"
	"net/http"

	"shopping-list/internal/article/adapter/gateway"
	"shopping-list/internal/article/adapter/presenter/view"
	"shopping-list/internal/article/driver/http/handler"
	"shopping-list/internal/article/driver/http/middleware"

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
	httpResponse := propre.NewHTTPResponse[view.Payload]()

	addArticleHTTPHandler := handler.NewAddArticleHandler(repository, httpResponse)
	getArticleHTTPHandler := handler.NewGetArticleHandler(repository, httpResponse)
	listAllArticlesHTTPHandler := handler.NewListAllArticlesHandler(repository, httpResponse)
	updateArticleHTTPHandler := handler.NewUpdateArticleHandler(repository, httpResponse)

	srv.Handle("POST /article", applyMiddlewares(ctx, addArticleHTTPHandler))
	srv.Handle("GET /article/{id}", applyMiddlewares(ctx, getArticleHTTPHandler))
	srv.Handle("GET /article", applyMiddlewares(ctx, listAllArticlesHTTPHandler))
	srv.Handle("PUT /article/{id}", applyMiddlewares(ctx, updateArticleHTTPHandler))

	return srv
}

func applyMiddlewares(ctx context.Context, next http.Handler) http.Handler {
	next = middleware.AcceptRequestHeaderMiddleware()(next)
	next = middleware.RequestLogMiddleware()(next)
	next = middleware.RequestContextMiddleware(ctx)(next)

	return next
}
