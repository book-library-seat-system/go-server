/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: seat的逻辑层，使用dao层接口，为上层控制层（server层）提供接口
	PS：该文件所有错误全都panic抛出，不进行错误返回
Date: 2018年5月4日 星期五 上午10:52
****************************************************************************/

package seat

/*************************************************
Function: GetAllTimeInterval
Description: 得到所有的时间间隔
InputParameter:
	school: 所查询的学校名字
Return: 可用时间间隔数组，以一小时为单位
*************************************************/
func GetAllTimeInterval(school string) []TimeInterval {
	return nil
}

/*************************************************
Function: GetAllSeatinfo
Description: 得到某时间段所有座位的信息，数组下标代表位置
InputParameter:
	school: 所查询的学校名字
	timeinterval: 查询的时间戳
Return: 该时间段的座位预约信息，用int数组保存
*************************************************/
func GetAllSeatinfo(school string, timeinterval TimeInterval) []int {
	return nil
}

/*************************************************
Function: GetAllUnbookSeatNumber
Description: 得到某时间段的未预约座位数量
InputParameter:
	school: 所查询的学校名字
	timeinterval: 查询的时间戳
Return: 未预约的座位数量
*************************************************/
func GetAllUnbookSeatNumber(school string, timeinterval TimeInterval) int {
	return 0
}

/*************************************************
Function: BookSeat
Description: 预约座位
InputParameter:
	school: 所查询的学校名字
	timeinterval: 查询的时间戳
	studentid: 预约学生ID
	seatid: 座位ID，即数组下标
Return: none
*************************************************/
func BookSeat(school string, timeinterval TimeInterval, studentid string, seatid string) {

}

/*************************************************
Function: UnbookSeat
Description: 取消预约座位
InputParameter:
	school: 所查询的学校名字
	timeinterval: 查询的时间戳
	studentid: 预约学生ID
	seatid: 座位ID，即数组下标
Return: none
*************************************************/
func UnbookSeat(school string, timeinterval TimeInterval, studentid string, seatid string) {

}
