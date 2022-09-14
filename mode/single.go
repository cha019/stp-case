package mode

import (
	SDK "github.com/h17600445140/stp-case/sdk"
	"github.com/myzhan/boomer"
	"log"
)

func Single(cid string, env string, verbose bool, caseName string, extraInfo string) {

	log.Println("single api stress test")

	caseConfig := &SDK.CaseConfig{
		CID:       cid,
		ENV:       env,
		Verbose:   verbose,
		CaseName:  caseName,
		ExtraInfo: extraInfo,
	}

	fn := SDK.GetCaseFn(caseName)
	if fn == nil {
		log.Println("case not found")
		panic("case not found")
	}

	task := &boomer.Task{
		Name:   "worker",
		Weight: 10,
		Fn: func() {
			_, _ = fn(caseConfig)
		},
	}

	if SDK.IsExistGlobalBoomer() {
		SDK.GetGlobalBoomer().Run(task)
	} else {
		boomer.Run(task)
	}

}
