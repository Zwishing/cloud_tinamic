package model

type PublishRequest struct {
	ServiceName     string `json:"serviceName" validate:"required"`
	SourceCategory  string `json:"sourceCategory" validate:"required,oneof= vector imagery"`
	SourceKey       string `json:"sourceKey" validate:"required"`
	ServiceCategory string `json:"serviceCategory"`
	Description     string `json:"description"`
}
