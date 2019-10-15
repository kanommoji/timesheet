*** Settings ***
Library    SeleniumLibrary

***Variables***
${URL_PAYMENTS}    http://localhost:3000/showSummaryTimesheet
${URL_TIMESHEET_BY_ID}    http://localhost:3000/showTimesheetByID

*** Testcases ***
เพิ่ม timesheet และดูผลสรุปในหน้า PAYMENTS
    เปิด Browser

***Keywords***
เปิด Browser 
    Open Browser    ${URL_PAYMENTS}    chrome
