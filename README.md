# swgoh.gg crawler and CLI tool

For Star Wars Galaxy of Heroes players and tech-savvy, tools provided in this repository allows them to crawl and use their data synced to the https://swgoh.gg website.

## API for Go lang

A package that crawls and fetches data is available at

    ronoaldo.gopkg.net/swgoh/swgohgg

If you aim to build a tool that uses the website data, please first ask site authors for their approval. Also, keep in mind that they may pose rate-limit for your requests in order to prevent website overload.

## CLI

You can use this command line interface to parse your characters, ships, and mods. The data is cached as a JSON file that you can parse and use in other apps. The output to stdout can also be used for several purposes, such as feed data into a spreadsheet.

### Install or Download

If you just want to use it, go to the [Release Page](https://github.com/ronoaldo/swgoh/releases) and download the one for your operating system.

You can also install this commands if you have the Go language tools:

    go get ronoaldo.gopkg.net/swgoh/cmd/swgoh

To see the full list of available options, run:

    swgoh --help

### Ships list

To list your ships, you can use the following command:

    swgoh -profile ronoaldo -ships

The result is a CSV list of your ships, their level, and stars.

### Character list

To list your characters, you can use the following command:

    swgoh -profile ronoaldo -collection

The result is a CSV list of your characters, their level, stars, and gear level.

### Mods and Mod Set suggestions

The tool has some experimental mod suggestions and is also capable of exporting your mods into CSV format for you spreadsheet lovers.

    swgoh -profile ronoaldo -mods

You can also ask the tool to suggest a set that contains the maximum amount of a given statistic (ignoring bonus sets!)

    swgoh -profile ronoaldo -mods -max-set 'Critical Chance'

Finally, you can ask the tool to optimize a set by trial and error. This is currently a dumb deep search in all possible mod sets and combinations. The algorithm is very simplified, does no optimizations and is *very slow*, but works. Often, max-set is enough but if you want to super maximize a given statistic:

    swgoh -profile ronoaldo -mods -optimize-set 'Critical Chance'

### Caching

Caching is performed in your personal folder, `$HOME` on *nix machines, using a file named `swgoh.*yourprofile*.*mods or roster*.json`.

You can delete these files to update the data from the website. Currently, the cache does not expire, but this can be changed in the future. You can also use the `--cache=false` command line switch to get always fresh data.
