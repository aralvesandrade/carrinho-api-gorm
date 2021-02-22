package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"carrinho-api-gorm/api/handler"
	"carrinho-api-gorm/business"
	"carrinho-api-gorm/config"
	"carrinho-api-gorm/controller"
	"carrinho-api-gorm/repository"
)

//Server ...
type Server struct{}

//StartServer ...
func (s *Server) StartServer() {

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := 3306
	port, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
	dbName := os.Getenv("MYSQL_NAME")

	mysql, error := config.InitDb(user, password, host, dbName, port)

	if error != nil {
		fmt.Println("Error on connecting to MySQL Database", error)
		os.Exit(1)
	}

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"X-Requested-With", "Content-Type", "Authorization", "accesstoken"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedOrigins: []string{"*"},
	})

	//Setting health endpoint
	handler.NewHealthCheckHandler(r)

	itensCarrinhoRepository := repository.NewItensCarrinhoRepository(mysql)
	itensCarrinhoBusiness := business.NewItensCarrinhoBusiness(itensCarrinhoRepository)

	carrinhoRepository := repository.NewCarrinhoRepository(mysql)
	carrinhoBusiness := business.NewCarrinhoBusiness(carrinhoRepository)
	carrinhoController := controller.NewCarrinhoController(carrinhoBusiness, itensCarrinhoBusiness)

	//Setting carrinho endpoint
	handler.NewCarrinhoHandler(r, carrinhoController)

	router := mux.NewRouter()
	router.Handle("/{_:.*}", r)

	port, _ = strconv.Atoi(os.Getenv("PORT"))

	if port == 0 {
		port = 5001
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
		Handler:      handlers.CompressHandler(handlers.RecoveryHandler()(c.Handler(router))),
	}

	if error := httpServer.ListenAndServe(); error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
}
