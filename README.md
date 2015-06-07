Slack - Karma
=============

A simple karma bot for Slack.

## Setup

**Binary**

```bash
$ karma --port=8080 --token=1234 --trigger=karma --storage=dynamodb
```

**Container**

@todo

## Usage

**Give karma**

```
karma nickschuch++
```

```
karma nickschuch+=10
```

**Take karma**

```
karma nickschuch--
```

```
karma nickschuch-=10
```

**Check karma my karma (if matches username)**

```
karma
```

**Check karma of others**

```
karma nickschuch
```
