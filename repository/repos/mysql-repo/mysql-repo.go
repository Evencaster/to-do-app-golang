package mysql_repo

import (
	"fmt"
	repo "github.com/Evencaster/to-do-app-golang/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"os"
)

type TaskModel struct {
	gorm.Model
	ID 		int
	Name 	string
}

type MySQLRepo struct {
	db *gorm.DB
}

func connectDB() *gorm.DB {
	DB_USERNAME := "root"
	DB_PASSWORD := "root"
	DB_NAME := "to-do-app-db"
	DB_HOST := "db"
	DB_PORT := "3306"

	connectionParams := url.Values{}
	connectionParams.Set("parseTime","true")
	connectionParams.Set("loc","Local")

	//DB_USERNAME := os.Getenv("DB_USERNAME")
	//DB_PASSWORD := os.Getenv("DB_PASSWORD")
	//DB_NAME := os.Getenv("DB_NAME")
	//DB_HOST := os.Getenv("DB_HOST")
	//DB_PORT := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)%s?%s",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
		connectionParams.Encode(),
		)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewMySQLRepo() *MySQLRepo {
	db := connectDB()
	return &MySQLRepo{db: db}
}

func (r *MySQLRepo) GetAllTasks() []repo.Task {

}

func (r *MySQLRepo) AddTask(name string)  {

}

func (r *MySQLRepo) RemoveTask(id uint64)  {

}

func (r *MySQLRepo) RemoveAllTasks()  {

}

