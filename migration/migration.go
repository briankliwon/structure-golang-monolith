package migration

import (
    "github.com/saiful344/structured_go_project/model"
    "github.com/jinzhu/gorm"
)

func Migration(db *gorm.DB){
    db.AutoMigrate(&model.Task{})
    db.AutoMigrate(&model.User{})
}
