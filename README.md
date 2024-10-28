![](./.github/banner.png)

<p align="center">
    SIDTool, A cross-platforms tool to work with Security Identifiers (SID) formats.
    <br>
    <a href="https://github.com/p0dalirius/SIDTool/actions/workflows/release.yaml" title="Build"><img alt="Build and Release" src="https://github.com/p0dalirius/SIDTool/actions/workflows/release.yaml/badge.svg"></a>
    <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/p0dalirius/SIDTool">
    <a href="https://twitter.com/intent/follow?screen_name=podalirius_" title="Follow"><img src="https://img.shields.io/twitter/follow/podalirius_?label=Podalirius&style=social"></a>
    <a href="https://www.youtube.com/c/Podalirius_?sub_confirmation=1" title="Subscribe"><img alt="YouTube Channel Subscribers" src="https://img.shields.io/youtube/channel/subscribers/UCF_x5O7CSfr82AfNVTKOv_A?style=social"></a>
    <br>
</p>

## Features

- [ ] Convert SID to and from formats:
   - [x] Parses SID in `hex`, `bytes` and `string` formats
   - [x] Outputs to `hex`, `bytes` and `string` formats.

## Usage

```              
$ ./SIDTool-linux-amd64 --help
SIDTool - by Remi GASCOU (Podalirius) - v1.1

Usage: SIDTool --value <string> [--to-hex] [--to-bytes] [--to-string]

  -v, --value <string> Input value to be converted.

  Conversion Options:
    -x, --to-hex    Convert the input value to hexadecimal. (default: false)
    -b, --to-bytes  Convert the input value to bytes. (default: false)
    -s, --to-string Convert the input value to string. (default: false)
```

## Contributing

Pull requests are welcome. Feel free to open an issue if you want to add other features.
