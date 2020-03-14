# maul ![License: MIT][mit-badge] [![GoDoc][godoc-badge]][godoc] [![Go Report Card][go-report]][go-report-badge]

[mit-badge]: https://img.shields.io/badge/License-MIT-blue.svg
[godoc]: https://godoc.org/github.com/ymt2/maul
[godoc-badge]: https://godoc.org/github.com/ymt2/maul?status.svg
[go-report]: https://goreportcard.com/badge/github.com/ymt2/maul
[go-report-badge]: https://goreportcard.com/report/github.com/ymt2/maul

maul is a multiplexer to call [GitHub REST API v3: Create a milestone](https://developer.github.com/v3/issues/milestones/#create-a-milestone).

## At a glance

```sh
$ go get -u github.com/ymt2/maul/cmd/maul
$ export GITHUB_AUTH_TOKEN={YOUR_TOKEN}
$ maul --repo owner/repo1 --repo owner/repo2 --due-after 14 --title Sprint1
```
