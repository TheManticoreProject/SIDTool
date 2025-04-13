![](./.github/banner.png)

<p align="center">
    A cross-platform tool to work with Security Identifiers (SID) formats.
    <br>
    <a href="https://github.com/TheManticoreProject/SIDTool/actions/workflows/release.yaml" title="Build"><img alt="Build and Release" src="https://github.com/TheManticoreProject/SIDTool/actions/workflows/release.yaml/badge.svg"></a>
    <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/TheManticoreProject/SIDTool">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/TheManticoreProject/SIDTool"> 
    <br>
</p>

## Features

- [ ] Convert SID to and from formats:
   - [x] Parses SID in `hex`, `bytes` and `string` formats
   - [x] Outputs to `hex`, `bytes` and `string` formats.

## Usage

```              
$ ./SIDTool-linux-amd64 --help
SIDTool - by Remi GASCOU (Podalirius) @ TheManticoreProject - v1.2

Usage: SIDTool --value <string> [--to-hex] [--to-bytes] [--to-string]

  -v, --value <string> Input value to be converted.

  Conversion Options:
    -x, --to-hex    Convert the input value to hexadecimal. (default: false)
    -b, --to-bytes  Convert the input value to bytes. (default: false)
    -s, --to-string Convert the input value to string. (default: false)
```

## Contributing

Pull requests are welcome. Feel free to open an issue if you want to add other features.

## Credits
  - [@p0dalirius](https://github.com/p0dalirius) for the creation of the [SIDTool](https://github.com/p0dalirius/SIDTool) project before transferring it to TheManticoreProject.