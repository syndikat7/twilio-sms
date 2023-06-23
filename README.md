# Twilio SMS

Simple CLI tool to check for new SMS messages from Twilio.

## Requirements

Create twilio configuration file in your home directory `~/.twilio.yaml` with the following content:

```yaml
Account_SID: ""
Auth_Token: ""
```

## How to use it
```
Usage:
twilio-sms receive [flags]

Flags:
-f, --from int    Timestamp to start from (default 1687257732)
-h, --help        help for receive
-l, --limit int   Limit of messages to fetch (default 20)
-w, --watch       Watch for new messages

Global Flags:
-c, --config string   config file (default is $HOME/.twilio.yaml)
```
