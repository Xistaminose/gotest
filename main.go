package main
import (
"log"
"os"
"github.com/gin-gonic/gin"
routes "github.com/cavdy-play/go_db/routes"
)
func main() {
// Init Router
router := gin.Default()
// Route Handlers / Endpoints
routes.Routes(router)
log.Fatal(router.Run(":4747"))
}