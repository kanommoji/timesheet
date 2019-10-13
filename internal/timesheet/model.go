package timesheet

type TransactionTimesheet struct {
	MemberID               string  `json:"member_id"`
	MemberNameTH           string  `json:"member_name_th"`
	Month                  int     `json:"month"`
	Year                   int     `json:"year"`
	Company                string  `json:"company"`
	Coaching               float64 `json:"coaching"`
	Training               float64 `json:"training"`
	Other                  float64 `json:"other"`
	TotalIncomes           float64 `json:"total_incomes"`
	Salary                 float64 `json:"salary"`
	IncomeTax1             float64 `json:"income_tax_1"`
	SocialSecurity         float64 `json:"social_security"`
	NetSalary              float64 `json:"net_salary"`
	Wage                   float64 `json:"wage"`
	IncomeTax53Percentage  int     `json:"income_tax_53_percentage"`
	IncomeTax53            float64 `json:"income_tax_53"`
	NetWage                float64 `json:"net_wage"`
	NetTransfer            float64 `json:"net_transfer"`
	StatusCheckingTransfer string  `json:"status_checking_transfer"`
	DateTransfer           string  `json:"date_transfer"`
	Comment                string  `json:"comment"`
}