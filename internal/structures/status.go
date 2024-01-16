package structures

type StatusType string

const (
	PROCESSING StatusType = "processing"
	PENDING    StatusType = "pending"
	CLOSED     StatusType = "closed"
)
