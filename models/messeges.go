package models

import "gorm.io/gorm"

type Message struct {
	ID      int		'gorm:"primaryKey;autoIncrement" json:"id"'
	Author  string	'json:"author"'
	Message string	'json:"message"'
}

type Messeges []Message

func (m *Messeges) TableName() string {
	return "messages"
}

func MigrateMessages(db *gorm.DB) error {
	err := db.AutoMigrate(&Message{})
	if err != nil {
		return err
	}
	return nil

}