package masak_openapi

type Domain struct {
	Name              string
	Step              string
	EstimateInMinutes string
}

type Repository interface {
	GetMenuApi(keyName string) (Domain, error)
}
