package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nandarahmat/my-fiber-api/middlewares"
	"github.com/nandarahmat/my-fiber-api/services"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// Routes untuk categories
	apiCategory := api.Group("/category", middlewares.AuthMiddleware)
	apiCategory.Get("/", services.GetCategories)
	apiCategory.Get("/:id", services.GetCategory)
	apiCategoryAdmin := api.Group("/category", middlewares.AdminMiddleware)
	apiCategoryAdmin.Post("/", services.CreateCategory)
	apiCategoryAdmin.Put("/:id", services.UpdateCategory)
	apiCategoryAdmin.Delete("/:id", services.DeleteCategory)

	// Routes untuk provinces
	apiProvince := api.Group("/provcity", middlewares.AuthMiddleware)
	apiProvince.Get("/listprovincies", services.GetProvinces)
	apiProvince.Get("/detailprovince/:id", services.GetDetailProvinces)
	apiProvince.Get("/listcities/:id", services.GetCity)
	apiProvince.Get("/detailcity/:id", services.GetDetailCity)

	// Routes untuk user
	apiUser := api.Group("/user", middlewares.AuthMiddleware)
	apiUser.Get("/", services.GetUser)
	apiUser.Put("/", services.UpdateUser)
	apiUser.Get("/alamat", services.GetUserAlamat)
	apiUser.Get("/alamat/:id", services.GetUserAlamatById)
	apiUser.Post("/alamat", services.CreateUserAlamat)
	apiUser.Put("/alamat/:id", services.UpdateUserAlamat)
	apiUser.Delete("/alamat/:id", services.DeleteUserAlamat)

	// Router untuk toko
	apiToko := api.Group("/toko", middlewares.AuthMiddleware)
	apiToko.Get("/my", services.GetMyToko)
	apiToko.Put("/my", services.UpdateMyToko)
	apiToko.Get("/", services.GetAllToko)
	apiToko.Get("/:id", services.GetTokoByID)

	// Routes Products
	apiProduk := api.Group("/product", middlewares.AuthMiddleware)
	apiProduk.Get("/", services.GetProducts)
	apiProduk.Get("/:id", services.GetProductById)
	apiProduk.Post("/", services.CreateProduct)
	apiProduk.Put("/:id", services.UpdateProduct)
	apiProduk.Delete("/:id", services.DeleteProduct)

	// Routes Transactions
	apiTrx := api.Group("/trx", middlewares.AuthMiddleware)
	apiTrx.Get("/", services.GetAllTrx)
	apiTrx.Get("/:id", services.GetTrxByID)
	apiTrx.Post("/", services.StoreTrx)

	// Auth routes
	api.Post("/auth/register", services.Register)
	api.Post("/auth/login", services.Login)
}
