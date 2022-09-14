package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/myzhan/boomer"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	responseBody CommonResponse
)

func CallServiceProxyHttp(r *HTTPRequest) (*CaseResponse, error) {

	var (
		url          = fmt.Sprintf("%s://%s%s", r.Request.URL.Scheme, r.Request.URL.Host, r.Request.URL.Path)
		caseResponse = &CaseResponse{URL: url}
		httpResponse *http.Response
	)

	httpClient := GetGlobalHttpClient()

	if r.Verbose {
		if r.Request.Body != nil {
			body, err := ioutil.ReadAll(r.Request.Body)
			if err != nil {
				GetSDKLogger().Printf("ioutil.ReadAll err is %v\n", err)
				return caseResponse, err
			}
			r.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			GetSDKLogger().Printf("request info : url is %s , header is %s , requestBody is %s\n",
				url, r.Request.Header, string(body))
		} else {
			GetSDKLogger().Printf("request body is nil\n")
		}
	}

	startTime := time.Now()
	httpResponse, err := httpClient.Do(r.Request)
	elapsed := time.Since(startTime)
	caseResponse.Cost = elapsed.Microseconds()

	if err != nil {
		if IsDebug() {
			if IsExistGlobalBoomer() {
				GetGlobalBoomer().RecordFailure(r.RequestName, "send req error", 0.0, err.Error())
			} else {
				boomer.RecordFailure(r.RequestName, "send req error", 0.0, err.Error())
			}
		}
		GetSDKLogger().Printf("httpClient.Do err is %v\n", err)
		return caseResponse, err
	}

	if httpResponse == nil {
		if IsDebug() {
			if IsExistGlobalBoomer() {
				GetGlobalBoomer().RecordFailure(r.RequestName, "resp is nil", 0.0, "resp is nil")
			} else {
				boomer.RecordFailure(r.RequestName, "resp is nil", 0.0, "resp is nil")
			}
		}
		GetSDKLogger().Printf("httpResponse == nil \n")
		return caseResponse, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			GetSDKLogger().Printf("io.ReadCloser err is %v\n", err)
		}
	}(httpResponse.Body)

	caseResponse.Code = uint32(httpResponse.StatusCode)

	if httpResponse.StatusCode == 200 {
		result, _ := ioutil.ReadAll(httpResponse.Body)
		caseResponse.Body = result
		err := json.Unmarshal(result, &responseBody)
		if err != nil {
			GetSDKLogger().Printf("json.Unmarshal err is %v\n", err)
			return caseResponse, err
		}
		if responseBody.Code == 0 {
			if IsDebug() {
				if IsExistGlobalBoomer() {
					GetGlobalBoomer().RecordSuccess(r.RequestName, "success",
						elapsed.Nanoseconds()/int64(time.Millisecond), httpResponse.ContentLength)
				} else {
					boomer.RecordSuccess(r.RequestName, "success",
						elapsed.Nanoseconds()/int64(time.Millisecond), httpResponse.ContentLength)
				}
			}
		} else {
			if IsDebug() {
				if IsExistGlobalBoomer() {
					GetGlobalBoomer().RecordFailure(r.RequestName, strconv.Itoa(responseBody.Code), elapsed.Nanoseconds()/int64(time.Millisecond), responseBody.Msg)
				} else {
					boomer.RecordFailure(r.RequestName, strconv.Itoa(responseBody.Code), elapsed.Nanoseconds()/int64(time.Millisecond), responseBody.Msg)
				}
			}
			GetSDKLogger().Printf("responseBody.Code == %v\n", responseBody.Code)
		}
	} else {
		if IsDebug() {
			if IsExistGlobalBoomer() {
				GetGlobalBoomer().RecordFailure(r.RequestName, "error", 0.0, strconv.Itoa(httpResponse.StatusCode))
			} else {
				boomer.RecordFailure(r.RequestName, "error", 0.0, strconv.Itoa(httpResponse.StatusCode))
			}
		}
		GetSDKLogger().Printf("httpResponse.StatusCode == %v\n", httpResponse.StatusCode)
	}

	return caseResponse, nil

}
