package abstract

import (
	"fmt"

	"github.com/neovg/kmptnzbot/internal/logger"
	"github.com/neovg/kmptnzbot/internal/telegram"
)

// Create a struct each matcher is inherited from
type Matcher struct{}

// Handle an error happening during matcher execution
func (m Matcher) HandleError(requestMessage telegram.RequestMessage, identifier string, err error) {
	// Print error information to logger
	logger.Log.Error(identifier, err.Error())

	// Send error notification to Telegram
	err = telegram.SendMessage(
		requestMessage,
		telegram.Message{
			Text: fmt.Sprintf("⚠️ %s %s", identifier, err.Error()),
		},
	)
	if err != nil {
		logger.Log.Warning(identifier, err.Error())
	}
}
