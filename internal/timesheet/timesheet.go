package timesheet

import "timesheet/internal/model"

type TimesheetGateways interface {
}

type Timesheet struct {
}

func (timesheet Timesheet) CalculatePaymentSummary(member []model.Member, incomes []model.Incomes) []model.TransactionTimesheet {
	var transactionTimesheetList []model.TransactionTimesheet
	for _, member := range member {
		coachingPaymentRate := CalculateTotalCoachingPaymentRate(incomes, member.Company)
		trainingWage := CalculateTotalTrainingWage(incomes, member.Company)
		otherWage := CalculateTotalOtherWage(incomes, member.Company, member.TravelExpense)
		paymentWage := CalculateTotalPaymentWage(coachingPaymentRate, trainingWage, otherWage)
		netSalary := CalculateNetSalary(member.Salary, member.IncomeTax1, member.SocialSecurity)
		wage := CalculateWage(paymentWage, member.Salary, member.Status)
		incomeTax53 := CalculateIncomeTax53(wage, member.IncomeTax53Percentage)
		netWage := CalculateNetWage(member.IncomeTax53Percentage, paymentWage, member.Salary, member.Status)
		netTransfer := CalculateNetTransfer(netSalary, netWage)
		transactionTimesheet := model.TransactionTimesheet{
			MemberID:              member.MemberID,
			MemberNameTH:          member.MemberNameTH,
			Company:               member.Company,
			Coaching:              coachingPaymentRate,
			Training:              trainingWage,
			Other:                 otherWage,
			TotalIncomes:          paymentWage,
			Salary:                member.Salary,
			IncomeTax1:            member.IncomeTax1,
			SocialSecurity:        member.SocialSecurity,
			NetSalary:             netSalary,
			Wage:                  wage,
			IncomeTax53Percentage: member.IncomeTax53Percentage,
			IncomeTax53:           incomeTax53,
			NetWage:               netWage,
			NetTransfer:           netTransfer,
		}
		transactionTimesheetList = append(transactionTimesheetList, transactionTimesheet)
	}

	return transactionTimesheetList
}

func CalculateTotalPaymentWage(coachingPaymentRate, trainingWage, otherWage float64) float64 {
	return coachingPaymentRate + trainingWage + otherWage
}

func CalculateNetSalary(salary, incomeTax1, socialSecurity float64) float64 {
	return salary - incomeTax1 - socialSecurity
}

func CalculateNetWage(incomeTax53Percentage int, paymentWage, salary float64, status string) float64 {
	wage := CalculateWage(paymentWage, salary, status)
	incomeTax53 := CalculateIncomeTax53(wage, incomeTax53Percentage)
	return wage - incomeTax53
}

func CalculateWage(paymentWage, salary float64, status string) float64 {
	if status != "salary" {
		return paymentWage - salary
	}
	return paymentWage
}

func CalculateIncomeTax53(wage float64, incomeTax53Percentag int) float64 {
	return wage * (float64(incomeTax53Percentag) / 100)
}

func CalculateNetTransfer(netSalary, netWage float64) float64 {
	return netSalary + netWage
}

func CalculateTotalCoachingCustomerCharging(incomes []model.Incomes, company string) float64 {
	totalCoachingCustomerCharging := 0
	if company == "siam_chamnankit" {
		for i := range incomes {
			if incomes[i].Company == "siam_chamnankit" {
				totalCoachingCustomerCharging += incomes[i].CoachingCustomerCharging
			}
		}
		return float64(totalCoachingCustomerCharging)
	}
	if company == "shuhari" {
		for i := range incomes {
			if incomes[i].Company == "shuhari" {
				totalCoachingCustomerCharging += incomes[i].CoachingCustomerCharging
			}
		}
		return float64(totalCoachingCustomerCharging)
	}
	return 0
}

func CalculateTotalOtherWage(incomes []model.Incomes, company string, travelExpense float64) float64 {
	totalOtherWage := 0
	if company == "siam_chamnankit" {
		for i := range incomes {
			if incomes[i].Company == "siam_chamnankit" {
				totalOtherWage += incomes[i].OtherWage
			}
		}
		return float64(totalOtherWage) + travelExpense
	}
	if company == "shuhari" {
		for i := range incomes {
			if incomes[i].Company == "shuhari" {
				totalOtherWage += incomes[i].OtherWage
			}
		}
		return float64(totalOtherWage) + travelExpense
	}
	return 0
}

func CalculateTotalCoachingPaymentRate(incomes []model.Incomes, company string) float64 {
	totalCoachingPaymentRate := 0
	if company == "siam_chamnankit" {
		for i := range incomes {
			if incomes[i].Company == "siam_chamnankit" {
				totalCoachingPaymentRate += incomes[i].CoachingPaymentRate
			}
		}
		return float64(totalCoachingPaymentRate)
	}
	if company == "shuhari" {
		for i := range incomes {
			if incomes[i].Company == "shuhari" {
				totalCoachingPaymentRate += incomes[i].CoachingPaymentRate
			}
		}
		return float64(totalCoachingPaymentRate)
	}
	return 0
}

func CalculateTotalTrainingWage(incomes []model.Incomes, company string) float64 {
	totalCoachingTrainingWage := 0
	if company == "siam_chamnankit" {
		for i := range incomes {
			if incomes[i].Company == "siam_chamnankit" {
				totalCoachingTrainingWage += incomes[i].TrainingWage
			}
		}
		return float64(totalCoachingTrainingWage)
	}

	for i := range incomes {
		if incomes[i].Company == "shuhari" {
			totalCoachingTrainingWage += incomes[i].TrainingWage
		}
	}
	return float64(totalCoachingTrainingWage)

}
