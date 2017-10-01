# swgoh.gg crawler and CLI tool

For Star Wars Galaxy of Heroes players and tech savvy,
tools provided in this reposity allows them to crawl and use
their data synced to https://swgoh.gg website.

## API for Go lang

A package that crawls and fetch data is available at

	ronoaldo.gopkg.net/swgoh/swgohgg

If you aim to build a tool that uses the website data,
please first ask site authors for their approval.

## CLI

You can use this command line interface to parse your
characters and mods.

### Install or Download

If you just want to use it, go to the [Relase Page](https://github.com/ronoaldo/swgoh/releases)
and download the one for your operating system.

From the command prompt, you can also run this commands if
you have Go setup and ready:

	go get ronoaldo.gopkg.net/swgoh/cmd/swgoh

To see the full list of available options, run:

	swgoh --help

### Ships list

To list your ships, you can use the following
command:

	swgoh -profile ronoaldo -ships

The result is a CSV list of your ships, their
level, and stars.

### Character list

To list your characters, you can use the following
command:

	swgoh -profile ronoaldo -chars

The result is a CSV list of your characters, their
level, stars and gear level.

### Mods and Mod Set suggestions

The tool has some experimental mod suggestion, and is also
capable to export your mdos into CSV format for your
spreadsheet lovers.

	swgoh -profile ronoaldo -mods

You can also ask the tool to suggest a set that contains a
maximum ammount of a given stat (ignoring bonus sets!)

	swgoh -profile ronoaldo -mods -max-set 'Critical Chance'

Finally, you can ask the tool to optimize a set by trial and error.
This is currently a dumb deep search in all possible mod sets and
combinations. The algorithm is very simplified, does no optimizations
and is *very slow*, but works. Often, max-set is enough but if you
want to super maximize a given stat:

	swgoh -profile ronoaldo -mods -optimize-set 'Critical Chance'

### Caching

Caching is performed in your HOME folder, in a file named
swgoh.*yourprofile*.*mods or roster*.json

You can delete these files to update the data from the website.
Currently, the cache does not expires, but this can be changed
in the future.
