package repositories

type RepositoryResult struct {
	Data       interface{}
	Error      error
	StatusCode int
}
