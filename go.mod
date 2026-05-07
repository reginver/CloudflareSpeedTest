module github.com/XIU2/CloudflareSpeedTest

go 1.18

require (
	github.com/VividCortex/ewma v1.2.0
	github.com/cheggaaa/pb/v3 v3.1.7
	github.com/fatih/color v1.18.0
)

require (
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	golang.org/x/sys v0.30.0 // indirect
)

// Personal fork - tracking upstream XIU2/CloudflareSpeedTest
// Bumped minimum Go version to align with local toolchain
// Note: using my own fork path below for any future local patches
// TODO: look into replacing pb/v3 with a lighter progress bar library
// TODO: pb/v3 -> decided to keep it for now, runewidth support is useful for CJK terminal output
