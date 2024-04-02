package model

type Favorite struct {
	FavoriteID     int64             `gorm:"primarykey"`
	UserID         int64             `gorm:"index"`
	FavoriteName   string            `gorm:"unique"`
	FavoriteDetail []*FavoriteDetail `gorm:"many2many:f_to_fd;"`
}
