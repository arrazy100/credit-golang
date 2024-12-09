package response

type ListDebtorInstallmentResponse struct {
	Data  []DebtorInstallmentResponse `json:"data"`
	Total int                         `json:"total"`
}

type DebtorInstallmentResponse struct {
	ID                    string                          `json:"id"`
	DebtorTransaction     DebtorTransactionResponse       `json:"debtor_transaction"`
	TenorDuration         int                             `json:"tenor_duration"`
	TotalInstallmentCount int                             `json:"total_installment_count"`
	PaidInstallmentCount  int                             `json:"paid_installment_count"`
	MonthlyInstallment    string                          `json:"monthly_installment"`
	StartDatePeriod       string                          `json:"start_date_period"`
	EndDatePeriod         string                          `json:"end_date_period"`
	Lines                 []DebtorInstallmentLineResponse `json:"lines"`
}

type DebtorInstallmentLineResponse struct {
	ID                string  `json:"id"`
	DueDate           string  `json:"due_date"`
	InstallmentNumber int     `json:"installment_number"`
	InstallmentAmount string  `json:"installment_amount"`
	PaymentDate       *string `json:"payment_date"`
	Status            string  `json:"status"`
}

type ByInstallmentNumber []DebtorInstallmentLineResponse

func (a ByInstallmentNumber) Len() int { return len(a) }
func (a ByInstallmentNumber) Less(i, j int) bool {
	return a[i].InstallmentNumber < a[j].InstallmentNumber
}
func (a ByInstallmentNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
