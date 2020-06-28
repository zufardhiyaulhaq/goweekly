# Goweekly
Get data from golangweekly.com and send to telegram via bot.

### Goweekly Scrapper
```
export GITHUB_TOKEN=YOURTOKEN
export REPOSITORY=USERNAME/REPOSITORY

make run-scrapper
```

### Goweekly Telegram Bot
```
export GITHUB_TOKEN=YOURTOKEN
export REPOSITORY=USERNAME/REPOSITORY
export TELEGRAM_TOKEN=TOKEN
export TELEGRAM_CHATID=CHATID

cd lib/dispatcher
ruby main.rb
```

