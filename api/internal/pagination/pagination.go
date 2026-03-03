package pagination

import (
	"net/http"
	"strconv"
)

type Params struct {
	Page   int
	Limit  int
	Offset int
}

func Parse(r *http.Request, defaultLimit, maxLimit int) Params {
	page := queryInt(r, "page", 1)
	limit := queryInt(r, "limit", defaultLimit)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	return Params{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
}

func queryInt(r *http.Request, key string, fallback int) int {
	value := r.URL.Query().Get(key)
	if value == "" {
		return fallback
	}

	num, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return num
}
