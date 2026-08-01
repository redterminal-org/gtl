CHANGELOG

## v1.5.3-fork-BUGFIX

### New Commits:
* FIX: Issue #4 'tui_color_box' uncommenting/usage crashed GTL on startup
* Updated mode-tui.md in docs to include new CTRL+t (select replied-to entry)

## v1.4.2-fork

### New Commits:
* Added feature: Select replied-to entry directly with "CTRL+t"
* FIX: translated some german comments to english

## v1.4.0-fork

### New Commits:
* Updated CHANGELOG.md and README docs
* Added EXPERIMENTAL OSC52 support
* FIX(defaults): set "tui_max_entries" to 100. Was 0, causing TUI to be SLOW!
* FIX: requests are now made as often as possible until the timeout in
  core/parser.go

## v1.3.0-fork

### New Commits:
* Updated to Go 1.26.4 and upgraded all dependencies to their newest versions.
* Fixed missing entries: Timeout raised to 6s. Requests are retried up to 6 times

## v1.2.0-fork

### New Commits:
* merged branch "tcellv3"
* Did some reformatting with Linter
* Fixed ALT/Enter when sidebar has focus
* removed "replace" line from go.mod
* WIP Change to cview tcell/v3 branch

## v1.1.3-fork

### New Commits:
* changed exclusion case for Search ("/"), Timeline ("t") and Highlights ("h") keys in Bookmarks page
* Updated Golang dependencies
* Fix regression: Display of Refresh Screen
* Changed timeout for reading tinylogs to 3 seconds
* Updated --help output
* Updated documentation (bookmarks)
* woodpecker-ci: added ldflags -s -w to disable symbol table and disable DWARF generation

## v1.1.2-fork

### New Commits:
* Made Bookmarks useable with "Open Links", "Open Thread", "Reply", "Delete" and "Open Tinylog"
* added Version string "devel" if building from source
* woodpecker-ci: Added version to binaries
* deleted unneeded Makefile
* .md files: Updated documentation a little
* Fixed some linting messages
* Removed protector from CLI mode preformatted text separator
* Disabled highlight of plain text without defined text to highlight
* woodpecker-ci: added Go Linter for Pull Requests

## v1.1.1-HOTFIX

### Bug Fixes
* HOTFIX: Regression - Fix #42: Entries with '#' at the beginning within a
  preformated text are failing. This should work as expected.

## v1.1.0-fork

### New Commits:
* Cleaned up Documentation and broken images
* Added support for relative Links (TUI and gemini mode)
* did some modernization fixes
* FIXED: Display of preformatted text blocks in TUI mode:
  Configuration option: 'tui_color_preformatted = "00A0000"'
* Added word wrapping to TUI mode
* Removed @GTL header in "gemini" mode
* GTL now also recognizes "http" links

## v1.0.2-fork

**Fork of original Release**
I added all the patches sent to bacardi55's dev-mailinglist on sourcehut since
the last release.

### New Commits:
* changed module to 'codeberg.org/fab/gtl'
* Updated to Golang version 1.25.4
* Update cbind and cview
* Fix crash on replies to posts with +0000 UTC offset
* Display only description of link if available, otherwise display URL
* Fixed getting date from selected entry
* Add links to replies
* Handle DOS line endings

## v1.0.0:

**Major Release**

### New features:
* Implements #34: Add bookmarks management
* Implements #51: Ignore highlight list

### Bug Fixes:
* Fix #50: World between bracket are not displayed
* Fix #52 part1: highlight bar in the subscription list does not change color
* Fix #52 part2: color of the buttons in the pop-up window does not change color
* Fix #52 part3: Wrong links in link modal when multiple links
* Fix #42: Entries with ## within a preformated text are failing

## v0.7.0:

### New features:
* Implements #44: Custom theme. See the documentation in [docs/mode-tui-theming.md]
* Implements #47: Search. Shortcut `/` will open a modal asking for a search. It will basically filter entries that contains exactly your search (case insensitive). It isn’t fuzzy or advance searching (more filtering)
* Implements #48: Add button to go to author tinylog on action modal (alt+enter)
* Implements #40: Add tinylog url in the gemini mode output (used for generating gemini://tinylogs.gmi.bacardi55.io)

### Bug Fixes
* Fix 42: Hide new line separator in response stub
* Fix subscription name with space in their name
* Fix #45: Crash when empty line in subscription file
* Fix #49: Allow link without space between `=>` and the actual link.

## v0.6.0:

**Major release**

### Breaking changes:
* Shortcut to create a new entry is now "N" (uppercase) and no longer ctrl+n.

