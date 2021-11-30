package statemechine

import (
	"encoding/json"
	"errors"
)

type StateMachine struct {
	onEvent    OnEvent
	transforms []Transform
	state      string
}

//AddTransforms 注册状态表
func (sm *StateMachine) AddTransforms(transList []Transform) {
	sm.transforms = transList
}

//AddTransformsFromJson 注册状态表
func (sm *StateMachine) AddTransformsFromJson(jsonStr string) error {
	transList := make([]Transform, 0)
	err := json.Unmarshal([]byte(jsonStr), &transList)
	if err != nil {
		return err
	}

	sm.AddTransforms(transList)
	return nil
}

//AddOnEvent 注册事件处理器
func (sm *StateMachine) AddOnEvent(event OnEvent) error {
	sm.onEvent = event
	return nil
}

//GetCurState 获取当前状态
func (sm *StateMachine) GetCurState() string {
	return sm.state
}

//SetCurState 设置当前状态
func (sm *StateMachine) SetCurState(s string) {
	sm.state = s
}

//Emit 触发事件
func (sm *StateMachine) Emit(event string, args ...interface{}) error {
	curState := sm.GetCurState()
	for _, v := range sm.transforms {
		if v.From == curState && v.Event == event {
			err := sm.handleEvent(v.Action, v.From, v.To, args)
			return err
		}
	}

	return errors.New("no target event")
}

func (sm *StateMachine) handleEvent(action string, fromState string, toState string, args []interface{}) error {
	if fromState != toState {
		sm.onEvent.OnExit(fromState, args)
	}

	err := sm.onEvent.OnAction(action, fromState, toState, args)
	if err != nil {
		sm.onEvent.OnActionFailure(action, fromState, toState, args, err)
		return err
	}

	sm.state = toState

	if fromState != toState {
		sm.onEvent.OnEnter(toState, args)
	}

	return nil
}
