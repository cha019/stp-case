package shopeepayUserAccount

import (
	"bytes"
	"encoding/json"
	SDK "git.garena.com/chao.huang/stp-case/sdk"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	CaseNameBatchGetSpUidByUid = "BatchGetSpUidByUid"
)

func init() {
	SDK.RegisterCaseFn(CaseNameBatchGetSpUidByUid, BatchGetSpUidByUid)
}

func BatchGetSpUidByUid(config *SDK.CaseConfig) (*SDK.CaseResponse, error) {

	dataLists := SDK.GetTestData()

	requestUrl := SDK.GetProxyUrlByEnvAndCid(config.ENV, config.CID)

	req := RequestUidList{}
	for i := 0; i < 20; i++ {
		shopeeId, _ := strconv.Atoi(dataLists[rand.Intn(len(dataLists))][0])
		req.UidList = append(req.UidList, shopeeId)
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))

	request.Header.Set("X-Client-Id", SDK.GetClientId(config.ENV, config.CID))
	request.Header.Set("X-Method", "/user.AccountSApi/BatchGetSpUidByUid")
	request.Header.Set("X-Service", "shopeepayuser-account")
	request.Header.Set("X-Tag", "account_limiter")
	//request.Header.Set("X-Tag", "rate_limiter")

	newRequest := &SDK.HTTPRequest{
		Request:     request,
		Verbose:     config.Verbose,
		RequestName: config.CaseName,
	}

	return SDK.CallServiceProxyHttp(newRequest)

}
