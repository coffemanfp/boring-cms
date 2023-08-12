package psql

// parsePagination calculates the limit and offset for pagination based on the provided page number.
func parsePagination(page int) (limit, offset int) {
	// Calculate the limit for the current page (adding 1 to the page and multiplying by 20).
	limit = (page + 1) * 20
	// Calculate the offset for the current page (multiplying the page by 20).
	offset = page * 20
	return
}
