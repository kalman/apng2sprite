# apng2sprite

Simple utility that converts apng (aka animated png files) to (horizontal) sprite sheets.
Useful for Unity and probably other things.

```
./apng2sprite -i myAnimatedPng.png -o mySpriteSheet.png
```

If you're building from source, requires:
 1. [Go](https://golang.org)
 2. [kettek/apng](https://github.com/kettek/apng) (`go get github.com/kettek/apng`)
