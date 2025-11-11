// © 2025 Amir. All rights reserved.
// Licensed under the MIT License with Commons Clause restriction.
// You may use this software freely for non-commercial purposes.
// Commercial use, resale, or offering as part of a paid service
// requires a separate commercial license from Amir.
// Contact: licensing@mcpgo.io

package commands_test

import (
	"context"
	"testing"

	"mcpgo/internal/application/commands"
	"mcpgo/internal/domain/server"
	"mcpgo/internal/domain/shared"
)

// MockServerRepo is a mock implementation of the ServerRepo for testing.
type MockServerRepo struct {
	SaveFunc func(ctx context.Context, s *server.Server) error
}

func (m *MockServerRepo) Save(ctx context.Context, s *server.Server) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(ctx, s)
	}
	return nil
}
func (m *MockServerRepo) FindByID(ctx context.Context, id shared.ID) (*server.Server, error) { return nil, nil }
func (m *MockServerRepo) FindAll(ctx context.Context) ([]*server.Server, error)           { return nil, nil }

// MockEventBus is a mock implementation of the EventBus for testing.
type MockEventBus struct {
	PublishFunc func(ctx context.Context, event shared.DomainEvent) error
}

func (m *MockEventBus) Publish(ctx context.Context, event shared.DomainEvent) error {
	if m.PublishFunc != nil {
		return m.PublishFunc(ctx, event)
	}
	return nil
}

func TestRegisterServerHandler_Handle(t *testing.T) {
	ctx := context.Background()
	mockRepo := &MockServerRepo{}
	mockEventBus := &MockEventBus{}

	var savedServer *server.Server
	mockRepo.SaveFunc = func(ctx context.Context, s *server.Server) error {
		savedServer = s
		return nil
	}

	var publishedEvent shared.DomainEvent
	mockEventBus.PublishFunc = func(ctx context.Context, event shared.DomainEvent) error {
		publishedEvent = event
		return nil
	}

	handler := commands.NewRegisterServerHandler(mockRepo, mockEventBus)

	cmd := commands.RegisterServerCommand{
		Name:     "Test Server",
		Address:  "http://localhost:9000",
		Protocol: "mcp/v1",
	}

	serverID, err := handler.Handle(ctx, cmd)
	if err != nil {
		t.Fatalf("Handle() error = %v, wantErr nil", err)
	}

	if serverID == "" {
		t.Errorf("Expected a server ID, but got empty string")
	}

	if savedServer == nil {
		t.Fatal("Expected server to be saved, but it was nil")
	}
	if savedServer.Name != cmd.Name {
		t.Errorf("Expected saved server name to be %s, got %s", cmd.Name, savedServer.Name)
	}

	if publishedEvent == nil {
		t.Fatal("Expected an event to be published, but it was nil")
	}
	if publishedEvent.EventName() != "domain.server.registered" {
		t.Errorf("Expected event name to be 'domain.server.registered', got '%s'", publishedEvent.EventName())
	}
}
