// Package main provides a simple bot experience using slack.Adapter with multiple plugin commands and scheduled tasks.
package main

import (
	"context"
	"flag"
	"github.com/oklahomer/go-kasumi/logger"
	"github.com/oklahomer/go-sarah/v4"
	"github.com/oklahomer/go-sarah/v4/alerter/line"
	_ "github.com/oklahomer/go-sarah/v4/slack"
	"github.com/oklahomer/go-sarah/v4/watchers"
	"gopkg.in/yaml.v2"
	"os"
	"os/signal"	
	"syscall"
	"tgbot/telegram"
)

type myConfig struct {
	CacheConfig     *sarah.CacheConfig `yaml:"cache"`
	Telegram        *telegram.Config   `yaml:"telegram"`
	Runner          *sarah.Config      `yaml:"runner"`
	LineAlerter     *line.Config       `yaml:"line_alerter"`
	PluginConfigDir string             `yaml:"plugin_config_dir"`
}

func newMyConfig() *myConfig {
	// Use a constructor function for each config struct, so default values are pre-set.
	return &myConfig{
		CacheConfig: sarah.NewCacheConfig(),
		Telegram:    telegram.NewConfig(),
		Runner:      sarah.NewConfig(),
		LineAlerter: line.NewConfig(),
	}
}

func main() {
	var path = flag.String("config", "", "path to application configuration file.")
	flag.Parse()
	if *path == "" {
		panic("./bin/examples -config=/path/to/config/app.yml")
	}

	// Read a configuration file.
	config := readConfig(*path)

	// When the Bot encounters critical states, send an alert to LINE.
	// Any number of Alerter implementations can be registered.
	sarah.RegisterAlerter(line.New(config.LineAlerter))

	// Set up a storage that can be shared among different Bot implementations.
	storage := sarah.NewUserContextStorage(config.CacheConfig)

	// Set up Slack Bot.
	//setupSlack(config.Slack, storage)
	setupTelegram(config.Telegram, storage)

	// Set up some commands.
	//todoCmd := todo.BuildCommand(&todo.DummyStorage{})
	//sarah.RegisterCommand(slack.SLACK, todoCmd)

	// Directly add Command to Bot.
	// This Command is not subject to config file supervision.
	//sarah.RegisterCommand(slack.SLACK, echo.Command)

	// Prepare Sarah's core context.
	ctx, cancel := context.WithCancel(context.Background())

	// Prepare a watcher that reads configuration from filesystem.
	if config.PluginConfigDir != "" {
		configWatcher, _ := watchers.NewFileWatcher(ctx, config.PluginConfigDir)
		sarah.RegisterConfigWatcher(configWatcher)
	}

	// Run.
	err := sarah.Run(ctx, config.Runner)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	select {
	case <-c:
		logger.Info("Stopping due to signal reception.")
		cancel()

	}
}

func readConfig(path string) *myConfig {
	configBody, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := newMyConfig()
	err = yaml.Unmarshal(configBody, config)
	if err != nil {
		panic(err)
	}

	return config
}

func setupTelegram(config *telegram.Config, storage sarah.UserContextStorage) {
	//adapter, err := slack.NewAdapter(config, slack.WithEventsPayloadHandler(slack.DefaultEventsPayloadHandler))
	adapter, err := telegram.NewAdapter(config)
	if err != nil {
		panic(err)
	}

	bot := sarah.NewBot(adapter, sarah.BotWithStorage(storage))

	// Register the bot to run.
	sarah.RegisterBot(bot)
}
