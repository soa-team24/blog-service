package model

type PagedResultBlog struct {
	Results    []Blog `json:"results"`
	TotalCount int    `json:"totalCount"`
}
