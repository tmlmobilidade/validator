package lib

type AssertionMessage struct {
	Expected int
	Actual int
	Message string
}

func Assert(assertion AssertionMessage) (string) {
	if assertion.Expected != assertion.Actual {
		return assertion.Message
	}
	
	return ""
}