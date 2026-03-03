package main

// Telegram Bot configuration
const (
	TelegramBotToken = "123456789:AABBCCDDEEFFaabbccddeeff-1234567890"
	TelegramChatID   = "-1001234567890" // use negative ID for groups/channels
)


// address monitoring variables
var validatorsToMonGnoland = []Validator{
	{Name: "Validator 1",       Address: "g1zmua93u53na9xsvtmlh5d8p7pm0w7p52ehewc9",  Telegram: "@Raul_CTO_AviaOne"},
	{Name: "Validator AviaOne", Address: "g1e5sxezpafa8lcv5xu3nmw4plz30mnepq2wv9xs",  Telegram: "@Raul_CTO_AviaOne @Aviaone"},
	{Name: "Validator 2",       Address: "g1k9cp2fz6n3fa4zftx7cz6wv2hwzaz6lkr0xk7z", Telegram: "@Raul_CTO_AviaOne"},
}