package tablecreator

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Creator 結構，用於連接資料庫
type Creator struct {
	DB *gorm.DB
}

// NewCreator 初始化 Creator 並建立資料庫連接
func NewCreator(dsn string) (*Creator, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Creator{DB: db}, nil
}

// AutoMigrate 自動建表
func (c *Creator) AutoMigrate(model interface{}) error {
	if err := c.DB.AutoMigrate(model); err != nil {
		log.Printf("failed to migrate table: %v", err)
		return err
	}
	log.Println("Table created successfully!")
	return nil
}

// 刪除表：
func (c *Creator) DropTable(model interface{}) error {
	if err := c.DB.Migrator().DropTable(model); err != nil {
		log.Printf("failed to drop table: %v", err)
		return err
	}
	log.Println("Table dropped successfully!")
	return nil
}

// 檢查表是否存在：
func (c *Creator) TableExists(tableName string) bool {
	return c.DB.Migrator().HasTable(tableName)
}
