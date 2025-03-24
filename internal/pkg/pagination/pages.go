// Package pagination provides support for pagination requests and responses.
package pagination

import (
	"fmt"
	"go-studi-kasus-kredit-plus/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

var (
	// DefaultLimit specifies the default page size
	DefaultLimit = 10
	// MaxPageSize specifies the maximum page size
	MaxPageSize = 1000
	// PageVar specifies the query parameter name for page number
	PageVar = "page"
	// PageSizeVar specifies the query parameter name for page size
	PageSizeVar = "limit"
	KeywordVar  = "keyword"
	Sort        = "sort"
	Order       = "order"

	// DefaultSort specifies the default sort column
	DefaultSort = "id"

	// Ascending specifies the ascending order
	Ascending = "asc"
	// Descending specifies the descending order
	Descending = "desc"

	// DefaultSortOrder specifies the default sort order
	DefaultSortOrder = Ascending
)

// Pages represents a paginated list of data items.
type Pages struct {
	Page       int    `json:"page" form:"page"`
	Limit      int    `json:"limit" form:"limit"`
	Sort       string `json:"sort" form:"sort"`
	Order      string `json:"order" form:"order"`
	PageCount  int    `json:"page_count"`
	TotalCount int    `json:"total_count"`
	Keyword    string `form:"keyword" json:"-"`
}

// New creates a new Pages instance.
// The page parameter is 1-based and refers to the current page index/number.
// The limit parameter refers to the number of items on each page.
// And the total parameter specifies the total number of data items.
// If total is less than 0, it means total is unknown.
func New(page, limit, total int, sort, order, keyword string) *Pages {
	if limit <= 0 {
		limit = DefaultLimit
	}
	if limit > MaxPageSize {
		limit = MaxPageSize
	}
	pageCount := -1
	if page < 1 {
		page = 1
	}

	return &Pages{
		Page:       page,
		Limit:      limit,
		TotalCount: total,
		PageCount:  pageCount,
		Sort:       sort,
		Order:      order,
	}
}

// NewFromRequest creates a Pages object using the query parameters found in the given HTTP request.
// count stands for the total number of items. Use -1 if this is unknown.
func NewFromRequest(req *http.Request) *Pages {
	keyword := req.URL.Query().Get(KeywordVar)
	page := parseInt(req.URL.Query().Get(PageVar), 1)
	limit := parseInt(req.URL.Query().Get(PageSizeVar), DefaultLimit)
	sort := req.URL.Query().Get(Sort)
	if sort == "" {
		sort = "id"
	}

	order := req.URL.Query().Get(Order)
	if order != Ascending && order != Descending {
		order = Ascending
	}

	return New(page, limit, 0, sort, order, keyword)
}

func NewFromRequestWithTotal(req *http.Request, total int) *Pages {
	keyword := req.URL.Query().Get(KeywordVar)
	page := parseInt(req.URL.Query().Get(PageVar), 1)
	limit := parseInt(req.URL.Query().Get(PageSizeVar), DefaultLimit)
	sort := req.URL.Query().Get(Sort)
	if sort == "" {
		sort = "id"
	}

	order := req.URL.Query().Get(Order)
	if order != Ascending && order != Descending {
		order = Ascending
	}

	data := New(page, limit, total, sort, order, keyword)
	data.UpdateTotal(total)
	return data
}

// parseInt parses a string into an integer. If parsing is failed, defaultValue will be returned.
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// Offset returns the OFFSET value that can be used in a SQL statement.
func (p *Pages) Offset() int {
	page := p.Page
	if page < 1 {
		page = 1
	}

	return (page - 1) * p.GetLimit()
}

func (p *Pages) GetLimit() int {
	limit := p.Limit
	if limit == 0 {
		limit = DefaultLimit
	}
	return limit
}

// OrderDB returns the formatted string that can be used in a SQL statement for ordering.
func (p *Pages) OrderDB() string {
	sort := p.Sort
	if sort == "" {
		sort = "id"
	}

	order := p.Order
	if order != Ascending && order != Descending {
		order = Ascending
	}

	return fmt.Sprintf("%s %s", utils.EscapeSpecial(sort), utils.EscapeSpecial(order))
}

// UpdateTotal update current total to use in pagination response
func (p *Pages) UpdateTotal(total int) {
	if total >= 0 {
		p.PageCount = (total + p.Limit - 1) / p.Limit
		p.TotalCount = total
		if p.Page > p.PageCount {
			p.Page = p.PageCount
		}
	}
}

// BuildLinkHeader returns an HTTP header containing the links about the pagination.
func (p *Pages) BuildLinkHeader(baseURL string, defaultLimit int) string {
	links := p.BuildLinks(baseURL, defaultLimit)
	header := ""
	if links[0] != "" {
		header += fmt.Sprintf("<%v>; rel=\"first\", ", links[0])
		header += fmt.Sprintf("<%v>; rel=\"prev\"", links[1])
	}
	if links[2] != "" {
		if header != "" {
			header += ", "
		}
		header += fmt.Sprintf("<%v>; rel=\"next\"", links[2])
		if links[3] != "" {
			header += fmt.Sprintf(", <%v>; rel=\"last\"", links[3])
		}
	}
	return header
}

// BuildLinks returns the first, prev, next, and last links corresponding to the pagination.
// A link could be an empty string if it is not needed.
// For example, if the pagination is at the first page, then both first and prev links
// will be empty.
func (p *Pages) BuildLinks(baseURL string, defaultLimit int) [4]string {
	var links [4]string
	pageCount := p.PageCount
	page := p.Page
	if pageCount >= 0 && page > pageCount {
		page = pageCount
	}
	if strings.Contains(baseURL, "?") {
		baseURL += "&"
	} else {
		baseURL += "?"
	}
	if page > 1 {
		links[0] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, 1)
		links[1] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, page-1)
	}
	if pageCount >= 0 && page < pageCount {
		links[2] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, page+1)
		links[3] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, pageCount)
	} else if pageCount < 0 {
		links[2] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, page+1)
	}
	if limit := p.Limit; limit != defaultLimit {
		for i := 0; i < 4; i++ {
			if links[i] != "" {
				links[i] += fmt.Sprintf("&%v=%v", PageSizeVar, limit)
			}
		}
	}

	return links
}
