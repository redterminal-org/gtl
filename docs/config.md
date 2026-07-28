# Configuration

## Default configuration file

```toml
# Default config file
#
# You have to edit this file to enable replying/editing
# in your tinylog

## General configuration:
# Path to subscribed tinylogs:
subscribed_data = "~/.config/gtl/subs"
# Date display in iternational format. Please don't change this.
date_format = "2006-01-02 15:04 -0700"
# Log file:
log_file = "/dev/null"
# Optional: Highlight when text is found in content.
# Separate values by a coma, eg:
# highlights = "@bacardi55, @bacardi, anything"
# Ignore Highlights from author:
# Must be the full author displayed name!
# ignore_author_highlights = "🤔 @bacardi55"
# Maximum number of entries showed in cli mode.
# If --limit is used, it will overide this setting.
# Will be ignored in tui mode.
cli_limit = 10
# Mode: either cli or tui
mode = "tui"
# If false, standard ascii characters will be used.
tui_status_emoji = false

## TUI options:
# Enable tinylog edition from gtl:
# This will use an external editor,
# configured in your EDITOR environment variable.
# You can check with 'echo $EDITOR' to see if it
# is configured correctly.
# ctrl+n is disabled when set to false.
allow_edit = false
# Path to tinylog file. This option is ignored if
# If not a valid file, editing will not be possible
# and ctrl+n will be disabled.
tinylog_path = "path/to/tinylog/file.gmi"
# Path to script to be executed after the edition is done.
# Should upload your tinylog path file to your Gemini Capsule.
# This script needs to be executable.
# If not a valid executable script, it will be ignored.
post_edit_script = "path/to/upload_script"
# Auto refresh feeds after editing the tinylog file.
# Only used when allow_edit = true
post_edit_refresh = false
# Limit the number of entries displayed in TUI.
# Indicate 0 for all entries.
# (This shouldn't be set to 0! GTL TUI would become VERY sluggish!)
# 100 entries are enough and reasonable.
tui_max_entries = 100
# Copy a pre formatted text to clipboard when creating a new entry
# On linux, requires 'xclip' or 'xsel'
tui_copy_stub_clipboard = false
# If you are running gtl without X, the copy to clipboard feature
# will not work (Or if you don't have xclip or xsel).
# In this case, enabling this option will allow gtl to display
# the sub text in a modal for easy copy in tools like tmux
# At this stage, the rendering is ugly because of an issue in cview:
# https://code.rocketnine.space/tslocum/cview/issues/72#issuecomment-3968
tui_show_stub = false

## TUI Theming:
# Define colors for GTL TUI mode.
# Default (dark) theme:
#tui_color_background = "000000"
#tui_color_box = "FFFFFF"
#tui_color_focus_box = "008000"
#tui_color_author_name = "FF0000"
#tui_color_links = "87CEEB"
#tui_color_elapsed_time = "87CEEB"
#tui_color_text = "FFFFFF"
#tui_color_selected_background = "FFFFFF"
#tui_color_selected_foreground = "000000"
#tui_color_highlight = "FF0000"
#tui_color_quote = "808080"
#tui_color_preformatted = "00A000"

# Bookmarks (only available via TUI for now):
bookmarks_enabled = false
bookmarks_file_path = "~/.config/gtl/bookmarks.json"
bookmarks_reverse_order = true
```

By default, gtl will look for ~/.config/gtl/gtl.toml as configuration file. It will create it if needed.

The --config option only looks for the file, it will not create it if the file given as argument of --config doesn't exist.

## Subscription management

You can add and remove tinylog entry either manually from the file directly, or use gtl to do it for you:
```
Subscription management usage:
	add --url url [--title title]	Indicate a new tinylog url (and optional title) to subscribe to.
	rm --url url			Indicate a tinylog url to be removed from the subscription.
```

