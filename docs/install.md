# Installation

## From Binary
Download a recent binary from the [Releases page](https://codeberg.org/fab/gtl/releases) for your OS and architecture, make it executable and put it in a directory in your $PATH (eg. `~/bin/gtl`).

After running for the first time, a configuration file `~/.config/gtl/gtl.toml`
is created which you need to edit, to enable editing/replying in your tinylog
and maybe make some other changes.

```bash
# Example binary for linux x86_64 (amd64) architecture in ~/Downloads/ directory
cd ~/Downloads/
# make binary executable
chmod +x gtl-linux-amd64
# move executable binary into your $PATH and rename it to 'gtl'
mv gtl-linux-amd64 ~/bin/gtl
```

## Updateing to a recent version
Updating GTL to a recent version should take no more than switching out the old
binary with the new released binary. Maybe you need to make it executable again.
If changes to the `gtl.toml` config file should be needed you should be informed
in the [README.md](../README.md) page about it.

## From Source
gtl requires go ≥ 1.26.4 since tag _v1.3.0-fork_

```bash
# clone GTL repository with git
git clone http://codeberg.org/fab/gtl.git
# cd into git repository
cd gtl
# build executable with Golang (Go)
go build
# move executable into $PATH
mv gtl ~/bin/
```

## Bookarks (since tag v1.1.2-fork)
If you're updating from an older release than v1.1.2-fork, you have to delete
your `bookmarks.json` file and let the gtl binary create a new one, because the
previous ones have created an invalid JSON file and therefore GTL will disable
bookmark functionality if using the old file.
