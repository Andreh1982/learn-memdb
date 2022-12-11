package worker

// SqsEntryEntity is the interface that receive message from SQS
type SqsEntryEntity struct {
	Name            string `json:"name"`
	Description     string `json:"description" validate:"required,max=140"`
	Severity        string `json:"severity" validate:"required"`
	Percentage      int    `json:"percentage" validate:"required"`
	Time            int    `json:"time" validate:"required"`
	Type            string `json:"type" validate:"required"`
	Threshold       int    `json:"threshold" validate:"required"`
	NamespaceUUID   string `json:"namespace_uuid"`
	Namespace       string `json:"namespace"`
	ApplicationUUID string `json:"application_uuid" validate:"required"`
	Application     string `json:"application"`
	UUID            string `json:"uuid"`
	Active          bool   `json:"active"`
	Key             string `json:"key"`
	Priority        string `json:"priority"`
}
