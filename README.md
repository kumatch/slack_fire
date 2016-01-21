slack_fire
======

Send a message to Slack using Incoming Webhooks.


## Installation

```
$ go get github.com/kumatch/slack_fire
```

## Examples

### Send text message

Set text message to arguments with Webhooks URL `-w` flag (and some optional flags).

```
$ slack_fire -whttps://hooks.slack.com/services/xxxxxxxxxxxxxxxxxxxx -ufire -c#notification "A text message"
```

### Set JSON string to stdin

Set JSON data to stdin with `-S` flag for Slack post messages (with Message Attachments). See [https://api.slack.com/incoming-webhooks](https://api.slack.com/incoming-webhooks).

```
$ echo '{ "text": "A text message", "username": "fire", "channel": "#notification" }' | slack_fire -w https://hooks.slack.com/services/xxxxxxxxxxxxxxxxxxxx -S
```
