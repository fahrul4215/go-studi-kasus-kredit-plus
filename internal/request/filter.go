package request

// GeneralFilter is a reusable filter for common query parameters
type QueryOptions struct {
	Keyword      string   `form:"keyword"`                  // Search keyword
	Columns      []string `form:"-"`                        // Columns to apply keyword search
	Page         int      `form:"page" default:"1"`         // Current page (default: 1)
	PerPage      int      `form:"per_page" default:"10"`    // Items per page (default: 10)
	SortBy       string   `form:"sort_by" default:"id"`     // Column to sort by
	SortOrder    string   `form:"sort_order" default:"asc"` // Sort order (asc/desc)
	WhereClauses []string `form:"-"`
	GroupBy      []string `form:"-"` // Columns to group by
}

func (f *QueryOptions) SetDefaultQueryOptions() {
	if f.Page == 0 {
		f.Page = 1
	}
	if f.PerPage == 0 {
		f.PerPage = 10
	}
	// if f.SortBy == "" {
	// 	f.SortBy = "id"
	// }
	// if f.SortOrder == "" {
	// 	f.SortOrder = "asc"
	// }
}
