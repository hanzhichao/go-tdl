package go_tdl

const (
	Passed = iota
	Failed
	Error
	Skipped
	XFailed
	XPassed
)

type TestSummary struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Duration  string `json:"duration"`
	IsSuccess bool   `json:"is_success"`

	Total    int `json:"total"`
	TotalRun int `json:"total_run"`
	Passed   int `json:"passed"`
	Failed   int `json:"failed"`
	Error    int `json:"error"`
	Skipped  int `json:"skipped"`
	XFailed  int `json:"xfailed"`
	XPassed  int `json:"xpassed"`
}

type TestCaseResult struct {
	TestCaseId          TestCaseId `json:"testcase_id"`
	TestCaseName        string     `json:"testcase_name"`
	TestCaseDescription string     `json:"testcase_description"`

	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Duration  string `json:"duration"`

	Status    int    `json:"status"`
	Output    string `json:"output"`
	ErrorInfo string `json:"error_info"`
}

type TestResult struct {
	Summary TestSummary      `json:"summary"`
	Details []TestCaseResult `json:"details"`
}

func (result *TestResult) Rerun() {
	// TODO
}
