package restgo

func GetPageOffset(pageSize int ,pageNum int)(offset int) {
	if pageSize!=0 || pageNum!=0 {
		offset = pageSize * (pageNum-1)
	}
	return
}
