package main
import (
"log"
"os"
"github.com/gin-gonic/gin"
routes "github.com/Xistaminose/gotest/routes"
)
func main() {
// Init Router
router := gin.Default()
// Route Handlers / Endpoints
routes.Routes(router)
log.Fatal(router.Run(":4747"))
}