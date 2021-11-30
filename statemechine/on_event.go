package statemechine

import (
	"log"
)

type OnEvent interface {
	OnEnter(toState string, args []interface{})
	OnAction(action string, fromState string, toState string, args []interface{}) error
	OnActionFailure(action string, fromState string, toState string, args []interface{}, err error)
	OnExit(fromState string, args []interface{})
}

//自定义Event
type DefaultOnEvent struct{}

func (p *DefaultOnEvent) OnEnter(toState string, args []interface{}) {
	log.Printf("OnEnter.... %+v %+v", toState, args)
}

func (p *DefaultOnEvent) OnAction(action string, fromState string, toState string, args []interface{}) error {
	log.Printf("OnAction.... %+v %+v", action, args)
	return nil
}

func (p *DefaultOnEvent) OnActionFailure(action string, fromState string, toState string, args []interface{}, err error) {
	log.Printf("OnActionFailure.... %+v %+v", action, err)
}

func (p *DefaultOnEvent) OnExit(fromState string, args []interface{}) {
	log.Printf("onExit.... %+v %+v", fromState, args)
}
