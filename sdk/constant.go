package sdk

import "strings"

func GetUrlByEnvAndCid(env string, cid string) string {

	if env == "live" {
		switch strings.ToLower(cid) {
		case "cn":
			return "https://www.test.baidu.com"
		default:
			return ""
		}
	} else if env == "test" {
		switch strings.ToLower(cid) {
		case "cn":
			return "https://www.test.baidu.com"
		default:
			return ""
		}
	} else {
		return ""
	}

}
