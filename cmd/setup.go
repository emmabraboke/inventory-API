package cmd

import (
	"inventory/cmd/middlewares"
	"inventory/cmd/routes"
	mongoDatabase "inventory/internals/database/mongo"
	mongoCustomerRepo "inventory/internals/repository/customerRepo/mongoRepo"
	mongoinvoiceRepo "inventory/internals/repository/invoiceRepo/mongoRepo"
	mongoProductRepo "inventory/internals/repository/productRepo/mongoRepo"
	mongoSaleRepo "inventory/internals/repository/saleRepo/mongoRepo"
	mongoTransactionRepo "inventory/internals/repository/transactionRepo/mongoRepo"
	mongoUserRepo "inventory/internals/repository/userRepo/mongoRepo"
	"inventory/internals/service/cryptoService"
	"inventory/internals/service/customerService"
	"inventory/internals/service/invoiceService"
	"inventory/internals/service/paymentService"
	"inventory/internals/service/productService"
	"inventory/internals/service/saleService"
	"inventory/internals/service/tokenService"
	"inventory/internals/service/transactionService"
	"inventory/internals/service/userService"
	"inventory/internals/service/validationService"
	"inventory/utils"
	"log"
	"net/http"

	_ "inventory/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() {

	config, err := utils.GetConfig(".")
	paystackUrl := "https://api.paystack.co/transaction/initialize"

	if err != nil {
		log.Fatal(err)
	}

	database := mongoDatabase.NewDatabase(config.DatabaseURI, "inventory")

	conn := database.Connect()
	validationSrv := validationService.NewValidationSrv()

	userDatabase := database.CreateCollection(conn, "user")
	productDatabase := database.CreateCollection(conn, "product")
	saleDatabase := database.CreateCollection(conn, "sale")
	transactionDatabase := database.CreateCollection(conn, "transaction")
	customerDatabase := database.CreateCollection(conn, "customer")
	invoiceDatabase := database.CreateCollection(conn, "invoice")

	userRepo := mongoUserRepo.NewUserRepo(userDatabase)
	customerRepo := mongoCustomerRepo.NewCustomerRepo(customerDatabase)
	productRepo := mongoProductRepo.NewproductRepo(productDatabase)
	saleRepo := mongoSaleRepo.NewSaleRepo(saleDatabase)
	transactionRepo := mongoTransactionRepo.NewTransactionRepo(transactionDatabase)
	invoiceRepo := mongoinvoiceRepo.NewInvoiceRepo(invoiceDatabase)

	tokenSrv := tokenService.NewTokenSrv(config.JWTSecret)
	crypoSrv := cryptoService.NewCryptoService()
	paymentSrv := paymentService.NewPaymentSrv(config.PaystackSecret, paystackUrl)

	userSrv := userService.NewUserSrv(userRepo, validationSrv, crypoSrv, tokenSrv)
	customerSrv := customerService.NewCustomerSrv(customerRepo, validationSrv)
	productSrv := productService.NewproductSrv(productRepo, validationSrv)
	saleSrv := saleService.NewsaleSrv(saleRepo, validationSrv)
	invoiceSrv := invoiceService.NewinvoiceSrv(invoiceRepo, validationSrv, saleSrv)
	transactionSrv := transactionService.NewtransactionSrv(transactionRepo, validationSrv, invoiceSrv, paymentSrv, productSrv)

	router := gin.New()
	router.Use(middlewares.CorsMiddleware())
	group := router.Group("/api/v1")

	// docs.SwaggerInfo.Host = "localhost:5000"
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Title = "Inventory API"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	group.Use(gin.Logger())

	routes.UserRoute(group, userSrv, tokenSrv)
	routes.InvoiceRoute(group, invoiceSrv, tokenSrv)
	routes.ProductRoute(group, productSrv, tokenSrv)
	routes.SaleRoute(group, saleSrv, tokenSrv)
	routes.CustomerRoute(group, customerSrv, tokenSrv)
	routes.TransactionRoute(group, transactionSrv, tokenSrv)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})

	router.Run(":5000")
}
