# SWGoH API client and CLI for Go

[![Build Status](https://travis-ci.org/ronoaldo/swgoh.svg?branch=master)](https://travis-ci.org/ronoaldo/swgoh)

This project implements API client and command line interface (CLI)
for Star Wars Galaxy of Heroes game. Data is retrieved from third party services
provided by other developers:

* https://swgoh.gg/ and https://swgoh.gg/api/
* https://api.swgoh.help

The CLI uses https://api.swgoh.help and the https://crinolo-swgoh.glitch.me/statCalc helper
to provide rich information about character, ships and mods.

A basic client for https://swgoh.gg/ (website crawling) and https://swgoh.gg/api/
is also provided. as of now this implementation is not extensivelly tested.

## API client for Go lang

The API client for the Go programming language is available using the import

    import "github.com/ronoaldo/swgoh/swgohhelp"

Read the full documentation at https://godoc.org/github.com/ronoaldo/swgoh/swgohhelp

## CLI

You can use this command line interface to parse your characters, ships, and mods.
The data is cached as a JSON file that you can parse and use in other apps.
The output to stdout can also be used for several purposes, such as feed data
into a spreadsheet.

### Install or Download

If you just want to use it, go to the
[Release Page](https://github.com/ronoaldo/swgoh/releases)
and download the one for your operating system.

You can also install the tool using the Go toolchain:

    go get github.com/ronoaldo/swgoh/cmd/swgoh

To see the full list of available options, run:

    swgoh -help

### Authentication

For each invocation you need to provide the https://api.swgoh.help credentials.
The parameters `-u "myuser"` and `-p "mypass"` are mandatory.

### Ships list

To list your ships, you can use the following command:

    swgoh -u "myuser" -p "mypass" -a "allycode" -ships

The result is a CSV list of your ships, their level, and stars.

### Character list

To list your characters, you can use the following command:

    swgoh -u "myuser" -p "mypass" -a "allycode" -characters

The result is a CSV list of your characters, their level, stars, and gear level.

### Mods

You can use the -mods switch to export all your mods to the standard output
formatted as CSV:

    swgoh -u "myuser" -p "mypass" -a "allycode" -mods

### Character and Arena stats

It is possible to display single character or your arena team statistics
using the -stats or -arena switches.

To display statistics from a single character, type:

    swgoh -u "myuser" -p "mypass" -a "allycode" -stats -char Tarkin

And to display the current stats of your arena team, use the -arena option.
Both character and ship arenas will be shown:

    swgoh -u "myuser" -p "mypass" -a "allycode" -arena

# Caching

Caching is done for several API consumed files. 
This serves two pourposes: avoid overloading the API
endpoints and gives you a local copy of the JSON files
for your parsing needs.

Cached data is stored in `$HOME/.config/api.swgoh.help/` folder.
You can change the value using the system variable `$SWGOH_CACHE_DIR`.
This folder is created if does not exists.

In that folder, a file `gamedata.json` holds cached data for several game
related info, such as unit names and other info.

The CLI also stores profile data in `ALLYCODE.json` files in the cache directory.
