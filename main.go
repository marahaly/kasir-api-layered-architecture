package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"	
	"os"
	"log"
	"marahaly-kasir-api/database"
	"marahaly-kasir-api/repositories"
	"marahaly-kasir-api/services"
	"marahaly-kasir-api/handlers"
	"github.com/spf13/viper"
)

type Config struct {
	Port	string `mapstructure:"PORT"`
	DBConn	string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}
	
	config := Config{
		Port: viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}
	
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()
	
	
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)
	
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)
	
	
	/*** API Health Check ***/
	/*** HEALTH - /health ***/
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "OK",
			"message": "API is Running",
		})
	})
	
	fmt.Println("Server is running at localhost:" + config.Port)
	
	errServ := http.ListenAndServe(":" + config.Port,nil)
	if errServ != nil {
		fmt.Println("Server is failed running at localhost:" + config.Port)
	}
}
