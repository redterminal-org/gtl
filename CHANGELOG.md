*CHANGELOG*

## v1.4.4-fork
### Config: Added option to configure color of timeline links
* Added option to "gtl.toml"" named "tui_color_timeline_links" to configure color of links in the timeline.
* removed ".woodpecker/golang-lint.yaml"
### CHANGELOG.md: Rebuild to preferred format
* Rebuild the CHANGELOG.md from commit history to preferred format

## v1.4.3-fork-BUGFIX
### CHANGELOG.md: Fixed typo in version number
### Merge pull request 'FIX-#4' (#6) from FIX-#4 into main
* Reviewed-on: https://codeberg.org/fab/gtl/pulls/6
### FIX: Issue #4 'tui_color_box' uncommenting/usage crashed GTL on startup
### Updated mode-tui.md in docs to include new CRTL+t (select replied-to entry)

## v1.4.2-fork
### Updated CHANGELOG.md
### Added feature: Select replied-to entry directly with "CTRL+t"
### FIX: translated some german comments to english

## v1.4.0-fork
### Updated CHANGELOG.md and README docs
### Added EXPERIMENTAL OSC52 support
### FIX(defaults): set "tui_max_entries" to 100. Was 0, causing TUI to be SLOW!
### FIX: requests are now made as often as possible until the timeout in core/parser.go

## v1.3.0-fork
### Updated to Go 1.26.4 and upgraded all dependencies to their newest versions.
### Fixed missing entries: Timeout raised to 6s. Requests are retried up to 6 times.

## v1.2.0-fork
### Fixed typo in CHANGELOG.md
### Merge pull request 'tcellv3 merge branch' (#2) from tcellv3 into main
* Reviewed-on: https://codeberg.org/fab/gtl/pulls/2
### Updated CHANGELOG.md
### Did some reformatting with Linter
### Fixed ALT/Enter when sidebar has focus
### removed "replace" line from go.mod
### WIP Change to cview tcell/v3 branch

## v1.1.3-fork
### Updated CHANGELOG.md
### changed exclusion case for Search ("/"), Timeline ("t") and Highlights ("h") keys in Bookmarks page
### Updated Golang dependencies
### Fix regression: Display of Refresh Screen
### Changed timeout for reading tinylogs to 3 seconds
### Updated --help output
### Updated documentation (bookmarks)
### woodpecker-ci: added ldflags -s -w to disable symbol table and disable DWARF generation

## v1.1.2-fork
### Updated CHANGELOG.md
### Made Bookmarks useable with "Open Links", "Open Thread", "Reply", "Delete" and "Open Tinylog"
### Fixed creation of bookmarks file and their display
### added Version string "devel" if building from source
### woodpecker-ci: Added version to binaries
### deleted unneeded Makefile
### .md files: Updated documentation a little
### Fixed some linting messages
### Removed protector from CLI mode preformatted text separator
### Disabled highlight of plain text without defined text to highlight
### woodpeckerci: added Go Linter for Pull Requests

## v1.1.1-HOTFIX
### FIXv2: Regression - Fix #42: Preformatted entries beginning with '#' shouldn't break the output anymore now
### HOTFIX: Regression - Fix #42: Entries with ## within a preformated text are failing

## v1.1.0-fork
### configured woodpecker-ci
### Fixed typo in CHANGELOG.md
### Updated CHANGELOG.md
### Fixed some typos in documentation
### Documentation and images cleaned up
### integrated with woodpecker-ci
### Added support for relative links to gemini mode
### Added support for relative links (first revision)
* Relative links are always relative to the gemini domain of the feed entry Uri, eg:
*   => /images/2025-09-08/Lindenplatz1.jpg Optional description
* will translate to:
*   => gemini://redterminal.org/images/2025-09-08/Lindenplatz1.jpg
*   optional description
* if its from my feed at gemini://redterminal.org/tinylog.gmi
* This is the first revision so there may be bugs in the regexp, which I didn't find yet. But it works in my tests.
### Fixed little bug in regexp removing preformatted mark lines in TUI
### some modernization of sources
### FIXED: Display of preformatted text blocks in TUI mode
### Fixed display of preformatted text in gemini mode. Watch out for regression!
### Added word wrapping
### Removed @GTL header in "gemini" mode
### also recognize "http" links
### Updated comment on international time format on generated config file

