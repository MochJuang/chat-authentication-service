package entity

import (
	"time"
)

type Conversation struct {
	ID           uint      `gorm:"primaryKey"`
	UUID         string    `gorm:"unique;not null"`
	Participants []User    `gorm:"many2many:conversation_participants"`
	Messages     []Message `gorm:"foreignKey:ConversationID"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
