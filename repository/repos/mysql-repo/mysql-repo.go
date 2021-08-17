package mysql_repo

import (
	"fmt"
	repo "github.com/Evencaster/to-do-app-golang/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	connectionParams := url.Values{}
	connectionParams.Set("parseTime","true")
	connectionParams.Set("loc","Local")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
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
	err := db.AutoMigrate(TaskModel{})
	if err != nil {
		log.Fatal(err)
	}
	return &MySQLRepo{db: db}
}

func (r *MySQLRepo) GetAllTasks() []repo.Task {
	var tasks []TaskModel
	err := r.db.Find(&tasks).Error
	if err != nil {
		log.Fatal(err)
	}
	var out []repo.Task

	for _, t := range tasks {
		out = append(out, repo.Task{
			Name: t.Name,
			ID: uint64(t.ID),
			Timestamp: t.CreatedAt.Unix(),
		})
	}

	return out
}

func (r *MySQLRepo) AddTask(name string) uint64 {
	task := TaskModel{Name: name}

	err := r.db.Create(&task).Error
	if err != nil {
		log.Fatal(err)
	}
	return uint64(task.ID)
}

func (r *MySQLRepo) RemoveTask(id uint64)  {
	r.db.Where("id = ?", id).Delete(&TaskModel{})
}

func (r *MySQLRepo) RemoveAllTasks()  {
	r.db.Delete(&TaskModel{})
}

