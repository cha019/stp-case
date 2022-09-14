package sdk

import (
	"github.com/myzhan/boomer"
	"log"
	"net/http"
)

type SDKConfig struct {
	testData         [][]string
	chanData         chan string
	dataProvider     DataProviderIface
	globalHttpClient *http.Client
	globalBoomer     *boomer.Boomer
	isDebug          bool
	logger           *log.Logger
}

var GlobalConfig = &SDKConfig{}

func SetTestData(dataPath string) error {
	var err error
	GlobalConfig.testData, err = ReadFileToSlice(dataPath)
	if err != nil {
		return err
	}
	return nil
}

func SetTestDataByData(data [][]string) {
	GlobalConfig.testData = data
}

func GetTestData() [][]string {
	return GlobalConfig.testData
}

func SetChanData(dataPath string) error {
	var err error
	GlobalConfig.chanData, err = ReadFileToChan(dataPath)
	if err != nil {
		return err
	}
	return nil
}

func GetChanData() chan string {
	return GlobalConfig.chanData
}

func SetGlobalHttpClient(c *http.Client) {
	GlobalConfig.globalHttpClient = c
}

func GetGlobalHttpClient() *http.Client {
	if GlobalConfig.globalHttpClient == nil {
		GlobalConfig.globalHttpClient = GetHttpClient()
		return GlobalConfig.globalHttpClient
	}
	return GlobalConfig.globalHttpClient
}

func SetGlobalBoomer(b *boomer.Boomer) {
	GlobalConfig.globalBoomer = b
}

func GetGlobalBoomer() *boomer.Boomer {
	return GlobalConfig.globalBoomer
}

func IsExistGlobalBoomer() bool {
	if GlobalConfig.globalBoomer != nil {
		return true
	}
	return false
}

func SetDataProvider(d DataProviderIface) {
	GlobalConfig.dataProvider = d
}

func GetDataProvider() DataProviderIface {
	return GlobalConfig.dataProvider
}

func SetDebug(d bool) {
	GlobalConfig.isDebug = d
}

func IsDebug() bool {
	return GlobalConfig.isDebug
}

func SetSDKLogger(log *log.Logger) {
	GlobalConfig.logger = log
}

func GetSDKLogger() *log.Logger {
	if GlobalConfig.logger == nil {
		GlobalConfig.logger = log.Default()
		return GlobalConfig.logger
	}
	return GlobalConfig.logger
}
