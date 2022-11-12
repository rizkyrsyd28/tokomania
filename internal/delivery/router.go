package delivery

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"tokomania/internal/delivery/product"
	productRepo "tokomania/internal/repository/product"
	productUC "tokomania/internal/usecase/product"
)

func Routes(r *gin.Engine){
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:ras28@localhost:5432/ecommerce")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	repo := productRepo.NewRepo(conn)
	usecase := productUC.NewUseCase(repo)

	r.GET("/products", product.GetProducts(usecase))
	r.GET("/product/:id", product.GetProducts(usecase))
}