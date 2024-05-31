package request

type PageInfo struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" from:"pageSize"`
	Keyword  string `json:"keyword" from:"keyword"`
}

type GetById struct {
	ID uint `json:"id" form:"id"`
}
