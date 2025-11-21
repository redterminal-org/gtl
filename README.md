# GTL: Gemini Tiny Logs

## This is a Fork of the GTL Gemini Tiny Log viewer by bacardi55
It seems bacardi55 has abandoned the project, so I put up this fork with the
patches included which came in on his development mailing list.

I really don't want to be the new maintainer but I think this app is worth to
stay alive.

I won't develop much on this project myself but you can send in Pull Requests
and I'll include them in this repository. You can also send in patches
via `git send-email` if proper formatted to `fab@redterminal.org` if you don't
have or want a Codeberg.org account. I'll commit the send in patches to this
repository as soon as possible.

There are still glitches with the view of Emojis sometimes so if you want that
repaired, send in a fix.

@bacardi55 still maintains the official list of Tinylogs and you'll find it on
Gemini here:

[bacardi55's list of tinylogs as plaintext file for GTL](gemini://tinylogs.gmi.bacardi55.io/subs.txt)

# Official Repository

Official repository and project is on [sourcehut](https://git.sr.ht/~bacardi55/gtl) (outdated).

Goal: A TUI for the [tinylogs](https://codeberg.org/bacardi55/gemini-tinylog-rfc/src/branch/main) format on the [gemini](gemini.circumlunar.space/) space.

![Gtl TUI screenshot](docs/images/gtl_tui_screenshot.png)

## Quick start:

**EDIT: There are no binarys on this repository, so you have to build it
yourself with Golang!**

*Assuming you downloaded the right binary for your system and put it in your binary, eg: ~/bin*:
```
# will create the configuration and subscription files in ~/.config/gtl/ and subscribe to tinylog aggregator (gemini://tinylogs.gmi.bacardi55.io).
~/bin/gtl
```

To read the global help:
```
gtl --help
```

## Installation

Please read the [installation page](docs/install.md).

## Configuration

Please read the [configuration page](docs/config.md)

## Usage

Gtl can run in 3 modes:
* TUI: An interactive TUI. See the [TUI mode documentation](docs/mode-tui.md).
* CLI: A command line interface. See the [CLI mode documentation](docs/mode-cli.md).
* Gemini: A mode that return a text/gemini output, compatible with the Tinylog format. See the [Gemini mode documentation](docs/mode-gemini.md).

