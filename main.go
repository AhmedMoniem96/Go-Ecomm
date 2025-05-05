package main

import (
	"fmt"
	"log"
	"my-go-project/db"
	"my-go-project/handlers"
	"my-go-project/repository"
	"my-go-project/routes"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	//**************************** connection to database **********************
	dbPool, err := db.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect with db", err)
	}
	defer dbPool.Close()

	//**************************** repository setup **********************
	productRepo := repository.NewProductRepository(dbPool)
	userRepo := repository.NewUserRepository(dbPool)

	//**************************** handler setup **********************
	productHandler := handlers.NewProductHandler(productRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	//**************************** route setup **********************
	r := routes.SetupRoutes(productHandler, userHandler)

	//**************************** CORS setup **********************
	c := cors.New(cors.Options{
		// السماح بالطلبات من هذه الـ Origins المختلفة (محليًا وإنتاجيًا)
		AllowedOrigins: []string{
			"http://localhost:5173", // Localhost for development
			"https://react-d-3c1e.vercel.app", // React app deployed on Vercel
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap the router with CORS handler
	handler := c.Handler(r)

	//**************************** run the server **********************
	port := "8080"
	fmt.Printf("✅ Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
