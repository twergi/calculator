package mocks

//go:generate mockgen -destination ./mock_repo.go -package mocks -mock_names Repository=MockRepository github.com/twergi/calculator/internal/app/usecases/calculator Repository
