package main

import (
	"flag"
	"fmt"
	_ "github.com/h17600445140/stp-case/case/api"
	_ "github.com/h17600445140/stp-case/case/chanRead"
	MODE "github.com/h17600445140/stp-case/mode"
	SDK "github.com/h17600445140/stp-case/sdk"
	"github.com/myzhan/boomer"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	caseName    string
	testData    string
	dataMode    string
	cid         string
	env         string
	verbose     bool
	mode        string
	isLocal     string
	maxRps      int
	concurrency int
	spawnRate   int
	extraInfo   string
)

func init() {

	rand.Seed(time.Now().UnixNano())
	SDK.SetGlobalHttpClient(SDK.GetHttpClient())

}

func main() {

	mixCaseNameMap := make(mapFlag)
	logger := log.Default()

	// cmd
	// local:
	// go run ./main.go -mode single -case_name BatchGetData -input "./data/test.csv" -cid cn -env test -local true -max_rps 100 -c 20 -r 2
	// go run ./main.go -mode single -case_name BatchGetData -input "./data/test.csv" -cid cn -env test -local true -max_rps 100 -c 20 -r 2
	// go run ./main.go -mode mix -input "./data/test.csv" -cid sg -env test -local true -max_rps 5 -c 2 -r 1 -mixCaseName BatchGetData=2 -mixCaseName test=1

	// remote:
	// locust --master --expect-slave=1 -f dummy.py --no-web -c 5 -r 1 -t 10s
	// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main.out
	// ./main.out -mode single -case_name BatchGetData -input "./data/test.csv" -cid cn -env test -master-host xx.xx.xx.xx
	// ./main.out -mode mix -input "./data/test.csv" -cid cn -env test -master-host xx.xx.xx.xx -mixCaseName BatchGetData=2 -mixCaseName test=1

	// chan
	// go run ./main.go -mode single -case_name test -input "./data/test.csv" -dataMode chanRead -cid cn -env test -local true -max_rps 100 -c 20 -r 2

	flag.StringVar(&caseName, "case_name", "demo", "test case which was registered")
	flag.StringVar(&testData, "input", "./data/test.csv", "can be a file path, json etc.")
	flag.StringVar(&dataMode, "dataMode", "randomRead", "1. randomRead 2. chanRead")
	flag.StringVar(&cid, "cid", "xx", "cid")
	flag.StringVar(&env, "env", "test", "can't be empty")
	flag.StringVar(&mode, "mode", "single", "stress test mode")
	flag.Var(&mixCaseNameMap, "mixCaseName", "example: -mixCaseName caseName1=3 -mixCaseName caseName2=3 -mixCaseName caseName3=4")
	flag.StringVar(&isLocal, "local", "false", "is local")
	flag.IntVar(&maxRps, "max_rps", 1, "max qps created by stress test")
	flag.IntVar(&concurrency, "c", 1, "concurrent requests")
	flag.IntVar(&spawnRate, "r", 1, "concurrency requests created per second")
	flag.StringVar(&extraInfo, "extraInfo", "", "case extraInfo")
	flag.BoolVar(&verbose, "verbose", false, "Print debug log")
	flag.Parse()

	logger.Printf(`running with these args : caseName = %s ,testData = %s , cid = %s , env = %s , mode = %s , 
		mixCaseName = %v , isLocal = %v , maxRps = %v , concurrency = %v , spawnRate = %v , extraInfo = %v, verbose = %t`,
		caseName,
		testData,
		cid,
		env,
		mode,
		mixCaseNameMap,
		isLocal,
		maxRps,
		concurrency,
		spawnRate,
		extraInfo,
		verbose)

	if dataMode == "randomRead" {
		err := SDK.SetTestData(testData)
		if err != nil {
			logger.Println(err)
		}
	} else if dataMode == "chanRead" {
		err := SDK.SetChanData(testData)
		if err != nil {
			logger.Println(err)
		}
	}

	SDK.SetGlobalHttpClient(SDK.GetHttpClient())
	SDK.SetDataProvider(&SDK.DataProvider{})
	SDK.SetSDKLogger(logger)

	// boomer local execute
	if isLocal != "false" {
		b := boomer.NewStandaloneBoomer(concurrency, float64(spawnRate))
		rateLimiter := boomer.NewStableRateLimiter(int64(maxRps), time.Second)
		b.SetRateLimiter(rateLimiter)
		SDK.SetGlobalBoomer(b)
	}

	if SDK.IsExistGlobalBoomer() {
		logger.Println("Global Boomer exit")
	} else {
		logger.Println("Global Boomer is not exit")
	}

	SDK.SetDebug(true)

	if mode == "single" {

		MODE.Single(cid, env, verbose, caseName, extraInfo)

	}

	if mode == "mix" {

		MODE.Mix(cid, env, mixCaseNameMap, extraInfo, verbose)

	}

}

type mapFlag map[string]string

func (f mapFlag) String() string {
	return fmt.Sprintf("%v", map[string]string(f))
}

func (f mapFlag) Set(value string) error {
	split := strings.SplitN(value, "=", 2)
	f[split[0]] = split[1]
	return nil
}
