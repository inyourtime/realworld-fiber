package port

type Repository interface {
	User() UserRepository
	Article() ArticleRepository
}
