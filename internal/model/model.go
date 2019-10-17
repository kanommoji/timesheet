package model

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
	ID                       int    `json:"id"`
	MemberID                 string `json:"member_id"`
	Month                    int    `json:"month"`
	Year                     int    `json:"year"`
	Day                      int    `json:"day"`
	StartTimeAMHours         int    `json:"start_time_am_hours"`
	StartTimeAMMinutes       int    `json:"start_time_am_minutes"`
	StartTimeAMSeconds       int    `json:"start_time_am_seconds"`
	EndTimeAMHours           int    `json:"end_time_am_hours"`
	EndTimeAMMinutes         int    `json:"end_time_am_minutes"`
	EndTimeAMSeconds         int    `json:"end_time_am_seconds"`
	StartTimePMHours         int    `json:"start_time_pm_hours"`
	StartTimePMMinutes       int    `json:"start_time_pm_minutes"`
	StartTimePMSeconds       int    `json:"start_time_pm_seconds"`
	EndTimePMHours           int    `json:"end_time_pm_hours"`
	EndTimePMMinutes         int    `json:"end_time_pm_minutes"`
	EndTimePMSeconds         int    `json:"end_time_pm_seconds"`
	Overtime                 int    `json:"overtime"`
	TotalHoursHours          int    `json:"total_hours_hours"`
	TotalHoursMinutes        int    `json:"total_hours_minutes"`
	TotalHoursSeconds        int    `json:"total_hours_seconds"`
	CoachingCustomerCharging int    `json:"coaching_customer_charging"`
	CoachingPaymentRate      int    `json:"coaching_payment_rate"`
	TrainingWage             int    `json:"training_wage"`
	OtherWage                int    `json:"other_wage"`
	Company                  string `json:"company"`
	Description              string `json:"description"`
}

type Payment struct {
	TotalHoursHours               int
	TotalHoursMinutes             int
	TotalHoursSeconds             int
	TotalCoachingCustomerCharging float64
	TotalCoachingPaymentRate      float64
	TotalTrainigWage              float64
	TotalOtherWage                float64
	PaymentWage                   float64
}

type Member struct {
	ID                    int     `json:"id"`
	MemberID              string  `json:"member_id"`
	Company               string  `json:"company"`
	MemberNameTH          string  `json:"member_name_th"`
	MemberNameENG         string  `json:"member_name_eng"`
	Email                 string  `json:"email"`
	OvertimeRate          float64 `json:"overtime_rate"`
	RatePerDay            float64 `json:"rate_per_day"`
	RatePerHour           float64 `json:"rate_per_hour"`
	Salary                float64 `json:"salary"`
	IncomeTax1            float64 `json:"income_tax_1"`
	SocialSecurity        float64 `json:"social_security"`
	IncomeTax53Percentage int     `json:"income_tax_53_percentage"`
	Status                string  `json:"status"`
	TravelExpense         float64 `json:"travel_expense"`
}

type Time struct {
	Hours   int
	Minutes int
	Seconds int
}
