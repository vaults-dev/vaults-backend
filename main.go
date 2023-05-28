package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vaults-dev/vaults-backend/initializers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vaults-dev/vaults-backend/graph"
)

func init() {
	// Load .env file
	godotenv.Load()

	initializers.ConnectDB()
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()

	// r.Use(cors.New(
	// 	cors.Config{
	// 		AllowAllOrigins:  true,
	// 		AllowCredentials: true,
	// 		AllowMethods:     []string{"POST", "GET", "PUT", "OPTIONS"},
	// 	}))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	// config.AllowHeaders = []string{}
	r.Use(cors.New(config))

	// r.Use(cors.Default())

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()

	// r.POST("/sign-up", controllers.SignUp)
	// r.POST("/login", controllers.Login)
	// r.GET("/home", middlewares.ValidateAuth, controllers.Home)
	r.Run() // listen and serve on 0.0.0.0:8080
}
