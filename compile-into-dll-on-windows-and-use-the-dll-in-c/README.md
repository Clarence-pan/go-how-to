Requirements
============

- GNU style `make` available in `%PATH%`
- `gcc` available in `%PATH%`
- `go` available in `%PATH%`
- `rm` available in `%PATH%` (for clean)

Keypoints
=========

Go cannot directly compile into DLL file on windows. But it can compile into static library file (`.a`).
So, we use `dll-bridge` to export the symbols to DLL. And then it can be used as other DLL.

Note
====

This requires cgo. But cgo cannot work with 32bit(x86) executables. So this way will only produce a 64bit DLL.
