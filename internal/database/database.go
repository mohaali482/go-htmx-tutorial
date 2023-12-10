package database

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Service interface {
	CreateCount() CountType
	GetAllCounts() ([]CountType, error)
	GetCount(id int) (CountType, error)
	IncrementCount(id int) error
	DecrementCount(id int) error
	DeleteCount(id int) error
}

type service struct {
	db *gorm.DB
}

var (
	dburl = os.Getenv("DB_URL")
)

type CountType struct {
	Id    int
	Value int
}

type Count struct {
	gorm.Model
	Value int
}

func (c Count) ToCountType() CountType {
	return CountType{Id: int(c.ID), Value: c.Value}
}

func New() Service {
	db, err := gorm.Open(sqlite.Open(dburl), &gorm.Config{})
	db.AutoMigrate(&Count{})
	if err != nil {
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) CreateCount() CountType {
	count := Count{Value: 0}
	s.db.Create(&count)
	return count.ToCountType()
}

func (s *service) GetAllCounts() ([]CountType, error) {
	var counts []Count
	s.db.Find(&counts)
	var countTypes []CountType
	for _, count := range counts {
		countTypes = append(countTypes, count.ToCountType())
	}
	return countTypes, nil
}

func (s *service) GetCount(id int) (CountType, error) {
	var count Count
	s.db.First(&count, id)
	countType := count.ToCountType()
	return countType, nil
}

func (s *service) IncrementCount(id int) error {
	var count Count
	s.db.First(&count, id)
	count.Value++
	s.db.Save(&count)
	return nil
}

func (s *service) DecrementCount(id int) error {
	var count Count
	s.db.First(&count, id)
	count.Value--
	s.db.Save(&count)
	return nil
}

func (s *service) DeleteCount(id int) error {
	var count Count
	return s.db.Delete(&count, id).Error
}
