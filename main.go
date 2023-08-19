package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vaults-dev/vaults-backend/controllers"
	"github.com/vaults-dev/vaults-backend/initializers"
	"github.com/vaults-dev/vaults-backend/libraries"
	"github.com/vaults-dev/vaults-backend/repositories"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vaults-dev/vaults-backend/graph"
)

func init() {
	// Load .env file
	godotenv.Load()

	initializers.GenerateJwk()

	controllers.GoogleOauthConfig = &oauth2.Config{
		RedirectURL: "https://vaults-backend-production.up.railway.app/google/callback",
		// RedirectURL:  "http://localhost:8080/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

// Defining the Graphql handler
func graphqlHandler(walletRepo *repositories.WalletRepository) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		WalletRepository: walletRepo,
	}}))

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
	// mysql
	gormDB := initializers.MysqlConnectDB()
	initializers.MigrateTable(gormDB)

	// repository
	userRepo := repositories.NewUserRepository(gormDB)
	userLib := libraries.NewUserLibrary(userRepo)
	userCtrl := controllers.NewUserController(userLib)

	walletRepo := repositories.NewWalletRepository(gormDB)

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.GET("/", playgroundHandler())
	r.POST("/query", graphqlHandler(walletRepo))

	r.GET("/jwk", controllers.GetJwk)
	r.POST("/sign-up", userCtrl.SignUp)
	r.POST("/login", userCtrl.Login)
	r.GET("/login-page", controllers.LoginPage)
	r.GET("/google-oauth", controllers.GoogleOAuth)
	r.GET("/google/callback", controllers.GoogleCallback)

	// r.GET("/home", middlewares.ValidateAuth, controllers.Home)
	r.Run() // listen and serve on 0.0.0.0:8080
}
