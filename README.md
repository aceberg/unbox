[![Main-Docker](https://github.com/aceberg/unbox/actions/workflows/main-docker-all.yml/badge.svg)](https://github.com/aceberg/unbox/actions/workflows/main-docker-all.yml)
[![Binary-release](https://github.com/aceberg/unbox/actions/workflows/binary-release.yml/badge.svg)](https://github.com/aceberg/unbox/actions/workflows/binary-release.yml)
[![Binary-Android](https://github.com/aceberg/unbox/actions/workflows/binary-android.yml/badge.svg)](https://github.com/aceberg/unbox/actions/workflows/binary-android.yml)

# unbox

Converts a list of `vless://` links to a sing-box config file.

## Quick start
Here `VLESS.txt` is a file with `vless://` links. Unbox will ignore anything else in the file, including other protocols and comments.
```sh
unbox -f VLESS.txt
```
In this example `sing-box.tmpl.json` is a [template](https://github.com/aceberg/unbox/blob/main/configs/sing-box.tmpl.json) sing-box config and `sing-box.json` is where unbox will put generated config.
```sh
unbox -f VLESS.txt -t sing-box.tmpl.json -o sing-box.json
```
### Docker
```sh
docker run -it -v /your/local/path:/data \
    aceberg/unbox \
     -f /data/VLESS.txt \
     -t /data/sing-box.tmpl.json \
     -o /data/sing-box.json
```

## Options

| Key | Description | Default |
| --- | ----------- | ------- |
| -f | Path to file with `vless://` links | VLESS.txt |
| -j | Validate and Indent json output (yes/no) | no |
| -n | Rename tags (yes/no). If `yes`, will rename tags to `tag1`, `tag2`... | no |
| -o | Path to output file |  |
| -t | Path to template sing-box config. Example [here](https://github.com/aceberg/unbox/blob/main/configs/sing-box.tmpl.json). There are only two variables available in template: `{{ .Unbox_tags }}` and `{{ .Unbox_outbounds }}` |  |

## Install
### Docker
There are DockerHub and GitHub images:
```sh
docker pull aceberg/unbox
```
```sh
docker pull ghcr.io/aceberg/unbox
```

### Binary
All available binaries are listed in the [latest](https://github.com/aceberg/unbox/releases/latest) release.    
For `amd64` there is an `apt` [repo](https://github.com/aceberg/ppa).

### Android and Termux
For `arm64` there are `android` and `termux.deb` [files](https://github.com/aceberg/unbox/releases/latest).