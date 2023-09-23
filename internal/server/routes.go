package server

import (
	categoryHttp "github.com/mrzalr/cashflo/internal/category/delivery/http"
	categoryRepo "github.com/mrzalr/cashflo/internal/category/repository/mysql"
	categoryUcase "github.com/mrzalr/cashflo/internal/category/usecase"
	transactionHttp "github.com/mrzalr/cashflo/internal/transaction/delivery/http"
	transactionRepo "github.com/mrzalr/cashflo/internal/transaction/repository/mysql"
	transactionUcase "github.com/mrzalr/cashflo/internal/transaction/usecase"
	userHttp "github.com/mrzalr/cashflo/internal/user/delivery/http"
	userRepo "github.com/mrzalr/cashflo/internal/user/repository/mysql"
	userUcase "github.com/mrzalr/cashflo/internal/user/usecase"
)

func (s *server) MapRoutes() {
	userRepository := userRepo.New(s.db)
	categoryRepository := categoryRepo.New(s.db)
	transactionRepository := transactionRepo.New(s.db)

	userUsecase := userUcase.New(userRepository)
	categoryUsecase := categoryUcase.New(categoryRepository)
	transactionUsecase := transactionUcase.New(transactionRepository)

	userHandler := userHttp.New(userUsecase)
	categoryHandler := categoryHttp.New(categoryUsecase)
	transactionHandler := transactionHttp.New(transactionUsecase)

	api := s.app.Group("/api")
	v1 := api.Group("/v1")
	usersRouter := v1.Group("/users")
	categoriesRouter := v1.Group("/categories")
	transactionsRouter := v1.Group("/transactions")

	userHandler.MapRoute(usersRouter)
	categoryHandler.MapRoute(categoriesRouter)
	transactionHandler.MapRoute(transactionsRouter)

}
