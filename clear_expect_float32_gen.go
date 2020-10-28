// +build !generate

// ⚠️⚠️⚠️ This file was autogenerated by generators/clear/clear ⚠️⚠️⚠️ //
package hit

import errortrace "github.com/Eun/go-hit/errortrace"

// IClearExpectFloat32 provides methods to clear steps.
type IClearExpectFloat32 interface {
	IStep
	// Between clears all matching Between steps
	Between(value ...float32) IStep
	// Equal clears all matching Equal steps
	Equal(value ...float32) IStep
	// GreaterOrEqualThan clears all matching GreaterOrEqualThan steps
	GreaterOrEqualThan(value ...float32) IStep
	// GreaterThan clears all matching GreaterThan steps
	GreaterThan(value ...float32) IStep
	// LessOrEqualThan clears all matching LessOrEqualThan steps
	LessOrEqualThan(value ...float32) IStep
	// LessThan clears all matching LessThan steps
	LessThan(value ...float32) IStep
	// NotBetween clears all matching NotBetween steps
	NotBetween(value ...float32) IStep
	// NotEqual clears all matching NotEqual steps
	NotEqual(value ...float32) IStep
	// NotOneOf clears all matching NotOneOf steps
	NotOneOf(value ...float32) IStep
	// OneOf clears all matching OneOf steps
	OneOf(value ...float32) IStep
}
type clearExpectFloat32 struct {
	cp callPath
	tr *errortrace.ErrorTrace
}

func newClearExpectFloat32(cp callPath) IClearExpectFloat32 {
	return &clearExpectFloat32{cp: cp, tr: ett.Prepare()}
}
func (v *clearExpectFloat32) trace() *errortrace.ErrorTrace {
	return v.tr
}
func (*clearExpectFloat32) when() StepTime {
	return CleanStep
}
func (v *clearExpectFloat32) callPath() callPath {
	return v.cp
}
func (v *clearExpectFloat32) exec(hit *hitImpl) error {
	if err := removeSteps(hit, v.callPath()); err != nil {
		return err
	}
	return nil
}
func (v *clearExpectFloat32) Between(value ...float32) IStep {
	return removeStep(v.callPath().Push("Between", float32SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat32) Equal(value ...float32) IStep {
	return removeStep(v.callPath().Push("Equal", float32SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat32) GreaterOrEqualThan(value ...float32) IStep {
	return removeStep(v.callPath().Push("GreaterOrEqualThan", float32SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat32) GreaterThan(value ...float32) IStep {
	return removeStep(v.callPath().Push("GreaterThan", float32SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat32) LessOrEqualThan(value ...float32) IStep {
	return removeStep(v.callPath().Push("LessOrEqualThan", float32SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat32) LessThan(value ...float32) IStep {
	return removeStep(v.callPath().Push("LessThan", float32SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat32) NotBetween(value ...float32) IStep {
	return removeStep(v.callPath().Push("NotBetween", float32SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat32) NotEqual(value ...float32) IStep {
	return removeStep(v.callPath().Push("NotEqual", float32SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat32) NotOneOf(value ...float32) IStep {
	return removeStep(v.callPath().Push("NotOneOf", float32SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat32) OneOf(value ...float32) IStep {
	return removeStep(v.callPath().Push("OneOf", float32SliceToInterfaceSlice(value)))
}