package cardtype

import (
	"steve/majong/interfaces"
	"steve/majong/settle/fan"
	majongpb "steve/server_pb/majong"
)

type cardTypeCalculator struct {
}

//CardTypeValueMap  卡牌类型与倍数映射
var CardTypeValueMap = make(map[interfaces.CardType]int)

// 初始番型与卡牌映射
func init() {
	// 平胡 1
	CardTypeValueMap[majongpb.CardType_PingHu] = 1
	// 清一色 4
	CardTypeValueMap[majongpb.CardType_QingYiSe] = 4
	// 七对 4
	CardTypeValueMap[majongpb.CardType_QiDui] = 4
	// 清七对 16
	CardTypeValueMap[majongpb.CardType_QingQiDui] = 16
	// 龙七对 8
	CardTypeValueMap[majongpb.CardType_LongQiDui] = 8
	// 清龙七对 32
	CardTypeValueMap[majongpb.CardType_QingLongQiDui] = 32
	// 碰碰胡 2
	CardTypeValueMap[majongpb.CardType_PengPengHu] = 2
	// 清碰 8
	CardTypeValueMap[majongpb.CardType_QingPeng] = 8
	// 金钩钓 4
	CardTypeValueMap[majongpb.CardType_JingGouDiao] = 4
	// 清金钩钓 16
	CardTypeValueMap[majongpb.CardType_QingJingGouDiao] = 16
	// 十八罗汉 64
	CardTypeValueMap[majongpb.CardType_ShiBaLuoHan] = 64
	// 清十八罗汉 256
	CardTypeValueMap[majongpb.CardType_QingShiBaLuoHan] = 254

}

//Calculate 获取能胡所有番型，及根
func (ctc *cardTypeCalculator) Calculate(params interfaces.CardCalcParams) (cardTypes []interfaces.CardType, gengCount int) {
	fanCardTypes := make([]interfaces.CardType, 0)
	// 遍历可行番型
	for i := 0; i < len(fan.AllFan); i++ {
		if fan.AllFan[i].Condition() {
			fanName := fan.AllFan[i].GetFanName()
			fanCardTypes = append(fanCardTypes, fanName)
		}
	}
	// 获取根数
	// gengCount = fan.GetGenCount(winner)
	// 番型名和根处理
	// scxlFanMutex
	cardTypes = fanCardTypes
	return cardTypes, gengCount
}

//CardTypeValue 获取总倍数
func (ctc *cardTypeCalculator) CardTypeValue(cardTypes []interfaces.CardType, gengCount int) int {
	total := 1
	// 叠乘番型
	for _, cardType := range cardTypes {
		if multiple, isExist := CardTypeValueMap[cardType]; isExist {
			total = total * multiple
		}
	}
	// 根平方
	genTotoal := 1 << uint(gengCount)
	// 根乘总番型倍数
	total = total * genTotoal
	return total
}
