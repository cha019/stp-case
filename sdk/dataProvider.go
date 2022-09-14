package sdk

import "math/rand"

type DataProviderIface interface {
	GetRandomData(num int) int64
}

type DataProvider struct {
}

func (d *DataProvider) GetRandomData(num int) int64 {
	return int64(rand.Intn(num))
}
