package mediator

import "testing"

func TestMediator(t *testing.T){
	m := NewConcreteMediator()

	ca := NewConcreteColleageA(m)
	cb := NewConcreteColleagueB(m)

	m.AddColleague(ca)
	m.AddColleague(cb)

	ca.Send("ca")
	cb.Send("cb")
}