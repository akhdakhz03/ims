package parameterforms

//define parameter variable

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
