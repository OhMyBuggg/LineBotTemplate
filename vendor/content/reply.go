package content

import (
	"strconv"
)

const (
	Status = "我還醒著"
	Usage = "蟲蟲機器人使用說明書\n目前機器人的活動時間在上午十點到晚上十一點，若在其他時間使用，會稍有延遲\n\n目前功能:\n  1. 測試 => 回覆訊息資訊\n  2. 蟲蟲機器人 => 回覆狀態\n  3. 重複 \"自訂訊息\" => 回覆相同訊息\n  4. 大頭貼 熊大/兔兔/饅頭人 => 可以更換傳送訊息時的頭貼\n  5. youtube \"影片ID\" => 回覆影片標題、頻道及下載連結（此功能尚未完全，有可能因不同影片而出錯）\n\n施工中..."
	StickerUsage = "很抱歉\n目前只有饅頭人、熊大、兔兔可以更換喔！\n\n請輸入:\n大頭貼 熊大\n來使用此功能。"
	GetInfoError = "無法取得影片資訊\n請重新確認 ID。"
	NoDownload = "很抱歉, 該影片不提供下載"
)
// const Unknown = "未知錯誤 請稍後再嘗試。"

func GetTestMessage(ReplyToken string, MessageID string, MessageText string, UserID string, QuotaValue int64) string {
	return "EventReplyToken: "+ReplyToken+"\nMessageID: "+MessageID+"\nMessageText: "+MessageText+"\nUserID: "+UserID+"\nremain message: "+strconv.FormatInt(QuotaValue, 10)
}

func GetStickerMessage(name string) string {
	return "嗨嗨 我是"+name+"， 歡迎使用蟲蟲機器人。"
}

func GetInfo(title string, author string, url string) string {
	return "影片標題: "+title+"\n頻道: "+author+"\n下載連結: "+url
}