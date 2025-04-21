# Knot
Minimal terminal email client with a twist.

## Prerequisites

To use knot you will need an email account set with an app-specific password. Please read the following articles for more information:
- [apple app-specific password](https://support.apple.com/en-us/102654)
- [gmail app-specific password](https://support.google.com/accounts/answer/185833?hl=en)

This password will be used to log into knot.

## Installation
1. Clone the repo to your computer
2. Run `go build`

## Tutorial
1. First check if your account is setup correctly by running the `boxes` command.
```
knot> boxes
```
2. Next pull the five most recent emails using the `pull` command.
```
knot> pull
```
3. Finally run the `list` command to see what was pulled.
```
knot> list
```
4. To read an email use the `read` command along with an index.
```
knot> read 0
```
5. To send an email use the `send` command and follow the prompts.
```
knot> send
```

# Further Reading

## Design Goals
1. Minimal: Keep dependencies to a minimum.
2. Modularity: Can add features when needed.
3. Portability: Attempts to run anywhere Go is supported.

## Name Origin
I'm currently playing [death stranding](https://en.wikipedia.org/wiki/Death_Stranding) and the concept of a knot is a central theme in the game. I won't spoil anything, so go play the game!
