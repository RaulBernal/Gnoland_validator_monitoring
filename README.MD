# GnoWatch: a validator signing  tool
Active monitoring asking last blockchain block for signed/commited blocks.

> One pic.... 
```
(venv) {raul@iMac-de-Raul:check_validators_v2}$ go run .
```
*Go Bot started, we are going to check some VALIDATORS every 60 seconds*

- Checking: `g1zmua93u53na9xsvtmlh5d8p7pm0w7p52ehewc9` - *Validator 1*
- Checking: `g1e5sxezpafa8lcv5xu3nmw4plz30mnepq2wv9xs` - *Validator AviaOne*
- Checking: `g1k9cp2fz6n3fa4zftx7cz6wv2hwzaz6lkr0xk7z` - *Validator 2*

👍 g1zmua93u53na9xsvtmlh5d8p7pm0w7p52ehewc9 is signing

❌ *Validator AviaOne* is not signing!!
`g1e5sxezpafa8lcv5xu3nmw4plz30mnepq2wv9xs`
cc: @Raul_CTO_AviaOne @Aviaone

👍 g1k9cp2fz6n3fa4zftx7cz6wv2hwzaz6lkr0xk7z is signing


## Go version
### Edit config.go with your Telegram's secrets.
- Replace with your Token and Group.

### Run 
```
go run .
```

### Build a binary
```
go build -o validator-monitor .
./validator-monitor
```

## Python version
### Edit config.go with your Telegram's secrets.
- Replace with your Token and Group.

### Install requests module
```
pip install requests
```

### Run 
```
python3 validator-monitor.py
```


## Cudos to the first version (developer of the BASH tool for Cosmos):
- https://github.com/DDeAlmeida/BitCannaTools
