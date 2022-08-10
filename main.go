package main

import (
	"SkateShop/models"
	"SkateShop/routes"
	"SkateShop/utils"
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
        "github.com/go-playground/validator/v10"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    router := gin.New()
    router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        // your custom format
        return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
        param.ClientIP,
        param.TimeStamp.Format(time.RFC1123),
        param.Method,
        param.Path,
        param.Request.Proto,
        param.StatusCode,
        param.Latency,
        param.Request.UserAgent(),
        param.ErrorMessage,
        )
    }))

    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
        if err, ok := recovered.(string); ok {
            c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
        }
        c.AbortWithStatus(http.StatusInternalServerError)
    }))

    // Handle db connection
    models.InitDB()

    // Handle routes
    v1 := router.Group("/api/v1")
    routes.UserRouterGroup(v1.Group("/users"))

    // Hanlde validators
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("roles", utils.RoleValidator)
    }

    // Will run on 8080 by default, but if a PORT environment variable
    // is defined, it will override the default.
    var port string = "8080"
    if os.Getenv("PORT") != "" {
        port = os.Getenv("PORT")
    }
    router.Run(":" + port)
}