## v1.0.2-fork
### Changed date to international format in config generation
### Updated README.md
### Updated README.md
### Updated CHANGELOG.md
### Updated .gitignore
### changed module location to "codeberg.org/fab/gtl"
### Updated README.md and install.md
### Changed gtl.toml sources to contain international date format
### updated go.mod and go.sum
### build: Update cbind and cview
* They were migrated to codeberg, but this means builds using old versions can cause issues as they still point to the old repository.
### Fix crash on replies to posts with +0000 UTC offset
* This fixes a crash when replying to an entry with offset +0000 and also now finds the original entry with offset +0000
### Display only description of link, if available, otherwise display URL
* This will only display the description of a link instead of the whole link with description. If no description is available, the URL is shown. When opening URLs, these are fetched from the original entry. I think, this looks much nicer.
* The patch will also show the thread of replies with an URL in the RE: line, like described in the specs.
### Fixed getting date from selected entry
* Sorry, messed up the subject in the last patch. Anyway this fixes a bug in getting the right date of an entry you want to reply to.
### Add links to replies
* I know that the link is optional in spec but it's so awesome.
### Handle DOS line endings
### Update README
### Update README
### Add .build.yml for sourcehut build
### Update README.md
* Add link to codeberg

## v1.0.1
### Add message for gemini render for aggregator page
### Add new date format…

## v1.0.0
### Add documentation and changelog
### bump version to v1.0.0
### update dependencies
### Fix #43: preformatted text starting with ##
### Fix #51: Ignore an author highlights
### Fix #50 - world between brackets not displayed correctly
### Finish #52 part 3: allow theming of modal buttons
### Fix part of #52 - Open link in link popup
### Fix color for selected tinylog in sidebar
### Add Bookmarks feature in TUI mode.

## v0.7.0
### bump version to v0.7.0
### Fix #49 - Fix link format
### Implements #40: [mode gemini] Add link to author tinylogs
### Implements #48: Add button to go to author tinylog on action modal (alt+enter)
### Fix search #47 hihglight issue
### Fix log message content
### Implements #47: Search
### Add help text for action modal
### Add Thread shortcut in help modal
### Implements #44 - Custom themes
### Update Changelog
### Fix #42: separator in response stub
### Update gitignore to ignore example_config directory (used for tests)
* Update gitignore to ignore example_config directory (used for tests)
### Fix subscription name with space in their name
### Fix #45: Crash when empty line in subscription file
### Update dependencies

## v0.6.0
### bump version to v0.6.0
### fix stub display in modal while cview bug is fixed
### Improve gemtext formatting in modal while cview bug isn't fixed
### Fix entry selection with mouse click
### add an option to display reply stub. Rendering blocked by cview
### Rmove extra empty line in Link Modal
### Implement a modal for entry details via alt+enter
### try to improve formatting

## v0.6.0-alpha
### bump version to v0.6.0-alpha
### Add changelog
### Update README and break it into multiple doc files
### Add Changelog
### Add custom mouse handler on Timeline box to avoid wrong entry selection on left click. Due to an issue in cview
### fix bug with selected entry when changing mode
### Fix thread modal formating because of cview bug
### Allow link selection with 'O'
### Improve modal display when opening multiple links or seeing a thread
### Fix Thread feature #36
### Fix entry selection limits
### Finish implementing stubs feature, resolves #28
### Remove date format duplicate
### Use gemini://tinylogs.gmi.bacardi55.io/ as default subscripton
### Start implementing #36: Show the tinylog entry being replied to
* Blocked by https://code.rocketnine.space/tslocum/cview/issues/69
* and https://code.rocketnine.space/tslocum/cview/issues/71
### Resolves #37: Implement a gemini compatible output
### Fix limit bug in cli mode
### Fix error for missing file
### fix esc shortcut behaviour
### Implements #35: parser v2
### clean code
### improve gemtext format
### Improve stubs for response #28
### Improve gemtext format
### Update dependencies
### update format for lists in text/gemini
### Start work to open links in entries
### Continue implementing #28: Reply stubs prework done. Rest Blocked for cview bug
### remove logs
### Start #28: Implements optional stubs for new entry
### Add J and K to README
### Add optional limit to number of displayed entries
### Remove deprecated footer related code
### Fix #29: Timeline refresh was previous refresh
### Implements region in timeline TextView

## v0.5.2
### bump version to v0.5.2
### Fix date parsing error
### Fix title in readme
### Fix case sensitivity in avatar/author metadata
### Fix readme typo

