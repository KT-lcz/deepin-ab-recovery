// Code generated by "dbusutil-gen -type Manager manager.go"; DO NOT EDIT.

package main

func (v *Manager) setPropBackingUp(value bool) (changed bool) {
	if v.BackingUp != value {
		v.BackingUp = value
		v.emitPropChangedBackingUp(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedBackingUp(value bool) error {
	return v.service.EmitPropertyChanged(v, "BackingUp", value)
}

func (v *Manager) setPropRestoring(value bool) (changed bool) {
	if v.Restoring != value {
		v.Restoring = value
		v.emitPropChangedRestoring(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedRestoring(value bool) error {
	return v.service.EmitPropertyChanged(v, "Restoring", value)
}

func (v *Manager) setPropConfigValid(value bool) (changed bool) {
	if v.ConfigValid != value {
		v.ConfigValid = value
		v.emitPropChangedConfigValid(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedConfigValid(value bool) error {
	return v.service.EmitPropertyChanged(v, "ConfigValid", value)
}