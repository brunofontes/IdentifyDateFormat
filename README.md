# IdentifyDateFormat

[![Go Reference](https://pkg.go.dev/badge/github.com/brunofontes/IdentifyDateFormat.svg)](https://pkg.go.dev/github.com/brunofontes/IdentifyDateFormat)

A very simple way of checking the possible date formats that matches the dates you have. Useful when you need to open files with lots of dates, but you don't know which format was used to export them.

## How it works?

It just tries to parse the dates against a custom slice of date formats. By default, it already has some formats that are commonly used, but you can change that. If the format were not parseable, it will be removed from the list of possible dates.
