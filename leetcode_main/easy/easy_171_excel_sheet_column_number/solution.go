package easy_171_excel_sheet_column_number

//https://leetcode.cn/problems/excel-sheet-column-number/

func titleToNumber(columnTitle string) int {
	var result int = 0
	length := len(columnTitle)
	for index := 0; index < length; index++ {
		tmpNum := int(columnTitle[index] - 'A' + 1)
		result = result*26 + tmpNum
	}
	return result
}
