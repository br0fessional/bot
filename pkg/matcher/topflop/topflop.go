package topflop

import (
	"fmt"
	"github.com/br0-space/bot/interfaces"
	"github.com/br0-space/bot/pkg/matcher/abstract"
	"github.com/br0-space/bot/pkg/telegram"
	"regexp"
	"strconv"
	"strings"
)

const identifier = "topflop"
const defaultLimit = 10

var pattern = regexp.MustCompile(`(?i)^/(top|flop)(@\w+)?($| )(\d+)?`)

var help = []interfaces.MatcherHelpStruct{{
	Command:     `top`,
	Description: `Zeigt eine Liste der am meisten geplusten Begriffe an.`,
	Usage:       `/top <optional: Anzahl der Einträge>`,
	Example:     `/top 10`,
}, {
	Command:     `flop`,
	Description: `Zeigt eine Liste der am meisten geminusten Begriffe an.`,
	Usage:       `/flop <optional: Anzahl der Einträge>`,
	Example:     `/flop 10`,
}}

const template = "```\n%s\n```"

type Matcher struct {
	abstract.Matcher
	repo interfaces.PlusplusRepoInterface
}

func MakeMatcher(
	logger interfaces.LoggerInterface,
	repo interfaces.PlusplusRepoInterface,
) Matcher {
	return Matcher{
		Matcher: abstract.MakeMatcher(logger, identifier, pattern, help),
		repo:    repo,
	}
}

func (m Matcher) Process(messageIn interfaces.TelegramWebhookMessageStruct) ([]interfaces.TelegramMessageStruct, error) {
	match := m.GetCommandMatch(messageIn)
	if match == nil {
		return nil, fmt.Errorf("message does not match")
	}

	cmd := match[0]
	limit := defaultLimit
	if match[3] != "" {
		res, err := strconv.ParseInt(match[3], 10, 0)
		if err != nil {
			return nil, err
		}
		limit = int(res)
	}

	var records []interfaces.Plusplus
	var err error
	switch cmd {
	case "top":
		records, err = m.repo.FindTops(limit)
	case "flop":
		records, err = m.repo.FindFlops(limit)
	}
	if err != nil {
		return nil, err
	}

	return makeReplies(records, messageIn.ID)
}

func makeReplies(records []interfaces.Plusplus, messageID int64) ([]interfaces.TelegramMessageStruct, error) {
	var lines []string
	for _, record := range records {
		lines = append(lines, fmt.Sprintf(
			"%5d | %s",
			record.Value,
			telegram.EscapeMarkdown(record.Name),
		))
	}

	text := fmt.Sprintf(
		template,
		strings.Join(lines, "\n"),
	)

	return []interfaces.TelegramMessageStruct{
		telegram.MakeMarkdownReply(text, messageID),
	}, nil
}