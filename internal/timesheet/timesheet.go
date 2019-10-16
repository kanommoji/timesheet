package timesheet

import "timesheet/internal/model"

type TimesheetGateways interface {
}

type Timesheet struct {
}

func (timesheet Timesheet) CalculatePaymentSummary(member []model.Member, incomes []model.Incomes) []model.TransactionTimesheet {
	var transactionTimesheetList []model.TransactionTimesheet
	for _, member := range member {
		coachingPaymentRate := CalculateCoachingPaymentRate(incomes, member.Company)
		trainingWage := CalculateTrainingWage(incomes, member.Company)
		otherWage := CalculateOtherWage(incomes, member.Company)
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

func CalculateCoachingPaymentRate(incomes []model.Incomes, company string) float64 {
	return 0.00
}

func CalculateTrainingWage(incomes []model.Incomes, company string) float64 {
	if company == "shuhari" {
		return 20000.00
	}
	return 155000.00
}

func CalculateOtherWage(incomes []model.Incomes, company string) float64 {
	return 0.00
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
