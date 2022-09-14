package chanRead

import (
	"fmt"
	SDK "github.com/h17600445140/stp-case/sdk"
)

const (
	CaseNameTest = "test"
)

func init() {
	SDK.RegisterCaseFn(CaseNameTest, Test)
}

func Test(config *SDK.CaseConfig) (*SDK.CaseResponse, error) {

	fmt.Println(config)

	var data string

	// 读完停止
	if v, err := SDK.GetDataFromChan(SDK.GetChanData()); err != nil {
		return nil, err
	} else {
		data = v
	}

	// 读完循环
	//if v, err := SDK.GetDataFromChan(SDK.GetChanData()); err != nil {
	//	fmt.Printf("appear err, err is %v", err)
	//} else {
	//	data = v
	//	SDK.GetChanData() <- data
	//}

	fmt.Println(data)

	res := &SDK.CaseResponse{}

	return res, nil

}
