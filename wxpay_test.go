package wxpay

import (
	"testing"
	"fmt"
)

func TestWxPay_Do(t *testing.T) {
	o := &WxPayUnifiedOrder{}
	o.Body = "wawaji积分"
	o.OutTradeNo = "2017071216411234567"
	o.TotalFee = 15
	o.SpbillCreateIp = "192.168.1.109"
	o.NotifyUrl = "https://www.baidu.com/"
	o.TradeType = APP
	cli := WxPayClient("wx", "138", "7c01cd5266b24989b")

	resp, _ := cli.DoRequest(o)
	var data WxPayUnifiedOrderResp
	cli.ReadResponse(resp, &data)

	fmt.Println(data)
}
