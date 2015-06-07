Slack - Karma
=============

A simple karma bot for Slack.

## Setup

### Slack

* Setup a https://api.slack.com/slash-commands with the key as "/karma".
* Setup an incoming callback and use this URL as the `--callback` config.

### CLI

**Binary**

```bash
$ karma --port=8080 --token=1234 --callback=http://example.com
```

```bash
$ export KARMA_PORT=8080
$ export KARMA_TOKEN=1234
$ export KARMA_CALLBACK=http://example.com
$ karma
```

**Docker**

Available here: https://github.com/nickschuch/karma-docker

```bash
$ cat << EOF > ~/karma.env
KARMA_PORT="8080"
KARMA_TOKEN="123456"
KARMA_STORAGE="memory"
KARMA_NAME="Karma"
KARMA_EMOJI=":slack:"
KARMA_CALLBACK="http://example.com"
EOF
$ docker run --env-file ~/karma.env -p 127.0.0.1:8080:8080 nickschuch/karma
```

## Usage

**Give karma**

```
/karma nickschuch++
```

```
/karma nickschuch+=10
```

**Take karma**

```
/karma nickschuch--
```

```
/karma nickschuch-=10
```

**Check my karma**

```
/karma
```

**Check karma of others**

```
/karma nickschuch
```
