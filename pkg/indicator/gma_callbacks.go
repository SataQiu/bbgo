// Code generated by "callbackgen -type GMA"; DO NOT EDIT.

package indicator

import ()

func (inc *GMA) OnUpdate(cb func(value float64)) {
	inc.UpdateCallbacks = append(inc.UpdateCallbacks, cb)
}

func (inc *GMA) EmitUpdate(value float64) {
	for _, cb := range inc.UpdateCallbacks {
		cb(value)
	}
}