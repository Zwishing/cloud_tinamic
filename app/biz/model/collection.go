package model

type PublishRequest struct {
	SourceCategory  string `json:"sourceCategory" validate:"required,oneof= vector imagery"`
	Key             string `json:"key" validate:"required"`
	ServiceCategory string `json:"serviceCategory" validate:"required,oneof= mvt feature"`
}
