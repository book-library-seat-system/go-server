package mylog

// // AddLog : 添加记录
// func AddLog(user string, command string, oldStr string, newStr string) {
// 	file := getFileHandle()
// 	l := log.New(file, "[INFO]", log.Ltime)
// 	outStr := ""
// 	if user != "" {
// 		outStr += "User:" + user + "  "
// 	}
// 	if command != "" {
// 		outStr += "Command:" + command + "\n"
// 	}
// 	if oldStr != "" {
// 		outStr += "From:" + oldStr + "\n"
// 	}
// 	if newStr != "" {
// 		outStr += "To:" + newStr + "\n"
// 	}
// 	l.Print(outStr)
// 	file.Close()
// }

// // AddErr 添加错误
// func AddErr(err error) {
// 	file := getFileHandle()
// 	l := log.New(file, "[INFO]", log.Ltime)
// 	l.Println(err)
// 	file.Close()
// }
