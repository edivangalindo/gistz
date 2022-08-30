# gistz
An easy way to dump gists.

This is inspired in [secretz](https://github.com/lc/secretz) and [jenkinz](https://github.com/lc/jenkinz) created by [lc](https://github.com/lc) but applicated to Github gists.

## Usage

Download all gists from one user:
```
echo nat | gistz
```

Download all gists from many users:
```
cat users.txt | gistz
```

## Installation

First, you'll need to [install go](https://golang.org/doc/install).

Then run this command to download + compile gistz:
```
go install github.com/edivangalindo/gistz@latest
```

You can now run `~/go/bin/gistz`. If you'd like to just run `gistz` without the full path, you'll need to `export PATH="/go/bin/:$PATH"`. You can also add this line to your `~/.bashrc` file if you'd like this to persist.