## v0.5.1
### bump version to v0.5.1
### Harmonized terminology around tinylogs and improve panels title
### Improve gemtext formatting for responses.
### Improve focus changes
### remove message post script if successful and add config to refresh feeds after editing entry
### Remove modal message after editing tinylog if no post script is defined.
### Improve main panel title by adding filtered tinylog name
### fix readme
### fix readme

## v0.5.0
### bump version to v0.5.0
### Hide run script button when not set
### update readme
### update readme
### Implements #10 - [TUI] Add new entry in tinylog
### cleaning
### Fix go formatting
### fix focus issue
### WIP - #10 [TUI] Add ability to add new entry in tinylog
### WIP - #10 [TUI] Add ability to add new entry in tinylog

## v0.4.8
### bump version to v0.4.8
### update readme
### Update TUI screenshot
### Add Mac/Windows help message
### Resolves #16: [TUI] Add hability to filter out tinylogs
### Add message while refreshing feeds
### improve padding
### Add UTC offset for date format

## v0.4.7
### bump version to v0.4.7
### fix gemtext formatting
### Add option for emoji usage. Also fixed some gemtext formatting
### Improve title readability
### Add sidebar toggle command in help box

## v0.4.6
### bump version to v0.4.6
### Improve TUI format #20
### Add help message in All sub secondary text
### add gemtext formatting in highlights
### Update sidebar color to match timeline color formatting
### Start implementing #20: Add formating for links and quote
### Add s shortcut in README
### Add shortcut to toggle sidebar display
### change dead link emoji
### Resolves #17: [TUI] Don't quit gtl when displaying help with q
### Add new date format…
### Reformat TUI Sub list, resolves #15

## v0.4.5
### bump version to v0.4.5
### TUI - Fix separator between old and new entries
### update readme for right version of go needed

## v0.4.4
### Bump to version v0.4.4
### default mode is now TUI
### Add display/hide help box and remove header
### t shortcut select back 'all sub' in sidebar
### reorganize ui code
### update shortcuts in readme
### move to cbind for TUI shortcuts
### Re-organize code before change in keybindings
### reduce size for header
### Resolves #13: Add color on border for focus panel
### Resolves #12: Add last refresh in timeline. Hide footer for now before it is used again
### Fix #1: Create config dir at first launch and fix first launch
### Implements #9: Show tinylogs with 0 compatible entries

## v0.4.3
### bump version to v0.4.3
### Fix separator regression
### fix code standards
### fix regression in filtering
### Add TAB in usage
### fix focus

## v0.4.2
### Bump to version v0.4.2
### update README
### Add indicator between new and old entries
### Add last refresh time in footer
### Refresh subscriptions status when refreshing feeds
### fix background
### Fix background to black
### Fix makefile
### Fix readme
### Fix go version in readme

## v0.4.1
### bump version to v0.4.1
### Improve readme and add status emoji

## v0.4.0
### bump version to v0.4.0
### add emoji in case of feed errors
### clean run code
### TUI v1
### mvt tui
### Add colors and hightlights in timeline
### Add mode option
### First try at TUI
### add new date format

## v0.3.1
### v0.3.1
### Update README with cli limit configuration
### Implement cli limit option
### add date format and small fix in case of only 1 entry in a tinylog
### fix formatting
### fix issues for noname tinylogs
### Fix mistake in Makefile
### update readme
### remove old comments
### udpate readme

## v0.3.0
### Bump to v0.3.0
### Update readme
### update dependencies
### Add make dependencies
### Add add and rm commands for subscription management
### Add unknown for author when no author metadata and no optional title

## v0.2.0
### update makefile for v0.2.0
### add screenshot in readme
### fix formatting
### fix metatags discovery
### improve elapsed time format in cli
### Add highlights config in README
### Add colors and highlights for CLI display
### update readme
### update readme
### update readme
### update readme
### update readme to point to release page

## v0.1.1
### Fix version display
### clean makefile
### Add binaries download in README

## v0.1.0
### update makefile
### update makefile
### update makefile
### update makefile
### Add Makefile
### Add version
### Improve README
### Format according to go standards
### Improve README
### remove debug
### replace --cli with --mode option
### Add log file configuration
### Add limit option for cli
### First version of cli display working with a sorted stream
### Fixing issue for entry stating with an empty line
### Parse tinylog entry date and content
### Start parsing tinylog files
### Retrieve all tinylog content (todo: parse)
### Initial commit
### Initial commit
