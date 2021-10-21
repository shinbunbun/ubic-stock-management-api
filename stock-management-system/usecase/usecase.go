package usecase

type Interactor struct {
	UserRepository        UserRepository
	TransactionRepository TransactionRepository
	StockRepository       StockRepository
}
