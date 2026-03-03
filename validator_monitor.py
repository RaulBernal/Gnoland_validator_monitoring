#python3

import requests
import time
from config import TELEGRAM_BOT_TOKEN, TELEGRAM_CHAT_ID, validators_to_mon_gnoland

time_msg = 60

# https://test11.testnets.gno.land/r/sys/validators/v2
url_gnoland = 'https://rpc.test11.testnets.gno.land/block'

# Telegram

def send_telegram(message: str):
    url = f"https://api.telegram.org/bot{TELEGRAM_BOT_TOKEN}/sendMessage"
    payload = {
        "chat_id":    TELEGRAM_CHAT_ID,
        "text":       message,
        "parse_mode": "Markdown",
    }
    try:
        resp = requests.post(url, json=payload, timeout=10)
        if not resp.ok:
            print(f"[Telegram] Error sending message: {resp.text}")
    except Exception as e:
        print(f"[Telegram] Connection error: {e}")

# Bot logic

def welcome_msg():
    msg = ''
    msg_intro = '*Python Bot started, we are going to check some VALIDATORS every ' + str(time_msg) + ' seconds*\n\n'
    for validators in validators_to_mon_gnoland:
        msg += "- Checking: `" + validators['validator'] + "` - *" + validators['name'] + '*\n'
    msg_formated = msg_intro + msg
    print(msg_formated)
    send_telegram(msg_formated)

def check_signers():
    try:
        response_check = requests.get(url_gnoland, headers={"Accept": "application/json"})
    except:
        error_msg = "⚠️ An error occurred getting the last block"
        print("\n" + error_msg)
        send_telegram(error_msg)
        return

    json_response = response_check.json()
    signatures    = json_response["result"]["block"]["last_commit"]["precommits"]

    # New var with signers.
    signer_addresses = {
        s["validator_address"]
        for s in signatures
        if s is not None  # This Validator is not signing...
    }

    for validator in validators_to_mon_gnoland:
        if validator['validator'] in signer_addresses:
            print(f"👍 {validator['validator']} is signing\n")
        else:
            alert = (
                f"❌ *{validator['name']}* is not signing!!\n"
                f"`{validator['validator']}`\n"
                f"cc: {validator['telegram']}"
            )
            print(alert + "\n")
            send_telegram(alert)

if __name__ == "__main__":
    welcome_msg()
    while True:
        check_signers()
        time.sleep(time_msg)
