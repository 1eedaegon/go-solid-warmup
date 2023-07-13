package main

import "time"

// Case 4 - Interface segregation

type Report interface {
	Report() string
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report){
	send(r.Report())
}
// 나는 Report만 필요한데 메타 정보까지 한꺼번에 모두 가져와야한다.

type Report interface{
	Report()
}

type WrittenInfo interface {
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) {
	send(r.Report())
}