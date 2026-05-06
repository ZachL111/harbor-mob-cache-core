# Harbor Mob Cache Core Walkthrough

I use this file as a small checklist before changing the Go implementation.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 164 | ship |
| stress | sync drift | 151 | ship |
| edge | local state | 171 | ship |
| recovery | conflict cost | 197 | ship |
| stale | form pressure | 218 | ship |

Start with `stale` and `stress`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

`stale` is the optimistic case; use it to make sure the scoring path still rewards strong signal.
