package utils

import "time"

// 获取时间并返回时间戳（如：20250429）
func GetTimePath() string {
	now := time.Now()
	formattedDate := now.Format("20060102_150405")
	// fmt.Println(formattedDate)
	return formattedDate
}
