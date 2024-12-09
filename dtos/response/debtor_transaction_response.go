package response

type DebtorTransactionResponse struct {
	ID             string `json:"id"`
	ContractNumber string `json:"contract_number"`
	OTR            string `json:"otr"`
	AdminFee       string `json:"admin_fee"`
	TotalLoan      string `json:"total_loan"`
	TotalInterest  string `json:"total_interest"`
	AssetName      string `json:"asset_name"`
	Status         string `json:"status"`
}
