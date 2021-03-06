package bot_test

import (
	"testing"

	"github.com/andygeiss/betago/business/bot"
)

type MockupController struct {
	Inbound  string
	Outbound string
}

func (c *MockupController) Connect() error {
	return nil
}
func (c *MockupController) Disconnect() error {
	return nil
}
func (c *MockupController) Read(message chan<- string) error {
	message <- c.Inbound
	return nil
}
func (c *MockupController) Write(message string) error {
	c.Outbound = message
	return nil
}

func TestBotIsDisconnectedAtStartup(t *testing.T) {
	b := bot.NewDefaultBot("DefaultBot", nil, nil)
	if b.State() != bot.StateDisconnected {
		t.Error("State should be Disconnected at startup!")
	}
}

func TestBotIsRegisteredAfterSetup(t *testing.T) {
	c := &MockupController{}
	c.Inbound = "REGISTERED"
	b := bot.NewDefaultBot("DefaultBot", c, nil)
	if err := b.Setup(); err != nil {
		t.Errorf("Setup should not return an error! %v", err.Error())
	}
	if b.State() != bot.StateRegistered {
		t.Error("State should be Registered at startup!")
	}
}
