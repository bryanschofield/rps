package echo

func NewEchoer() Echoer {
	return &simpleEchoer{}
}

type simpleEchoer struct{}

func (_ *simpleEchoer) Echo(in string) string {
	return in
}
