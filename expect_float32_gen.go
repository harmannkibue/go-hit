// +build !generate_numeric

// ⚠️⚠️⚠️ This file was autogenerated by generators/expect/numeric ⚠️⚠️⚠️ //
package hit

import minitest "github.com/Eun/go-hit/internal/minitest"

// IExpectFloat32 provides assertions for the float32 type.
type IExpectFloat32 interface {
	// Equal expects the float32 to be equal to the specified value.
	Equal(value float32) IStep

	// NotEqual expects the float32 to be not equal to the specified value.
	NotEqual(value ...float32) IStep

	// OneOf expects the float32 to be equal to one of the specified values.
	OneOf(values ...float32) IStep

	// NotOneOf expects the float32 to be not equal to one of the specified values.
	NotOneOf(values ...float32) IStep

	// GreaterThan expects the float32 to be not greater than the specified value.
	GreaterThan(value float32) IStep

	// LessThan expects the float32 to be less than the specified value.
	LessThan(value float32) IStep

	// GreaterOrEqualThan expects the float32 to be greater or equal than the specified value.
	GreaterOrEqualThan(value float32) IStep

	// LessOrEqualThan expects the float32 to be less or equal than the specified value.
	LessOrEqualThan(value float32) IStep

	// Between expects the float32 to be between the specified min and max value (inclusive, min >= float32 >= max).
	Between(min, max float32) IStep

	// NotBetween expects the float32 to be not between the specified min and max value (inclusive, min >= float32 >= max).
	NotBetween(min, max float32) IStep
}
type expectFloat32ValueCallback func(hit Hit) float32
type expectFloat32 struct {
	cp            callPath
	valueCallback expectFloat32ValueCallback
}

func newExpectFloat32(cp callPath, valueCallback expectFloat32ValueCallback) IExpectFloat32 {
	return &expectFloat32{cp: cp, valueCallback: valueCallback}
}
func (v *expectFloat32) Equal(value float32) IStep {
	return &hitStep{Trace: ett.Prepare(), When: ExpectStep, CallPath: v.cp.Push("Equal", []interface{}{value}), Exec: func(hit *hitImpl) error {
		return minitest.Error.Equal(v.valueCallback(hit), value)
	}}
}
func (v *expectFloat32) NotEqual(values ...float32) IStep {
	return &hitStep{Trace: ett.Prepare(), When: ExpectStep, CallPath: v.cp.Push("NotEqual", float32SliceToInterfaceSlice(values)), Exec: func(hit *hitImpl) error {
		return minitest.Error.NotEqual(v.valueCallback(hit), float32SliceToInterfaceSlice(values)...)
	}}
}
func (v *expectFloat32) OneOf(values ...float32) IStep {
	return &hitStep{Trace: ett.Prepare(), When: ExpectStep, CallPath: v.cp.Push("OneOf", float32SliceToInterfaceSlice(values)), Exec: func(hit *hitImpl) error {
		return minitest.Error.OneOf(v.valueCallback(hit), float32SliceToInterfaceSlice(values)...)
	}}
}
func (v *expectFloat32) NotOneOf(values ...float32) IStep {
	return &hitStep{Trace: ett.Prepare(), When: ExpectStep, CallPath: v.cp.Push("NotOneOf", float32SliceToInterfaceSlice(values)), Exec: func(hit *hitImpl) error {
		return minitest.Error.NotOneOf(v.valueCallback(hit), float32SliceToInterfaceSlice(values)...)
	}}
}
func (v *expectFloat32) GreaterThan(value float32) IStep {
	return &hitStep{Trace: ett.Prepare(), When: ExpectStep, CallPath: v.cp.Push("GreaterThan", []interface{}{value}), Exec: func(hit *hitImpl) error {
		l := v.valueCallback(hit)
		if l <= value {
			return minitest.Error.Errorf("expected %f to be greater than %f", l, value)
		}
		return nil
	}}
}
func (v *expectFloat32) LessThan(value float32) IStep {
	return &hitStep{Trace: ett.Prepare(), When: ExpectStep, CallPath: v.cp.Push("LessThan", []interface{}{value}), Exec: func(hit *hitImpl) error {
		l := v.valueCallback(hit)
		if l >= value {
			return minitest.Error.Errorf("expected %f to be less than %f", l, value)
		}
		return nil
	}}
}
func (v *expectFloat32) GreaterOrEqualThan(value float32) IStep {
	return &hitStep{Trace: ett.Prepare(), When: ExpectStep, CallPath: v.cp.Push("GreaterOrEqualThan", []interface{}{value}), Exec: func(hit *hitImpl) error {
		l := v.valueCallback(hit)
		if l < value {
			return minitest.Error.Errorf("expected %f to be greater or equal than %f", l, value)
		}
		return nil
	}}
}
func (v *expectFloat32) LessOrEqualThan(value float32) IStep {
	return &hitStep{Trace: ett.Prepare(), When: ExpectStep, CallPath: v.cp.Push("LessOrEqualThan", []interface{}{value}), Exec: func(hit *hitImpl) error {
		l := v.valueCallback(hit)
		if l > value {
			return minitest.Error.Errorf("expected %f to be less or equal than %f", l, value)
		}
		return nil
	}}
}
func (v *expectFloat32) Between(min, max float32) IStep {
	return &hitStep{Trace: ett.Prepare(), When: ExpectStep, CallPath: v.cp.Push("Between", []interface{}{min, max}), Exec: func(hit *hitImpl) error {
		l := v.valueCallback(hit)
		if l < min || l > max {
			return minitest.Error.Errorf("expected %f to be between %f and %f", l, min, max)
		}
		return nil
	}}
}
func (v *expectFloat32) NotBetween(min, max float32) IStep {
	return &hitStep{Trace: ett.Prepare(), When: ExpectStep, CallPath: v.cp.Push("NotBetween", []interface{}{min, max}), Exec: func(hit *hitImpl) error {
		l := v.valueCallback(hit)
		if l >= min && l <= max {
			return minitest.Error.Errorf("expected %f not to be between %f and %f", l, min, max)
		}
		return nil
	}}
}