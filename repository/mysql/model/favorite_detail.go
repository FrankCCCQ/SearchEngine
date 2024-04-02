package model

type FavoriteDetail struct {
	FavoriteDetailID int64 `gorm:"primarykey"`
	UserID           int64
	UrlID            int64
	Url              string
	Desc             string
	Favorite         []*Favorite `gorm:"many2many:f_to_fd;"`
}
