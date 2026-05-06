# Review Journal

This journal records the domain cases that matter before widening the public API.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 164, lane `ship`
- `stress`: `sync drift`, score 151, lane `ship`
- `edge`: `local state`, score 171, lane `ship`
- `recovery`: `conflict cost`, score 197, lane `ship`
- `stale`: `form pressure`, score 218, lane `ship`

## Note

The useful failure mode here is a wrong decision on a named case, not a vague style disagreement.
