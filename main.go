//   Approver Bot
//   Copyright (C) 2021 Fraud boy bgm)

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func main() {
	bot, err := gotgbot.NewBot(
		os.Getenv("TOKEN"),
		&gotgbot.BotOpts{
			APIURL:      "",
			Client:      http.Client{},
			GetTimeout:  gotgbot.DefaultGetTimeout,
			PostTimeout: gotgbot.DefaultPostTimeout,
		},
	)
	if err != nil {
		fmt.Println("Failed to create bot:", err.Error())
	}
	updater := ext.NewUpdater(
		&ext.UpdaterOpts{
			ErrorLog: nil,
			DispatcherOpts: ext.DispatcherOpts{
				Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
					fmt.Println("an error occurred while handling update:", err.Error())
					return ext.DispatcherActionNoop
				},
				Panic:       nil,
				ErrorLog:    nil,
				MaxRoutines: 0,
			},
		})
	dp := updater.Dispatcher

	
	// Commands
	dp.AddHandler(handlers.NewCommand("start", Start))
	dp.AddHandler(handlers.NewChatJoinRequest(nil, Approve))

	// Start Polling()
	poll := updater.StartPolling(bot, &ext.PollingOpts{DropPendingUpdates: true})
	if poll != nil {
		fmt.Println("Failed to start bot:", poll.Error())
	}

	fmt.Printf("@%s has been sucesfully started\n💝Made by @fbb_alone\n", bot.Username)
	updater.Idle()
}

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveChat.Type != "private" {
		return nil
	}

	user := ctx.EffectiveSender.User
	text := `
<b>Hello <a href="tg://user?id=%v">%v</a></b>
I am a bot for accepting newly coming join requests.

 <b><a href="https://telegra.ph/Accept-Join-Request-Help-03-03">Help Me </a> - Its Help You To know How to Use Me</b> 

Bot made with 💝 by <a href="https://t.me/+Ngd7XKW_pZcxYWY1">Fraud Boy Bgm</a>  for you!

<b> My Developer :</b> <b><a href="t.me/fbb_alone">My Father</a></b>

<b> You Must Join Below channels to Use Me </b>
	`
	ctx.EffectiveMessage.Reply(
		bot,
		fmt.Sprintf(text, user.Id, user.FirstName),
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Main Channel", Url: "https://t.me/+Ngd7XKW_pZcxYWY1"},
					{Text: "Update Channel", Url: "https://t.me/+vz5Bij9vfANhNzNl"},
				}},
			},
			ParseMode:             "html",
			DisableWebPagePreview: true,
		},
	)
	return nil
}

func Approve(bot *gotgbot.Bot, ctx *ext.Context) error {
	_, err := bot.ApproveChatJoinRequest(ctx.EffectiveChat.Id, ctx.EffectiveSender.User.Id)
	if err != nil {
		fmt.Println("Error while approving:", err.Error())
	}
	return nil
}
