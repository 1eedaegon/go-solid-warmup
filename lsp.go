package main

// Case 3 - Liskov subsitution
type Duck interface {
	Kwack()
}

type S struct {}
func (s *S) Kwack() {}

type U struct {}
func (u *U) Kwack() {}

func MarshalDuck(t T) {}

var s = &S{}
var u = &U{}


// 오리로써 작동하게 끔 만들어야한다.
// Golang이 상속을 지원하지 않기에 더욱 강화된다.
MarshalDuck(u)
MarshalDuck(s)
