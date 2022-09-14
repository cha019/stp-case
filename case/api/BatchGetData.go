package api

import (
	"bytes"
	"encoding/json"
	SDK "github.com/h17600445140/stp-case/sdk"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	CaseNameBatchGetData = "BatchGetData"
)

func init() {
	SDK.RegisterCaseFn(CaseNameBatchGetData, BatchGetData)
}

func BatchGetData(config *SDK.CaseConfig) (*SDK.CaseResponse, error) {

	dataLists := SDK.GetTestData()

	requestUrl := SDK.GetUrlByEnvAndCid(config.ENV, config.CID)

	req := RequestUidList{}
	for i := 0; i < 20; i++ {
		id, _ := strconv.Atoi(dataLists[rand.Intn(len(dataLists))][0])
		req.UidList = append(req.UidList, id)
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))

	request.Header.Set("X-Service", "xxx-xxx")

	newRequest := &SDK.HTTPRequest{
		Request:     request,
		Verbose:     config.Verbose,
		RequestName: config.CaseName,
	}

	return SDK.CallServiceProxyHttp(newRequest)

}
