package model

type InstallmentPlan struct {
	Date string  `json:"date"`
	SumToPay float64 `json:"sum_to_pay"`
	ParticipationFee float64 `json:"participation_fee"`
}
