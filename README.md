# gistz
An easy way to dump gists.

This is inspired in [secretz](https://github.com/lc/secretz) and [jenkinz](https://github.com/lc/jenkinz) created by [lc](https://github.com/lc) but applicated to Github gists.

Pre-requisites:

* You need to config an env called ```GH_AUTH_TOKEN``` with your personal access token, to do the requests

## Usage

Download all gists from one user:
```
echo nat | gistz
```

Download all gists from many users:
```
cat users.txt | gistz
```

This works fine with [gh-members](https://github.com/edivangalindo/gh-members) to get gists from members of a organization:
```
echo github | gh-members | gistz
```

## Installation

First, you'll need to [install go](https://golang.org/doc/install).

Then run this command to download + compile gistz:
```
go install github.com/edivangalindo/gistz@latest
```

You can now run `~/go/bin/gistz`. If you'd like to just run `gistz` without the full path, you'll need to `export PATH="/go/bin/:$PATH"`. You can also add this line to your `~/.bashrc` file if you'd like this to persist.
