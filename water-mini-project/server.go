package main

import (
	"net/http"
	"water-mini-project/app/repository"
	"water-mini-project/app/utils"

	"github.com/labstack/echo/v4"
)

func main() {

	err := utils.InitializeDB()
	if(err != nil) {
		panic("connection failed!")
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
		
	e.GET("/token", repository.GetAuthGenerateTokenRepository)

	e.GET("/customer/:id", repository.GetPelangganByIdRepository, repository.ValidateAuthTokenRepository)
	e.POST("/customer", repository.CreatePelangganRepository, repository.ValidateAuthTokenRepository)
	e.PUT("/customer/:id", repository.UpdatePelangganRepository, repository.ValidateAuthTokenRepository)
	e.DELETE("/customer/:id", repository.DeletePelangganRepository, repository.ValidateAuthTokenRepository)

	e.GET("/bill/:id", repository.GetBillByIdRepository, repository.ValidateAuthTokenRepository)
	e.POST("/bill", repository.CreateBillRepository, repository.ValidateAuthTokenRepository)
	e.PUT("/bill/:id", repository.UpdateBillRepository, repository.ValidateAuthTokenRepository)
	e.DELETE("/bill/:id", repository.DeleteBillRepository, repository.ValidateAuthTokenRepository)

	e.GET("search-customer/:query", repository.SearchCustomerRepository, repository.ValidateAuthTokenRepository);
	e.GET("search-bill/:query", repository.SearchBillRepository, repository.ValidateAuthTokenRepository);
	e.GET("search-bill-per-customer/:id", repository.GetBillPerCustomerRepository, repository.ValidateAuthTokenRepository);
	e.POST("search-by-period", repository.SearchByPeriodRepository, repository.ValidateAuthTokenRepository);

	e.Logger.Fatal(e.Start(":8080"))
}