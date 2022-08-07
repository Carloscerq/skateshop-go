package main;

import (
    "github.com/gin-gonic/gin"
    "os"
    "net/http"
    "fmt"
    "time"
)

func main() {
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

    // Will run on 8080 by default, but if a PORT environment variable
    // is defined, it will override the default.
    var port string = "8080"
    if os.Getenv("PORT") != "" {
        port = os.Getenv("PORT")
    }
    router.Run(":" + port)
}
