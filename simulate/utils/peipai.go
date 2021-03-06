package utils

import (
	"fmt"
	"net/http"
	"steve/client_pb/room"
	"steve/simulate/config"

	"github.com/Sirupsen/logrus"
)

// Peipai 配牌
// step 1. 将手牌和墙牌转换成接口可识别的字符串， 参数： cards = 11,12,13,,...  len=xxx
// step 2. 将换三张方向转换成接口可识别的字符串 ， hszfx=dui, shun, ni
// step 3. 庄家位置   zhuang= number
func Peipai(game string, seatCards [][]uint32, wallCards []uint32, hszDir room.Direction, bankerSeat int) error {

	// url 基本字符串 + 牌数据
	url := fmt.Sprintf("%s?game=%s&%s", config.GetPeipaiURL(), game, translatePeipaiCards(seatCards, wallCards))

	// 换三张方向
	hszfx := translateHszDir(hszDir)
	if hszfx != "" {
		url = fmt.Sprintf("%s&%s", url, hszfx)
	}

	// 庄家座位号
	url = fmt.Sprintf("%s&zhuang=%d", url, bankerSeat)

	fmt.Println("utils.peipai() url = ", url)

	// 发出url请求
	return requestPeipai(url)
}

func requestPeipai(url string) error {
	logrus.WithField("url", url).Info("请求配牌")
	_, err := http.DefaultClient.Get(url)
	return err
}

// ClearPeiPai 配牌
func ClearPeiPai(game string) error {
	//不传cards视为清除
	url := fmt.Sprintf("%s?game=%s", config.GetPeipaiURL(), game)
	fmt.Println("utils.ClearPeiPai() url = ", url)

	// 发出url请求
	return requestPeipai(url)
}

// translatePeipaiCards 将卡牌转换成配牌接口字符串
// 返回：  cards=11,22,33,44,...&len=22
func translatePeipaiCards(seatCards [][]uint32, wallCards []uint32) string {
	result := "cards="
	first := true
	count := 0
	for _, cards := range seatCards {
		for _, card := range cards {
			if first {
				result = fmt.Sprintf("%s%s", result, translatePeipaiCard(card))
				first = false
			} else {
				result = fmt.Sprintf("%s,%s", result, translatePeipaiCard(card))
			}
			count++
		}
	}
	for _, wallCard := range wallCards {
		result = fmt.Sprintf("%s,%s", result, translatePeipaiCard(wallCard))
		count++
	}
	result = fmt.Sprintf("%s&len=%d", result, count)
	return result
}

// translatePeipaiCard 转换单张牌
func translatePeipaiCard(card uint32) string {
	return fmt.Sprint(card)
}

// 格式化换三张的字符串
func translateHszDir(dir room.Direction) string {
	switch dir {
	case room.Direction_AntiClockWise:
		{
			return "hszfx=ni"
		}
	case room.Direction_ClockWise:
		{
			return "hszfx=shun"
		}
	case room.Direction_Opposite:
		{
			return "hszfx=dui"
		}
	}
	return ""
}
