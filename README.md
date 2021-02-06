# `merge-mgr`

_merge manager_

TUI wrapper around
[`gh`](https://github.com/cli/cli/)
to merge pull requests,
with a similar workflow to [`github.com/pulls`](github.com/pulls)

## Installation

`go get -u github.com/tlwr/merge-mgr`

## Usage

Ensure that [`gh`](https://github.com/cli/cli/) is installed and you are logged in

```
$ merge-mgr
2021/02/06 21:16:05 getting pull requests for user "tlwr"
choose a PR to merge

> ( ) tlwr/some-gallery-thing [595] https://github.com/tlwr/some-gallery-thing/pull/595 build(deps): bump webpack from 5.20.1 to 5.21.1
  (•) tlwr/some-gallery-thing [593] https://github.com/tlwr/some-gallery-thing/pull/593 build(deps): bump sass-loader from 10.1.1 to 11.0.0
  (•) tlwr/some-gallery-thing [592] https://github.com/tlwr/some-gallery-thing/pull/592 build(deps): bump ts-loader from 8.0.14 to 8.0.15
  ( ) tlwr/some-gallery-thing [586] https://github.com/tlwr/some-gallery-thing/pull/586 build(deps): bump @types/react from 17.0.0 to 17.0.1
  ( ) tlwr/some-gallery-thing [585] https://github.com/tlwr/some-gallery-thing/pull/585 build(deps): bump webpack-cli from 4.4.0 to 4.5.0
  ( ) tlwr/operator-tools [9] https://github.com/tlwr/operator-tools/pull/9 build(deps): bump github.com/onsi/gomega from 1.10.4 to 1.10.5
  ( ) tlwr/some-gallery-thing [583] https://github.com/tlwr/some-gallery-thing/pull/583 build(deps): bump sass from 1.32.5 to 1.32.6
  ( ) tlwr/some-gallery-thing [582] https://github.com/tlwr/some-gallery-thing/pull/582 build(deps): bump pino-pretty from 4.4.0 to 4.5.0

(press space to select ; enter to proceed ; q to quit)

2021/02/06 21:16:09 merging https://github.com/tlwr/some-gallery-thing/pull/593
✔ Merged pull request #593 (build(deps): bump sass-loader from 10.1.1 to 11.0.0)
2021/02/06 21:16:11 merging https://github.com/tlwr/some-gallery-thing/pull/592
✔ Merged pull request #592 (build(deps): bump ts-loader from 8.0.14 to 8.0.15)
```
