package sdk

import "strings"

var (
	clientIdTest = "28de986e-d334-11eb-8553-acde48001122"
	clientId     = "7203e231-521b-4ba2-b8c0-f40e7f80077c"
)

func GetProxyUrlByEnvAndCid(env string, cid string) string {

	if env == "live" {
		switch strings.ToLower(cid) {
		case "sg":
			return "https://serviceproxy.airpay.sg/grpc"
		case "my":
			return "https://serviceproxy.airpay.com.my/grpc"
		case "id":
			return "https://serviceproxy.airpay.co.id/grpc"
		case "ph":
			return "https://serviceproxy.airpay.com.ph/grpc"
		default:
			return ""
		}
	} else if env == "test" {
		switch strings.ToLower(cid) {
		case "sg":
			return "https://serviceproxy.test.airpay.sg/grpc"
		case "my":
			return "https://serviceproxy.test.airpay.com.my/grpc"
		case "id":
			return "https://serviceproxy.test.airpay.co.id/grpc"
		case "ph":
			return "https://serviceproxy.test.airpay.com.ph/grpc"
		default:
			return ""
		}
	} else {
		return ""
	}

}

func GetGwUrlByEnvAndCid(env string, cid string) string {

	if env == "live" {
		switch strings.ToLower(cid) {
		case "sg":
			return ""
		default:
			return ""
		}
	} else if env == "test" {
		switch strings.ToLower(cid) {
		case "sg":
			return "https://api.gw.test.airpay.sg"
		case "my":
			return "https://api.gw.test.airpay.com.my"
		case "id":
			return "https://api.gw.test.airpay.co.id"
		case "ph":
			return "https://api.gw.test.airpay.com.ph"
		default:
			return ""
		}
	} else {
		return ""
	}

}

func GetClientId(env string, cid string) string {

	if env == "live" {
		switch strings.ToLower(cid) {
		case "id":
			return clientId
		case "sg":
			return clientId
		case "ph":
			return clientId
		case "my":
			return clientId
		default:
			return ""
		}
	} else if env == "test" {
		switch strings.ToLower(cid) {
		case "id":
			return clientIdTest
		case "sg":
			return clientIdTest
		case "ph":
			return clientIdTest
		case "my":
			return clientIdTest
		default:
			return ""
		}
	} else {
		return ""
	}

}
