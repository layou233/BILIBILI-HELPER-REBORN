package utils

import (
	"fmt"
	"github.com/iyear/biligo"
	"time"
)

type vip struct {
	Type   string
	Due    time.Time
	Status string
}

func (v vip) String() string {
	if v.Type == "" {
		return "未开通, 过期时间: " + v.Due.String()
	}
	return fmt.Sprintf("%s, 到期时间: %s, 状态:%s", v.Type, v.Due.String(), v.Status)
}

func GetVIPStat(stat *biligo.VipStat) (v vip) {
	v = vip{
		Due: GetCST8Time(time.UnixMilli(stat.VipDueDate)),
	}
	if stat.VipDueDate < time.Now().UnixMilli() {
		v.Status = "已过期"
		return
	}
	switch stat.VipType {
	case 0: // 未开通
		return
	case 1:
		v.Type = "月度大会员"
		break
	case 2:
		v.Type = "年度大会员"
		break
	default:
		v.Type = "出错了"
	}
	switch stat.VipStatus {
	case 0:
		v.Status = "正常"
		break
	case 1:
		v.Status = "由于IP地址更换过于频繁，服务被冻结."
		break
	case 2:
		v.Status = "你的大会员账号风险过高，大会员功能已被锁定"
		break
	}
	return
}
