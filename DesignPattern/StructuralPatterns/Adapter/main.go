package main

//适配器模式（Adapter):将一个类的接口转换成客户希望的另外一个接口
//使得原本由于接口不兼容而不能一起工作的那些类能一起工作
func main() {
	sms := NewAdapterSMS("你好北京", "北京土著")
	sms.SendSms()
}
