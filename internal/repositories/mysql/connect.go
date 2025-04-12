package mysql

import (
	"database/sql"
	"fmt"

	//"github.com/andybeak/hexagonal-demo/orm"
	"log"
	"os"
	"strconv"

	"github.com/waskull/gopharm/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Models = []interface{}{
	&domain.User{},
}

type DatabaseParameters struct {
	User string
	Pass string
	Host string
	Name string
	Port int
}

func ProvideDatabaseParameters() *DatabaseParameters {
	port, _ := strconv.ParseInt(os.Getenv("MYSQL_PORT"), 10, 64)
	return &DatabaseParameters{
		User: os.Getenv("MYSQL_USER"),
		Pass: os.Getenv("MYSQL_PASS"),
		Host: os.Getenv("MYSQL_HOST"),
		Port: int(port),
		Name: os.Getenv("MYSQL_DATABASE"),
	}
}

func ProvideDatabase(dbParams *DatabaseParameters) *gorm.DB {
	log.Println("Conectando a la base de datos...")
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci&loc=Local",
			dbParams.User, dbParams.Pass, dbParams.Host, dbParams.Port, dbParams.Name)

	var err error
	sqlDB, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic("Could not connect to database")
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	migrate(*gormDB)
	return gormDB
}

func migrate(g gorm.DB) {
	log.Println("Setting up automigration...")
	db := g.AutoMigrate(Models...)
	if db != nil && db.Error != nil {
		log.Fatal(fmt.Sprintf("Fallo la migración con el siguiente error %s", db.Error()))
	}
	log.Println(fmt.Sprintf("La base de datos se ha migrado exitosamente"))
}
