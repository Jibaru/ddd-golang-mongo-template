package entities

type Version struct {
	value string
}

func NewVersion(value string) *Version {
	return &Version{
		value: value,
	}
}

func (v *Version) Value() string {
	return v.value
}
