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
2021/01/24 12:23:31 getting pull requests for user "tlwr"
choose a PR to merge

( ) tlwr/some-gallery-thing [567] https://github.com/tlwr/some-gallery-thing/pull/567 build(deps): bump sass from 1.32.4 to 1.32.5
(•) tlwr/some-gallery-thing [561] https://github.com/tlwr/some-gallery-thing/pull/561 build(deps): bump @types/node-fetch from 2.5.7 to 2.5.8
( ) tlwr/some-gallery-thing [523] https://github.com/tlwr/some-gallery-thing/pull/523 build(deps): bump cheerio from 1.0.0-rc.3 to 1.0.0-rc.5

(press q to quit)
2021/01/24 12:23:36 merging https://github.com/tlwr/some-gallery-thing/pull/561
? What merge method would you like to use? Create a merge commit
? Delete the branch locally and on GitHub? No
? Submit? Yes
✔ Merged pull request #561 (build(deps): bump @types/node-fetch from 2.5.7 to 2.5.8)
```

## Problems

Selecting "Yes" when asked to "Delete the branch locally" may fail,
this is behaviour of the `gh` tool
