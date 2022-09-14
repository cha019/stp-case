package mode

import (
	SDK "git.garena.com/chao.huang/stp-case/sdk"
	"github.com/myzhan/boomer"
	"log"
	"strconv"
	"sync"
)

var (
	caseConfigMap sync.Map
	taskList      []*boomer.Task
)

func Mix(cid string, env string, maxCaseNameMap map[string]string, extraInfo string, verbose bool) {

	log.Println("max api stress test")

	for key, value := range maxCaseNameMap {
		stressCase := SDK.NewCaseConfig(cid, env, key, extraInfo, verbose)
		scale, _ := strconv.Atoi(value)
		caseConfigMap.Store(*stressCase, scale)
	}

	log.Println(caseConfigMap)

	caseConfigMap.Range(setTask)

	log.Println(taskList)

	if SDK.IsExistGlobalBoomer() {
		SDK.GetGlobalBoomer().Run(taskList...)
	} else {
		boomer.Run(taskList...)
	}

}

func setTask(key, value interface{}) bool {

	caseConfig := key.(SDK.CaseConfig)
	weight := value.(int)

	fn := SDK.GetCaseFn(caseConfig.CaseName)
	if fn == nil {
		log.Println("case not found")
		panic("case not found")
		return false
	}

	task := &boomer.Task{
		Name:   "worker",
		Weight: weight,
		Fn: func() {
			_, _ = fn(&caseConfig)
		},
	}

	taskList = append(taskList, task)

	return true
}
