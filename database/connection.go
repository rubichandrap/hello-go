package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseEnvironmentSetup struct {
	DBHost     string
	DBPort     uint64
	DBUser     string
	DBPassword string
	DBName     string
}

func Connect() (*gorm.DB, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbSetup := DatabaseEnvironmentSetup{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
	}

	var (
		db  *gorm.DB
		err error
	)

	if dbDriver == "postgres" {
		db, err = postgreSQLConnection(&dbSetup)
	} else {
		db, err = mySQLConnection(&dbSetup)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		defer sqlDB.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}

func postgreSQLConnection(d *DatabaseEnvironmentSetup) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", d.DBHost, d.DBPort, d.DBUser, d.DBPassword, d.DBName)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func mySQLConnection(d *DatabaseEnvironmentSetup) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.DBUser, d.DBPassword, d.DBHost, d.DBPort, d.DBName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
