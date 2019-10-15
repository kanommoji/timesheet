function showSummary(){
    var year = $("#year").val();
    var month = $("#month").val()
    
    var request = new XMLHttpRequest();
    var url = "http://localhost:8080/showSummaryTimesheet";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
            var json = JSON.parse(request.responseText);
            var siamChamnankit = "";
            var countSiamChamnankit = 0 ;
            var totalCoachingSiamChamnankit = 0;
            var totalTraningSiamChamnankit = 0;
            var totalOtherSiamChamnankit = 0;
            var totalIncomesSiamChamnankit = 0;
            var totalSalarySiamChamnankit = 0;
            var totalIncomeTax1SiamChamnankit = 0;
            var totalSocialSecuritySiamChamnankit = 0;
            var totalNetSalarySiamChamnankit = 0;
            var totalWageSiamChamnankit = 0;
            var totalIncomeTax53SiamChamnankit = 0;
            var totalNetWageSiamChamnankit = 0;
            var totalNetTransferSiamChamnankit = 0;

            var shuhari = "";
            var countShuhari = 0;
            var totalCoachingShuhari = 0;
            var totalTraningShuhari = 0;
            var totalOtherShuhari = 0;
            var totalIncomesShuhari = 0;
            var totalSalaryShuhari = 0;
            var totalIncomeTax1Shuhari = 0;
            var totalSocialSecurityShuhari = 0;
            var totalNetSalaryShuhari = 0;
            var totalWageShuhari = 0;
            var totalIncomeTax53Shuhari = 0;
            var totalNetWageShuhari = 0;
            var totalNetTransferShuhari = 0;

            for (var i = 1; i <= json.length; i++) {  
                if(json[i-1].company === "Siam Chamnankit"){
                    countSiamChamnankit++;
                    siamChamnankit += "<tr>";
                    siamChamnankit += "<td>"+countSiamChamnankit+"</td>";
                    siamChamnankit += "<td>"+json[i-1].member_name_th+"</td>";
                    siamChamnankit += "<td>"+json[i-1].coaching+"</td>";
                    totalCoachingSiamChamnankit += json[i-1].coaching;
                    siamChamnankit += "<td>"+json[i-1].training+"</td>";
                    totalTraningSiamChamnankit +=json[i-1].training;
                    siamChamnankit += "<td>"+json[i-1].other+"</td>";
                    totalOtherSiamChamnankit += json[i-1].other;
                    siamChamnankit += "<td>"+json[i-1].total_incomes+"</td>";
                    totalIncomesSiamChamnankit += json[i-1].total_incomes;
                    siamChamnankit += "<td>"+json[i-1].salary+"</td>";
                    totalSalarySiamChamnankit += json[i-1].salary;
                    siamChamnankit += "<td>"+json[i-1].income_tax_1+"</td>";
                    totalIncomeTax1SiamChamnankit += json[i-1].income_tax_1;
                    siamChamnankit += "<td>"+json[i-1].social_security+"</td>";
                    totalSocialSecuritySiamChamnankit += json[i-1].social_security;
                    siamChamnankit += "<td>"+json[i-1].net_salary+"</td>";
                    totalNetSalarySiamChamnankit += json[i-1].net_salary;
                    siamChamnankit += "<td>"+json[i-1].wage+"</td>";
                    totalWageSiamChamnankit += json[i-1].wage;
                    siamChamnankit += "<td>"+json[i-1].income_tax_53_percentage+"</td>";
                    siamChamnankit += "<td>"+json[i-1].income_tax_53+"</td>";
                    totalIncomeTax53SiamChamnankit += json[i-1].income_tax_53;
                    siamChamnankit += "<td>"+json[i-1].net_wage+"</td>";
                    totalNetWageSiamChamnankit +=json[i-1].net_wage;
                    siamChamnankit += "<td>"+json[i-1].net_transfer+"</td>";
                    totalNetTransferSiamChamnankit += json[i-1].net_transfer;
                    siamChamnankit += "<td>"+json[i-1].status_checking_transfer+"</td>";
                    siamChamnankit += "<td>"+json[i-1].date_transfer+"</td>";
                    siamChamnankit += "<td>"+json[i-1].comment+"</td>";
                    siamChamnankit += "</tr>";
                }else{
                    countShuhari++;
                    shuhari += "<tr>";
                    shuhari += "<td>"+countShuhari+"</td>";
                    shuhari += "<td>"+json[i-1].member_name_th+"</td>";
                    shuhari += "<td>"+json[i-1].coaching+"</td>";
                    totalCoachingShuhari += json[i-1].coaching;
                    shuhari += "<td>"+json[i-1].training+"</td>";
                    totalTraningShuhari += json[i-1].training;
                    shuhari += "<td>"+json[i-1].other+"</td>";
                    totalOtherShuhari += json[i-1].other;
                    shuhari += "<td>"+json[i-1].total_incomes+"</td>";
                    totalIncomesShuhari += json[i-1].total_incomes;
                    shuhari += "<td>"+json[i-1].salary+"</td>";
                    totalSalaryShuhari += json[i-1].salary;
                    shuhari += "<td>"+json[i-1].income_tax_1+"</td>";
                    totalIncomeTax1Shuhari += json[i-1].income_tax_1;
                    shuhari += "<td>"+json[i-1].social_security+"</td>";
                    totalSocialSecurityShuhari += json[i-1].social_security;
                    shuhari += "<td>"+json[i-1].net_salary+"</td>";
                    totalNetSalaryShuhari += json[i-1].net_salary;
                    shuhari += "<td>"+json[i-1].wage+"</td>";
                    totalWageShuhari += json[i-1].wage;
                    shuhari += "<td>"+json[i-1].income_tax_53_percentage+"</td>";
                    shuhari += "<td>"+json[i-1].income_tax_53+"</td>";
                    totalIncomeTax53Shuhari += json[i-1].income_tax_53;
                    shuhari += "<td>"+json[i-1].net_wage+"</td>";
                    totalNetWageShuhari += json[i-1].net_wage;
                    shuhari += "<td>"+json[i-1].net_transfer+"</td>";
                    totalNetTransferShuhari += json[i-1].net_transfer;
                    shuhari += "<td>"+json[i-1].status_checking_transfer+"</td>";
                    shuhari += "<td>"+json[i-1].date_transfer+"</td>";
                    shuhari += "<td>"+json[i-1].comment+"</td>";
                    shuhari += "</tr>";
                }
            }
            $("#table_siam_chamnankit").html(siamChamnankit);
            $("#total_coaching_siam_chamnankit").html(totalCoachingSiamChamnankit);
            $("#total_traning_siam_chamnankit").html(totalTraningSiamChamnankit);
            $("#total_other_siam_chamnankit").html(totalOtherSiamChamnankit);
            $("#total_incomes_siam_chamnankit").html(totalIncomesSiamChamnankit);
            $("#total_salary_siam_chamnankit").html(totalSalarySiamChamnankit);
            $("#total_income_tax_1_siam_chamnankit").html(totalIncomeTax1SiamChamnankit);
            $("#total_social_security_siam_chamnankit").html(totalSocialSecuritySiamChamnankit);
            $("#total_net_salary_siam_chamnankit").html(totalNetSalarySiamChamnankit);
            $("#total_wage_siam_chamnankit").html(totalWageSiamChamnankit);
            $("#total_income_tax_53_siam_chamnankit").html(totalIncomeTax53SiamChamnankit);
            $("#total_net_wage_siam_chamnankit").html(totalNetWageSiamChamnankit);
            $("#total_net_transfer_siam_chamnankit").html(totalNetTransferSiamChamnankit);

            $("#table_shuhari").html(shuhari);
            $("#total_coaching_shuhari").html(totalCoachingShuhari);
            $("#total_traning_shuhari").html(totalTraningShuhari);
            $("#total_other_shuhari").html(totalOtherShuhari);
            $("#total_incomes_shuhari").html(totalIncomesShuhari);
            $("#total_salary_shuhari").html(totalSalaryShuhari);
            $("#total_income_tax_1_shuhari").html(totalIncomeTax1Shuhari);
            $("#total_social_security_shuhari").html(totalSocialSecurityShuhari);
            $("#total_net_salary_shuhari").html(totalNetSalaryShuhari);
            $("#total_wage_shuhari").html(totalWageShuhari);
            $("#total_income_tax_53_shuhari").html(totalIncomeTax53Shuhari);
            $("#total_net_wage_shuhari").html(totalNetWageShuhari);
            $("#total_net_transferShuhari").html(totalNetTransferShuhari);
            
        }
    };
    var data = JSON.stringify({"year":year, "month": month});
    request.send(data);
    
}

function addTimsheet(){
    
}