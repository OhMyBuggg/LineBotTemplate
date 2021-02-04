package content

import (
	"strconv"
)

const (
	Status = "我還醒著"
	Usage = "蟲蟲機器人使用說明書\n目前機器人的活動時間在上午十點到晚上十一點，若在其他時間使用，會稍有延遲\n\n目前功能:\n	 1. 測試 => 回覆訊息資訊\n  2. 蟲蟲機器人 => 回覆狀態\n  3. 重複 \"自訂訊息\" => 回覆相同訊息\n\n施工中..."
)

func GetTestMessage(ReplyToken string, MessageID string, MessageText string, UserID string, QuotaValue int64) string {
	return "EventReplyToken: "+ReplyToken+"\nMessageID: "+MessageID+"\nMessageText: "+MessageText+"\nUserID: "+UserID+"\nremain message: "+strconv.FormatInt(QuotaValue, 10)
}