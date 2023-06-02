# Welcome to the Crypto Bot
The Crypto Bot retrieve some information about to major crypto currency Bitcoin and Ethereum. Enough to 
digit the name of the Crypto currency and the bot will show for exsample :
Symbol: BTC
Price in USD: 29413.41
Variation of price in 24H: 0.42%

For use the bot in local mode you need to clone the repo and install the dependencies:
```bash 
git clone https://github.com/Skrypnyk81/Dev-ops-prometheus.git
go install
```
You must have a working telegram token to insert in variable scope with command:
```bash
read -s TELE_TOKEN
{{insert token}}
export TELE_TOKEN
```
After build the bot with command:
```bash
go build -ldflags "-X="github.com/Skrypnyk81/Dev-ops-prometheus/cmd.appVersion=v1.0.3
```
For starting the binary file of bot run command:
```bash
./Dev-ops-prometheus start
```
You must see output: kbot started v1.0.3 

You can try bot with this URL https://t.me/titanic_prediction_bot
