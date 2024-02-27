package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/nvlhnn/gommerce/internal/api/middlewares"
	"github.com/nvlhnn/gommerce/internal/handlers"
	"github.com/nvlhnn/gommerce/internal/repositories"
	"github.com/nvlhnn/gommerce/internal/services"
	"gorm.io/gorm"
)

func SetupRoutes(engine *gin.Engine, db *gorm.DB, cache *redis.Client) {

	jwtService := services.NewJWTService()

	cutomerRepo := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(cutomerRepo)
	customerHandler := handlers.NewCustomerHandler(customerService, jwtService)

	productRepo := repositories.NewProductRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	productService := services.NewProductService(productRepo, categoryRepo)
	productHandler := handlers.NewProductHandler(productService)
	
	cartRepo := repositories.NewCartRepository(db)
	cartService := services.NewCartService(cartRepo, productRepo, cache)
	cartHandler := handlers.NewCartHandler(cartService)

	txRepo := repositories.NewTransactionRepository(db)
	txService := services.NewTransactionService(db, cache, txRepo, cartRepo)
	txHandler := handlers.NewTransactionHandler(txService)


	// rate limiting
	engine.Use(middlewares.RateLimitMiddleware(time.Minute, 20)) // Allow 100 requests per minute


	// group api url to api/v11
	api := engine.Group("/api/v1")

	api.POST("/customers/register", customerHandler.Register)
	api.POST("/customers/login", customerHandler.Login)

	api.GET("/products/:category_id", productHandler.List)


	// protect routes
	api.Use(middlewares.AuthorizeJWT(jwtService, customerService))

	api.POST("/carts", cartHandler.AddCart)
	api.GET("/carts", cartHandler.ListCarts)
	api.DELETE("/carts/:product_id", cartHandler.DeleteCart)


	api.POST("/orders", txHandler.CreateOrder)
	api.GET("/orders", txHandler.GetOrders)
}
