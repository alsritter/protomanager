package protomanager

import (
	"context"

	"github.com/jhump/protoreflect/desc"
)

// Provider is used to read, parse and manage Proto files
type Provider interface {
	Start(ctx context.Context, cancelFunc context.CancelFunc) error
	// GetMethod is used to get descriptor of specified grpc path
	GetMethod(name string) (*desc.MethodDescriptor, bool)
}
