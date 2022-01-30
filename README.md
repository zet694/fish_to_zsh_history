# Fish to zsh history converter

This small program is designed to convert the history from fish shell to zsh shell when you migrate between shells.

## How to use

1. Download the last version from [releases](https://github.com/zet694/fish_to_zsh_history/releases)
2. Unpack binary file from archive and execute it
```
$ tar -xvzf fish_to_zsh_history-*.tar.gz && \
  ./fish_to_zsh_history
```
3. After the conversion is finished you need replace your current zsh history file:
```
$ cp ~/.zsh_history ~/.zsh_history.backup && \
  mv ~/.zsh_history.convert ~/.zsh_history
```
4. Cleanup after work:
```
$ rm fish_to_zsh_history fish_to_zsh_history-v*
```