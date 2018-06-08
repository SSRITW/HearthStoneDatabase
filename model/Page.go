package model

type Page struct {
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"`
	PageNum int `form:"pageNum" json:"pageNum" binding:"required"`
}
