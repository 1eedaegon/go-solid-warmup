package main

// Case 5 - Dependency inversion

// Embedding(Composition)
type Mail struct {
	alarm Alarm
}
type Alarm struct {}

type (m *Mail) OnRecv() {
	m.alarm.Alarm()
}

// 상속보다 낫지만 Mail은 Alarm에 의존적이다.
// 모듈간 결합도 최상

// Mail이라는 event를 구현하고 Mail event가 발생했을 때
// Handle할 수 있도록 Alarm이라는 listener를 구현하겠다.

type Event interface {
	Register(EventListener) // Event에 Listener를 입력 시 등록한다.
}

type EventListener interface {
	OnFire()
}

type Msg struct {
	id string
	content string
}
type ProduceQueue interface {
	Produce(msg Msg)
}

type Mail struct {
	listner EventListener
}

func (m *Mail) Register(listner EventListener) {
	m.listner = listner
}

func (m *Mail) OnRecv() {
	m.listner.OnFire()
}

type KafkaQMsg struct {
	path string
	msg Msg
}
func (k *KafkaQMsg) Produce(msg Msg) {
	k.msg = msg
	k.Send(msg)
}

type Alarm struct {}
// OnFire를 구현한 순간 Eventlistner가 된다.
func (a *Alarm) OnFire() {
	k:= &KafkaQMsg{...}
	k.Produce()
} 

var main = &Mail{}
var listner EventListner = &Alarm{}

mail.Register(listner)
mail.OnRecv()
