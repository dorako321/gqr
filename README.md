# gqr
`gqr` shows QR code what you made in your terminal.

[![Build Status](https://travis-ci.com/dorako321/gqr.svg?branch=master)](https://travis-ci.com/dorako321/gqr)

![intro](https://github.com/dorako321/gqr/raw/master/doc/intro.gif)


# Install

```
$ go get -u github.com/dorako321/gqr
```

# Usage

`-m` option can set the message.

```
$ gqr -m foo
```

Also, you can set the message to use pipe command.
```
$ echo "bar" | gqr
```

## License
MIT