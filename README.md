ltsv-cli.go
========================================

LTSV Command Line Tool

Requirement
----------------------------------------

### Go 1.7+

Available cross compile

on mac

```
brew install go --cross-compile-all
```

on linux, see: http://dave.cheney.net/2013/07/09/an-introduction-to-cross-compilation-with-go-1-1


### Gom

https://github.com/mattn/gom



Development
----------------------------------------

Install libraries.

```
gom install
```

After coding, build.

```
make generate
```

Generate sample ltsv (./sample.ltsv).

```
make sample
```

Format source codes

```
make format
```



Install
----------------------------------------

ex

on `linux x86_64`:

```
curl https://raw.githubusercontent.com/sugilog/ltsv-cli.go/master/gen/lc.linux.amd64 > ~/bin/lc && chmod 0755 ~/bin/lc
```

on `Mac x86_64`:

```
curl https://raw.githubusercontent.com/sugilog/ltsv-cli.go/master/gen/lc.darwin.amd64 > ~/bin/lc && chmod 0755 ~/bin/lc
```

Usage
----------------------------------------

Filter ltsv keys

```
cat {ltsv file} | lc f --key={key1,key2,..}
```

Grep ltsv lines

```
cat {ltsv file} | lc f --key={key1,key2,..} {word}
```


