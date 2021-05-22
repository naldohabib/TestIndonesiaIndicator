package models

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Data Scrape ...
type YoutubeData struct {
	gorm.Model
	ChannelId   string `gorm:"not null;size:255" json:"channelId"`
	Title       string `gorm:"not null;size:255" json:"title"`
	ChannelName string `gorm:"not null;size:500" json:"channelName"`
	PublishedAt string `gorm:"not null;size:500" json:"publishedAt"`
}

// TableName ..
func (u YoutubeData) TableName() string {
	return "tb_data"
}

// Validate ...
func (u *YoutubeData) Validate() error {

	if err := validation.Validate(u.Title, validation.Required); err != nil {
		return errors.New("title cannot be blank")
	}

	if err := validation.Validate(u.ChannelName, validation.Required); err != nil {
		return errors.New("Channel Name cannot be blank")
	}

	if err := validation.Validate(u.PublishedAt, validation.Required); err != nil {
		return errors.New("PublishedAt cannot be blank")
	}
	return nil

}