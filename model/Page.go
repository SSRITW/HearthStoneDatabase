package model

type Page struct {
	PageSize int `form:"page_size" json:"page_size" binding:"required"`
	PageNum int `form:"page_num" json:"page_num" binding:"required"`
}
