package sdk

var (
	cases = make(map[string]CaseFn)
)

type CaseConfig struct {
	CID       string
	ENV       string
	CaseName  string
	ExtraInfo string
	Verbose   bool
}

func NewCaseConfig(cid string, env string, caseName string, extraInfo string, verbose bool) *CaseConfig {
	c := &CaseConfig{
		CID:       cid,
		ENV:       env,
		CaseName:  caseName,
		ExtraInfo: extraInfo,
		Verbose:   verbose,
	}
	return c
}

type CaseFn = func(config *CaseConfig) (*CaseResponse, error)

func RegisterCaseFn(caseName string, f CaseFn) {
	cases[caseName] = f
}

func GetCaseFn(caseName string) CaseFn {
	return cases[caseName]
}

func GetAllCaseName() []string {
	var caseNames []string
	for k, _ := range cases {
		caseNames = append(caseNames, k)
	}
	return caseNames
}

type CaseResponse struct {
	URL  string
	Code uint32
	Body []byte
	Cost int64
}
