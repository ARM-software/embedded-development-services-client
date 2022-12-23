package client

//*************************************************************************************
// NOTE: this file is not generated.
// It defines generic models.
//*************************************************************************************

type IModel interface {
	// FetchType returns the resource type
	FetchType() string
	// FetchLinks returns the resource links if present
	FetchLinks() (links any, err error)
	// FetchName returns the resource name if present, or else an error
	FetchName() (string, error)
	// FetchTitle returns the resource title if present, or else an error
	FetchTitle() (string, error)
}
