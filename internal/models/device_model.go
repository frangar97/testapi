package models

type DeviceModel struct {
	Name string `json:"name"`
}

func (d DeviceModel) ValidateDevice() map[string]string {
	errores := make(map[string]string)
	if d.Name == "" {
		errores["name"] = "field name is required"
	}

	return errores
}
