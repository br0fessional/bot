package interfaces

import (
	"gorm.io/gorm"
	"time"
)

type MessageStats struct {
	gorm.Model
	ChatID int64 `gorm:"<-:create;index"`
	UserID int64 `gorm:"<-:create;index"`
	//UserStats Stats     `gorm:"foreignKey:user_id;references:user_id;constraint:OnDelete:CASCADE"`
	Time  time.Time `gorm:"<-:create;index"`
	Words int       `gorm:"<-:create"`
}

type MessageStatsWordCountStruct struct {
	UserID   int64
	Username string
	Words    int
}

type MessageStatsRepoInterface interface {
	InsertMessageStats(chatID int64, userID int64, words int) error
	GetKnownUserIDs(chatID int64) ([]int64, error)
	GetWordCounts(chatID int64) ([]MessageStatsWordCountStruct, error)
}
