package ui

import (
	"context"
	"errors"
	"strings"

	log "github.com/myLogic207/gotils/logger"
)

var (
	ErrCommandNotFound = errors.New("command not found")
)

type ShellWrapper struct {
	logger        log.Logger
	knownCommands map[string]func(context.Context, []string) ([]byte, error)
}

func NewShellWrapper(logger log.Logger) *ShellWrapper {
	commands := map[string]func(context.Context, []string) ([]byte, error){
		"echo": echo,
	}
	return &ShellWrapper{
		logger:        logger,
		knownCommands: commands,
	}
}

func (sw *ShellWrapper) Execute(ctx context.Context, command string) ([]byte, error) {
	// check if command is known
	sw.logger.Debug(ctx, "Executing command: %s", command)
	parts := strings.Split(command, " ")
	if cmd, ok := sw.knownCommands[parts[0]]; ok {
		return cmd(ctx, parts[1:])
	} else {
		return nil, ErrCommandNotFound
	}
}

func echo(ctx context.Context, args []string) ([]byte, error) {
	return []byte("echo: " + strings.Join(args, " ")), nil
}
