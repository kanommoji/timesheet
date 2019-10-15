package timesheet

type TransactionTimesheet struct {
	ID                     int     `json:"id"`
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

type Incomes struct {
	Day                      int    `json:"day"`
	StartTimeAM              string `json:"start_time_am"`
	EndTimeAM                string `json:"end_time_am"`
	StartTimePM              string `json:"start_time_pm"`
	EndTimePM                string `json:"end_time_pm"`
	Overtime                 int    `json:"overtime"`
	TotalHours               int    `json:"total_hours"`
	CoachingCustomerCharging int    `json:"coaching_customer_charging"`
	CoachingPaymentRate      int    `json:"coaching_payment_rate"`
	TrainingWage             int    `json:"training_wage"`
	OtherWage                int    `json:"other_wage"`
	Company                  string `json:"company"`
	Description              string `json:"description"`
}
