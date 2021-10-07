package request

type TableInsert struct {
	Capacity int  `json:"capacity"`
	Status   bool `json:"status"`
}
