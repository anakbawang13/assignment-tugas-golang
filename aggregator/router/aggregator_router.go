package router

import (
	"assignment-tugas-golang/aggregator/handler/grpc"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, aggrHandler *grpc.AggregatorHandler) {
	// User routes
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", aggrHandler.CreateUser)
		userRoutes.GET("/:id", aggrHandler.GetUserByID)
		userRoutes.PUT("/:id", aggrHandler.UpdateUser)
		userRoutes.DELETE("/:id", aggrHandler.DeleteUser)
		userRoutes.GET("", aggrHandler.GetAllUsers)
	}

	// Wallet routes
	walletRoutes := r.Group("/wallets")
	{
		walletRoutes.POST("", aggrHandler.CreateWallet)
		walletRoutes.GET("/:id", aggrHandler.GetWalletByID)
		walletRoutes.PUT("/:id", aggrHandler.UpdateWallet)
		walletRoutes.DELETE("/:id", aggrHandler.DeleteWallet)
		walletRoutes.GET("", aggrHandler.GetAllWallets)
	}

	// Transaction routes
	transactionRoutes := r.Group("/transactions")
	{
		transactionRoutes.POST("", aggrHandler.CreateTransaction)
		transactionRoutes.GET("/:id", aggrHandler.GetTransactionByID)
		transactionRoutes.GET("/wallet/:walletid", aggrHandler.GetTransactionByWalletID)
		transactionRoutes.DELETE("/:id", aggrHandler.DeleteTransaction)
	}
}
