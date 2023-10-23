package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "os"
    "fmt"
)

// BasicAuthCmsMiddleware adalah middleware untuk otentikasi Basic Auth
func BasicAuthCmsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Ambil nilai username dan password dari variabel lingkungan
        username := os.Getenv("BASIC_AUTH_CMS_USERNAME")
        password := os.Getenv("BASIC_AUTH_CMS_PASSWORD")

        // Ekstrak kredensial Basic Auth dari header permintaan
        user, pass, ok := c.Request.BasicAuth()


        // Periksa apakah kredensial sesuai
        if !ok || user != username || pass != password {
            c.JSON(http.StatusUnauthorized, gin.H{"data": nil, "error_message": "Unauthorized", "status": 401})
            c.Abort()
            return
        }

        // Lanjutkan eksekusi berikutnya jika kredensial valid
        c.Next()
    }
}
