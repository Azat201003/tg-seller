package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/markup"
	"github.com/gotd/td/telegram/message/peer"
	"github.com/gotd/td/tg"
)

func main() {

	dispatcher := tg.NewUpdateDispatcher()
	var opts telegram.Options
	opts = telegram.Options{
		UpdateHandler: dispatcher,
	}
	if os.Getenv("PROXY_ENABLE") == "1" {
		opts.Resolver = getResolver()
	}
	client := telegram.NewClient(APP_ID, APP_HASH, opts)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	err := client.Run(ctx, func(ctx context.Context) error {
		log.Println("Successfully started!")

		if _, err := client.Auth().Bot(ctx, os.Getenv("BOT_TOKEN")); err != nil {
			return err
		}
		state, err := client.API().UpdatesGetState(ctx)
		if err != nil {
			return err
		}

		log.Println(state)

		api := tg.NewClient(client)
		sender := message.NewSender(api)

		dispatcher.OnNewMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewMessage) error {
			message, ok := update.Message.(*tg.Message)
			if !ok || message.Out {
				return nil
			}
			log.Println(message.Message)
			if message.Message == "/start" {
				inputPeer, _ := peer.EntitiesFromUpdate(e).ExtractPeer(message.PeerID)
				_, err := sender.To(inputPeer).Markup(markup.InlineKeyboard(
					markup.Row(markup.Callback("Найти товары", []byte("search0:[]"))),
					markup.Row(markup.Callback("Разместить товар", []byte("place0"))),
				)).Text(ctx, START_MESSAGE)
				log.Println(err)

				return err
			}
			return nil
		})

		dispatcher.OnBotCallbackQuery(func(ctx context.Context, e tg.Entities, update *tg.UpdateBotCallbackQuery) error {
			if strings.HasPrefix(string(update.Data), "search0:") {
				log.Println(string(update.Data))
				filters := getAllFilters()

				selectedFiltersBlob, _ := strings.CutPrefix(string(update.Data), "search0:")
				var selectedFilters []int64
				json.Unmarshal([]byte(selectedFiltersBlob), &selectedFilters)

				var rows []tg.KeyboardButtonRow
				rows = append(rows, markup.Row(markup.Callback(SEARCH_BUTTON_1, []byte("search1:"+selectedFiltersBlob))))
				for _, filter := range filters {
					if !slices.Contains(selectedFilters, filter.FilterId) {
						selectedFiltersBytes, _ := json.Marshal(append(slices.Clone(selectedFilters), filter.FilterId))
						rows = append(rows, markup.Row(markup.Callback(filter.Name, fmt.Appendf(nil, "search0:%v", string(selectedFiltersBytes)))))
					} else {
						selectedFiltersBytes, _ := json.Marshal(slices.DeleteFunc(
							slices.Clone(selectedFilters),
							func(el int64) bool { return el == filter.FilterId },
						))
						rows = append(rows, markup.Row(markup.Callback(filter.Name+" ✅", fmt.Appendf(nil, "search0:%v", string(selectedFiltersBytes)))))
					}
				}
				inputPeer, _ := peer.EntitiesFromUpdate(e).ExtractPeer(update.Peer)
				_, err := sender.To(inputPeer).Markup(markup.InlineKeyboard(rows...)).Text(ctx, FIND_MESSAGE)
				log.Println(err)

				return err
			} else if strings.HasPrefix(string(update.Data), "search1:") {
				log.Println(string(update.Data))
				selectedFiltersBlob, _ := strings.CutPrefix(string(update.Data), "search1:")
				var selectedFilters []int64
				json.Unmarshal([]byte(selectedFiltersBlob), &selectedFilters)

				advertises := getFilteredAdvertises(selectedFilters)

				inputPeer, _ := peer.EntitiesFromUpdate(e).ExtractPeer(update.Peer)
				builder := sender.To(inputPeer)
				for _, advertise := range advertises {
					builder.Markup(markup.InlineKeyboard(markup.Row(markup.URL("Перейти", advertise.Link)))).Text(ctx, advertise.Name)
				}

				return err
			}
			return nil
		})

		<-ctx.Done()
		return nil
	})
	if err != nil {
		log.Fatalf("Client error: %v", err)
	}
	fmt.Println("Client finished.")
}
