package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql
	"gopkg.in/gormigrate.v1"
)

// Session DB
var Session *gorm.DB

// Init DB
func init() {
	var err error
	Session, err = gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME")),
	)
	Session = Session.Omit("created_at", "updated_at")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Printf("DB Connected!")

	// Migrations
	m := gormigrate.New(Session, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create users table
		{
			ID: "201608301400",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					ID        int       `gorm:"PRIMARY_KEY"`
					Name      string    `gorm:"fullname" binding:"required" valid:"required"`
					Username  string    `gorm:"username" binding:"required" valid:"required"`
					Password  string    `gorm:"password,omitempty" binding:"required" valid:"stringlength(5|9),required"`
					Email     string    `gorm:"email" binding:"required" valid:"email,required"`
					DOB       string    `gorm:"dob"`
					Gender    string    `gorm:"gender"`
					Avatar    string    `gorm:"avatar"`
					Payload   string    `gorm:"payload"`
					Active    bool      `gorm:"active"`
					CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
					UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
				}
				return tx.AutoMigrate(&User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
