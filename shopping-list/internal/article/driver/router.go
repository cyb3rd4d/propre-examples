package driver

import (
	"context"
	"net/http"
	"shopping-list/internal/article/adapter/decoder"
	"shopping-list/internal/article/adapter/gateway"
	"shopping-list/internal/article/adapter/presenter"
	"shopping-list/internal/article/adapter/presenter/view"
	usecase "shopping-list/internal/article/business/use_case"
	"shopping-list/internal/article/driver/http/middleware"

	"github.com/cyb3rd4d/propre"
	"github.com/spf13/viper"
)

func NewRouter(ctx context.Context) *http.ServeMux {
	router := http.NewServeMux()

	db := NewMysqlConnection(ctx, MysqlOpts{
		User:     viper.GetString("db_user"),
		Password: viper.GetString("db_password"),
		Host:     viper.GetString("db_host"),
		Port:     viper.GetInt("db_port"),
		DBName:   viper.GetString("db_name"),
	})

	repository := gateway.NewMysqlArticleRepository(db)
	httpResponse := propre.NewHTTPResponse[view.Payload]()

	addArticle := propre.NewHTTPHandler(
		decoder.NewAddArticleRequestDecoder(
			propre.NewRequestPayloadExtractor[decoder.AddArticleRequest](propre.JSONDecoder),
		),
		usecase.NewAddArticleInteractor(repository),
		presenter.NewAddArticlePresenter(httpResponse),
	)

	getArticle := propre.NewHTTPHandler(
		decoder.NewGetArticleRequestDecoder(),
		usecase.NewGetArticleInteractor(repository),
		presenter.NewGetArticlePresenter(httpResponse),
	)

	listAllArticles := propre.NewHTTPHandler(
		decoder.NewListAllArticlesRequestDecoder[any](),
		usecase.NewListAllArticlesInteractor[any](repository),
		presenter.NewListAllArticlesPresenter(httpResponse),
	)

	updateArticle := propre.NewHTTPHandler(
		decoder.NewUpdateArticleRequestDecoder(
			propre.NewRequestPayloadExtractor[decoder.UpdateArticleRequest](propre.JSONDecoder),
		),
		usecase.NewUpdateArticleInteractor(repository),
		presenter.NewUpdateArticlePresenter(httpResponse),
	)

	deleteArticle := propre.NewHTTPHandler(
		decoder.NewDeleteArticleRequestDecoder(),
		usecase.NewDeleteArticleInteractor(repository),
		presenter.NewDeleteArticlePresenter(httpResponse),
	)

	router.Handle("POST /article", applyMiddlewares(ctx, addArticle))
	router.Handle("GET /article/{id}", applyMiddlewares(ctx, getArticle))
	router.Handle("GET /article", applyMiddlewares(ctx, listAllArticles))
	router.Handle("PUT /article/{id}", applyMiddlewares(ctx, updateArticle))
	router.Handle("DELETE /article/{id}", applyMiddlewares(ctx, deleteArticle))

	return router
}

func applyMiddlewares(ctx context.Context, next http.Handler) http.Handler {
	next = middleware.AcceptRequestHeaderMiddleware()(next)
	next = middleware.RequestLogMiddleware()(next)
	next = middleware.RequestContextMiddleware(ctx)(next)

	return next
}
