# Red Time

## Description

This is a small piece of software to run on a raspberry to help manage a photographic enlarger by using a f-stop like timer.

It provides an API to setup up an endpoint that will accept f-stops and will transform that into seconds for working with the enlarger in a more accurate way.

## How to build

A Makefile is provided which will help building, running and installing the software, the only requirement is to have an environment variable called `REDTIMEHOST` that will point to the right host where to copy and execute the command.

A simple `make run` will take care of everything and run the program in the host.

In order to make everything easier, I'll recommend to setup your ssh config to login using SSH keys, that way with the environment variable from above no password will be needed.

## API

Usually port will be `:3000` at $REDTIMEHOST

### Commands

This is a POST command, use json as the body request.

#### start
URL: host:port/api/commands/start
Parameters:
- `fstops` _float_ number of f-stops to use the enlarger.
- `delay` _float_ number of seconds until the enlarger is turned on. This is useful to have some time to turn off the screen in case of color enlargement or if no red filter is used for black and white photography