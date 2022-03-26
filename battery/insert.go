package battery

import "fmt"

func insert_pack(resp_float [37]float32) (str_result [31]string) {
	var result [37]float32

	//int16
	result[0] = resp_float[0] / 100

	//uint16
	result[1] = resp_float[1] / 100
	result[2] = resp_float[2]
	result[3] = resp_float[3]
	result[4] = resp_float[4] / 100
	result[5] = resp_float[5] / 100
	result[6] = resp_float[6] / 100
	result[7] = resp_float[7]

	//uint16
	result[8] = resp_float[15] / 1000
	result[9] = resp_float[16] / 1000
	result[10] = resp_float[17] / 1000
	result[11] = resp_float[18] / 1000
	result[12] = resp_float[19] / 1000
	result[13] = resp_float[20] / 1000
	result[14] = resp_float[21] / 1000
	result[15] = resp_float[22] / 1000
	result[16] = resp_float[23] / 1000
	result[17] = resp_float[24] / 1000
	result[18] = resp_float[25] / 1000
	result[19] = resp_float[26] / 1000
	result[20] = resp_float[27] / 1000
	result[21] = resp_float[28] / 1000
	result[22] = resp_float[29] / 1000
	result[23] = resp_float[30] / 1000

	//int16
	result[24] = resp_float[31] / 10
	result[25] = resp_float[32] / 10
	result[26] = resp_float[33] / 10
	result[27] = resp_float[34] / 10
	result[28] = resp_float[35] / 10
	result[29] = resp_float[36] / 10

	//int16
	result[30] = resp_float[12]

	str_result[0] = fmt.Sprintf("%.2f", result[0])   //batcur
	str_result[1] = fmt.Sprintf("%.2f", result[1])   //batvolt
	str_result[2] = fmt.Sprintf("%.2f", result[2])   //soc
	str_result[3] = fmt.Sprintf("%.2f", result[3])   //soh
	str_result[4] = fmt.Sprintf("%.2f", result[4])   //rem_cap
	str_result[5] = fmt.Sprintf("%.2f", result[5])   //full_cap
	str_result[6] = fmt.Sprintf("%.2f", result[6])   //des_cap
	str_result[7] = fmt.Sprintf("%.2f", result[7])   //cycle_count
	str_result[8] = fmt.Sprintf("%.2f", result[8])   //cell1
	str_result[9] = fmt.Sprintf("%.2f", result[9])   //cell2
	str_result[10] = fmt.Sprintf("%.2f", result[10]) //cell3
	str_result[11] = fmt.Sprintf("%.2f", result[11]) //cell4
	str_result[12] = fmt.Sprintf("%.2f", result[12]) //cell5
	str_result[13] = fmt.Sprintf("%.2f", result[13]) //cell6
	str_result[14] = fmt.Sprintf("%.2f", result[14]) //cell7
	str_result[15] = fmt.Sprintf("%.2f", result[15]) //cell8
	str_result[16] = fmt.Sprintf("%.2f", result[16]) //cell9
	str_result[17] = fmt.Sprintf("%.2f", result[17]) //cell10
	str_result[18] = fmt.Sprintf("%.2f", result[18]) //cell11
	str_result[19] = fmt.Sprintf("%.2f", result[19]) //cell12
	str_result[20] = fmt.Sprintf("%.2f", result[20]) //cell13
	str_result[21] = fmt.Sprintf("%.2f", result[21]) //cell14
	str_result[22] = fmt.Sprintf("%.2f", result[22]) //cell15
	str_result[23] = fmt.Sprintf("%.2f", result[23]) //cell16
	str_result[24] = fmt.Sprintf("%.2f", result[24]) //temp1
	str_result[25] = fmt.Sprintf("%.2f", result[25]) //temp2
	str_result[26] = fmt.Sprintf("%.2f", result[26]) //temp3
	str_result[27] = fmt.Sprintf("%.2f", result[27]) //temp4
	str_result[28] = fmt.Sprintf("%.2f", result[28]) //temp_mossf
	str_result[29] = fmt.Sprintf("%.2f", result[29]) //temp_env
	str_result[30] = fmt.Sprintf("%.2f", result[30]) //Balance status

	return str_result
}
