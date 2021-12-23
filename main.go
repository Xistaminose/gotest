package main
import (
"log"
// "os"
"github.com/gin-gonic/gin"
routes "github.com/Xistaminose/gotest/routes"
models "github.com/Xistaminose/gotest/config"
)

func main() {
// Connecting to db
db := models.Connect()
router := gin.Default()
// Route Handlers / Endpoints
routes.Routes(router)
log.Fatal(router.Run(":4747"))
}