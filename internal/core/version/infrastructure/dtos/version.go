package dtos

type VersionDto struct {
	Value string `json:"value"`
}

func NewVersionDto(value string) *VersionDto {
	return &VersionDto{
		value,
	}
}
