[![Main-Docker](https://github.com/aceberg/unbox/actions/workflows/main-docker-all.yml/badge.svg)](https://github.com/aceberg/unbox/actions/workflows/main-docker-all.yml)
[![Binary-release](https://github.com/aceberg/unbox/actions/workflows/binary-release.yml/badge.svg)](https://github.com/aceberg/unbox/actions/workflows/binary-release.yml)
[![Binary-Android](https://github.com/aceberg/unbox/actions/workflows/binary-android.yml/badge.svg)](https://github.com/aceberg/unbox/actions/workflows/binary-android.yml)

<h1><a href="https://github.com/aceberg/unbox">
    <img src="https://raw.githubusercontent.com/aceberg/unbox/main/assets/logo.png" width="20" />
</a>unbox</h1>
<br/>

Unbox is a CLI tool for [sing-box](https://github.com/SagerNet/sing-box) that can do 3 things:
1. Convert a file with `vless://`,`hysteria2://`,`trojan://` links to full sing-box config
2. Remove unreachable nodes from sing-box config file (uses sing-box Clash API)
3. Keep connection alive and switch to another proxy immediately if not (uses sing-box Clash API)

## Screenshot

<details>
  <summary>Expand</summary>

![Screenshot_1](https://raw.githubusercontent.com/aceberg/unbox/main/assets/Screenshot_1.png)   
</details> 

## Install
<details><summary>Expand</summary>

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
</details>

## Convert file with URLs to sing-box config
<details><summary>Expand</summary>

Here `VLESS.txt` is a file with `vless://`,`hysteria2://`,`trojan://` links. Unbox will ignore anything else in the file, including other protocols and comments.
```sh
unbox -f VLESS.txt
```
In this example `sing-box.tmpl.json` is a [template](https://github.com/aceberg/unbox/blob/main/configs/sing-box.tmpl.json) sing-box config and `sing-box.json` is where unbox will put generated config.
```sh
unbox -f VLESS.txt -t sing-box.tmpl.json -o sing-box.json -j
```
Docker
```sh
docker run -it -v /your/local/path:/data \
    aceberg/unbox \
    -f /data/VLESS.txt \
    -t /data/sing-box.tmpl.json \
    -o /data/sing-box.json
```

| Key | Description | Default |
| --- | ----------- | ------- |
| -f | Path to file with links | VLESS.txt |
| -j | Validate and Indent JSON output |  |
| -n | Rename tags. If used, will rename tags to `tag1`, `tag2`... | |
| -o | Path to output file |  |
| -t | Path to template sing-box config. Example [here](https://github.com/aceberg/unbox/blob/main/configs/sing-box.tmpl.json). There are only two variables available in template: `{{ .Unbox_tags }}` and `{{ .Unbox_outbounds }}` |  |
</details>

## Remove unreachable nodes from sing-box config
<details><summary>Expand</summary>

```sh
unbox -a "http://127.0.0.1:9090" -o sing-box.json
```

| Key | Description | Default |
| --- | ----------- | ------- |
| -a | URL to sing-box Clash API |  |
| -o | Path to sing-box config file |  |

</details>

## Keep connection alive and switch if not
<details><summary>Expand</summary>

Use the `-k` flag to run `unbox` in keepalive mode
```sh
unbox -a "http://127.0.0.1:9090" -k
```

| Key | Description | Default |
| --- | ----------- | ------- |
| -a | URL to sing-box Clash API |  |
| -da | Delay between checks of all proxy servers (seconds). Use 0 to disable | 300 |
| -db | Delay between checks of 3-4 backup servers (seconds). Use 0 to disable | 30 |
| -dm | Delay between checks of the main server (seconds). Use 0 to disable | 5 |
| -k | Keepalive mode | |
| -u | URL to test proxies. `https://www.gstatic.com/generate_204` will be used if empty |  |


</details>

## Thanks

- [SagerNet/sing-box](https://github.com/SagerNet/sing-box)
- <a href="https://www.flaticon.com/free-icons/unboxing" title="Unboxing icons">Unboxing icons created by Futuer - Flaticon</a>