### New features:
* Implements region in timeline TextView: This will allow to navigate between entries using J and K (uppercase)
* #28: stubs for new entry: Clicking on "N" (uppercase) will copy in the clipboard a stub for a new entry with the current date if the option `tui_copy_stub_clipboard` is set (see README).
* #28: stubs for response: When selecting entry with J/K, using "R" (uppercase) will copy a pre-formatted response in your clipboard if the option `tui_copy_stub_clipboard` is set (see README).
* #33: Open links in entries, relates to #23. When using J/K to select an entry, using "O" will automatically open the link in the entry. If more than one link is in the entry, a popup will ask for one (or all) link(s) to open.
* #37: Implement a gemini compatible output
* #36: Show the tinylog entry being replied to: When selecting an entry with J/K, if the entry is a response to another tinylog entry and in the RFC format for a response, it will show the original entry in a popup when using "T" (uppercase)?
* #35: parser v2: Improve parser to be compatible with proposition 15 of the RFC (https://codeberg.org/bacardi55/gemini-tinylog-rfc/issues/15) that allow multiple break line in posts. Also improve parsing in general.
* Implement a modal for entry details via alt+enter
* Implement `tui_show_stub` option to show the stub to copy instead (or additionally) of copying it to the clipboard. Can be useful if you run gtl without X and/or are using tools like tmux to copy texts. Idea is from @szczezuja.
* Implement entry selection via mouse left click and simplify highlights code.

### Smaller improvements:
* [TUI] Add optional limit to number of displayed entries: Allow to limit the number of entries displayed in the timeline. See `tui_max_entries` option in README.
* Use gemini://tinylogs.gmi.bacardi55.io/ as default subscripton when no subscription file is found.
* Update dependencies
* Remove dead code

### Bug fixes
* Improve gemtext formating for level 3 headers and lists.
* Fix #29: Timeline refresh was previous refresh
* Fix error for missing file: if a tinylog isn't available on a working capsule, status was wrong.
* Fix limit bug in cli mode
* Remove date format duplicate

### Known Issues:
* Text in modal is not nicely formatted (eg: Thread or reply stub). This is due to an issue in cview: https://code.rocketnine.space/tslocum/cview/issues/72

Please read the README that has been updated accordingly.

## 0.5.2

**Critical Bug Fixing release**

* Fix case sensitivity in avatar/author metadata
* Fix date parsing error

## 0.5.1

*Minor release*

* Improve main panel title by adding filtered tinylog name
* Remove modal message after editing tinylog if no post script is defined.
* Remove message post script if successful and add config to refresh feeds after editing entry
* Improve focus changes
* Improve gemtext formatting for responses.
* Harmonized terminology around tinylogs and improve panels title

## 0.5.0

**Major release**

* Add ability to edit tinylog and launch a post edit script

## 0.4.8

*Minor release*

* Add UTC offset for date format
* Resolves #21: [TUI] Add a refresh message while feeds are refreshed
* Resolves #16: [TUI] Add ability to filter out tinylogs

## 0.4.7

*Minor release*

**Breaking change**: emoji status is now an option and disabled by default.
To enable it tui_status_emoji = true in gtl.toml

* Add sidebar toggle command in help box
* Improve title readability
* Add option for emoji usage. Also fixed some gemtext formatting
* fix gemtext formatting

## 0.4.6

*Minor release*

* Reformat TUI Sub list, resolves #15
* Resolves #17: [TUI] Don't quit gtl when displaying help with q
* change dead link emoji
* Add shortcut 's' to toggle sidebar display
* Add formating for links, lists and quotes
* Update sidebar color to match timeline color formatting
* Add help message in All sub secondary text
* Add new date format…

## 0.4.5

*Minor release*

**Breaking changes:**
* TUI shortcuts have been changed to single letter, no needs to ctrl anymore. Look at the README or display the help in TUI mode with ?.

* Fix #1: Create config dir at first launch and fix first launch
* Default mode is now TUI
* TUI - Implements #9: Show tinylogs with 0 compatible entries
* TUI - Resolves #12: Add last refresh in timeline. Hide footer for now before it is used again
* TUI - Resolves #13: Add color on border for focus panel
* TUI - Add display/hide help box and remove header
* TUI - shortcuts have changed, press ? for the help

## 0.4.3

*Minor release*

* fix focus
* Add TAB in usage
* fix regression in filtering
* fix code standards
* Fix separator regression

## 0.4.2

*Minor release*

New:
* Refresh subscriptions status when refreshing feeds (#6)
* Add last refresh time in footer
* Add indicator between new and old entries

Fixes:
* Makefile tcell dependency error
* Background color when terminal emulator background isn't black



## 0.4.1

*Minor release*

* Improve readme and add status emoji

## 0.4.0

**Major release**
First version of the TUI.

* TUI v1
* add emoji in case of feed errors
* Update README with cli limit configuration
* add new date format
* First try at TUI
* Add mode option
* Add colors and hightlights in timeline
