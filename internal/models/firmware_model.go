package models

type FirmwareModel struct {
	Name         string `json:"name"`
	DeviceID     string `json:"device_id"`
	Version      string `json:"version"`
	ReleaseNotes string `json:"release_notes"`
	ReleaseDate  string `json:"release_date"`
	Url          string `json:"url"`
}

func (f FirmwareModel) ValidateFirmware() map[string]string {
	errores := make(map[string]string)
	if f.Name == "" {
		errores["name"] = "field name is required"
	}

	if f.DeviceID == "" {
		errores["device_id"] = "field device_id is required"
	}

	if f.Version == "" {
		errores["version"] = "field version is required"
	}

	if f.ReleaseNotes == "" {
		errores["release_notes"] = "field release_notes is required"
	}

	if f.ReleaseDate == "" {
		errores["release_notes"] = "field release_date is required"
	}

	if f.Url == "" {
		errores["url"] = "field url is required"
	}

	return errores
}
