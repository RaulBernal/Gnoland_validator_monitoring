package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const timeMsgSeconds = 60

// https://test11.testnets.gno.land/r/sys/validators/v2
const urlGnoland = "https://rpc.test11.testnets.gno.land/block"

// address monitoring variables
type Validator struct {
	Name     string
	Address  string
	Telegram string
}

// JSON response structs

type BlockResponse struct {
	Result struct {
		Block struct {
			LastCommit struct {
				Precommits []*Precommit `json:"precommits"`
			} `json:"last_commit"`
		} `json:"block"`
	} `json:"result"`
}

type Precommit struct {
	ValidatorAddress string `json:"validator_address"`
}

// Telegram

type telegramPayload struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func sendTelegram(message string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", TelegramBotToken)

	payload := telegramPayload{
		ChatID:    TelegramChatID,
		Text:      message,
		ParseMode: "Markdown",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("[Telegram] Error building payload: %v\n", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("[Telegram] Connection error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		fmt.Printf("[Telegram] Error sending message: %s\n", respBody)
	}
}

// Bot logic

func welcomeMsg() {
	msg := fmt.Sprintf("*Go Bot started, we are going to check some VALIDATORS every %d seconds*\n\n", timeMsgSeconds)
	for _, v := range validatorsToMonGnoland {
		msg += fmt.Sprintf("- Checking: `%s` - *%s*\n", v.Address, v.Name)
	}
	fmt.Println(msg)
	sendTelegram(msg)
}

func checkSigners() {
	resp, err := http.Get(urlGnoland)
	if err != nil {
		errorMsg := "⚠️ An error occurred getting the last block"
		fmt.Println("\n" + errorMsg)
		sendTelegram(errorMsg)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("\n⚠️ An error occurred reading the response body")
		return
	}

	var blockResp BlockResponse
	if err := json.Unmarshal(body, &blockResp); err != nil {
		fmt.Println("\n⚠️ An error occurred parsing the JSON response")
		return
	}

	precommits := blockResp.Result.Block.LastCommit.Precommits

	// New var with signers.
	signerAddresses := make(map[string]struct{})
	for _, s := range precommits {
		if s == nil { // This Validator is not signing...
			continue
		}
		signerAddresses[s.ValidatorAddress] = struct{}{}
	}

	for _, validator := range validatorsToMonGnoland {
		if _, found := signerAddresses[validator.Address]; found {
			fmt.Printf("👍 %s is signing\n\n", validator.Address)
		} else {
			alert := fmt.Sprintf(
				"❌ *%s* is not signing!!\n`%s`\ncc: %s",
				validator.Name, validator.Address, validator.Telegram,
			)
			fmt.Println(alert + "\n")
			sendTelegram(alert)
		}
	}
}

func main() {
	welcomeMsg()
	for {
		checkSigners()
		time.Sleep(timeMsgSeconds * time.Second)
	}
}
