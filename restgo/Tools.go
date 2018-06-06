package restgo

func GetPageOffset(pageSize int ,pageNum int)(offset int) {
	offset = pageSize * (pageNum-1)
	return
}
