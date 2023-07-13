package main

// Case 2 - Open closed principle
func SendReport(r *Report, method SendType, reciever string) {
	switch method {
	case Email:
	case Fax:
	case Printer:
	default:
	}
}
// 새로운 전송방식을 추가하기 위해 SendReport함수를 수정해야한다.
// 함수에 의존적인 코드 전부 영향

// Duck typing
type ReportSender interface {
	Send(r *Report)
}

type EmailSender struct {}
func (e *EmailSender) Send(r *Report) {}

type FaxSender struct {}
func (f *FaxSender) Send(r *Report) {}

