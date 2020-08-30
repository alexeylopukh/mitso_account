package model

type Student struct {
	Login int64  `json:"login"`
	Fio   string `json:"fio"`
	Balance float64 `json:"balance"`
	Principal float64 `json:"principal"`
	Fine float64 `json:"fine"`
	LastUpdate string `json:"last_update"`
	Prolongation int64 `json:"prolongation"`
	Plans []*InstallmentPlan `json:"plans"`
}
