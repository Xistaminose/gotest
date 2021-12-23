package config
import (
"log"
"os"
"github.com/go-pg/pg/v10"
"github.com/go-pg/pg/v10/orm"
)
// Connecting to db
func Connect() *pg.DB {
opts := &pg.Options{
User: "postgres",
Password: "postgres",
Addr: "localhost:5432",
Database: "segurospromo",
}
var db *pg.DB = pg.Connect(opts)
if db == nil {
log.Printf("Failed to connect")
os.Exit(100)
}
log.Printf("Connected to db")
return db
}