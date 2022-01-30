# Fish to zsh history converter

This small program is designed to convert the history from fish shell to zsh shell when you migrate between shells.

## How to use

1. Download last version from releases
2. Unpack binary file from archive and execute it
```
$ tar -xvzf fish_to_zsh_history-*.tar.gz && \
  ./fish_to_zsh_history
```
3. After conversion finished you need replace you'r current zsh history file:
```
$ cp ~/.zsh_history ~/.zsh_history.backup && \
  mv ~/.zsh_history.convert ~/.zsh_history
```