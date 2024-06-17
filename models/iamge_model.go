package models

type ImageModel struct {
	MODEL
	Path      string    `json:"path"`
	ImageType ImageType `gorm:"type=smallint(6)" json:"image_type"`
	Key       string    `gorm:"column:image_key" json:"key"`
	Name      string    `json:"name"`
}

type ImageType int
