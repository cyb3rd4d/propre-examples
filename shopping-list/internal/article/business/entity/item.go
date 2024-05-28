package entity

var (
	ArticleEventCreated = ArticleEvent{name: "article_created"}
	ArticleEventUpdated = ArticleEvent{name: "name_updated"}
)

type ArticleEvent struct {
	name string
}

type Article struct {
	id     int
	name   string
	events []ArticleEvent
}

type ArticleOption func(*Article)

func ArticleWithID(id int) ArticleOption {
	return func(article *Article) {
		article.id = id
	}
}

func ArticleWithName(name string) ArticleOption {
	return func(article *Article) {
		article.name = name
	}
}

func NewArticle(opts ...ArticleOption) Article {
	article := Article{}

	for _, option := range opts {
		option(&article)
	}

	if article.id == 0 {
		article.events = []ArticleEvent{ArticleEventCreated}
	}

	return article
}

func (article Article) ID() int {
	return article.id
}

func (article Article) Name() string {
	return article.name
}

func (article Article) Events() []ArticleEvent {
	return article.events
}

func (article *Article) SetName(name string) {
	article.events = append(article.events, ArticleEventUpdated)
	article.name = name
}
