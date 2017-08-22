check-cert-expire
=================

Installation
------------

### Binary
Download binary from https://github.com/tetsun/check-cert-expire/releases.

### Go get
```zsh
go get github.com/tetsun/check-cert-expire
```

Usage
-----
See help.
```zsh
check-cert-expire -h
```

Examples
--------
Basic
```zsh
$ check-cert-expire google.com
2017-10-31 12:39:00 +0000 UTC
# => 0
```

Check delay
```zsh
$ check-cert-expire -d 90 google.com
2017-10-31 12:39:00 +0000 UTC
# => 65
```

Silent mode (for script)
```zsh
$ check-cert-expire -s google.com
# => 0
```