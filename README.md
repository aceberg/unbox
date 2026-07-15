[![Main-Docker](https://github.com/aceberg/unbox/actions/workflows/main-docker-all.yml/badge.svg)](https://github.com/aceberg/unbox/actions/workflows/main-docker-all.yml)
[![Binary-release](https://github.com/aceberg/unbox/actions/workflows/binary-release.yml/badge.svg)](https://github.com/aceberg/unbox/actions/workflows/binary-release.yml)
[![Binary-Android](https://github.com/aceberg/unbox/actions/workflows/binary-android.yml/badge.svg)](https://github.com/aceberg/unbox/actions/workflows/binary-android.yml)

<h1><a href="https://github.com/aceberg/unbox">
    <img src="https://raw.githubusercontent.com/aceberg/unbox/main/assets/logo.png" width="30" />
</a>unbox</h1>
<br/>

Unbox is a CLI tool for [sing-box](https://github.com/SagerNet/sing-box) that has commands:   
1. `examples`: Show usage examples
2. `conf`: Works with sing-box config file to
    - [Deduplicate](https://github.com/aceberg/unbox#deduplicate) (compares all fields but `tag`)
    - [Convert sing-box config outbounds to URLs](https://github.com/aceberg/unbox#convert-sing-box-config-outbounds-to-urls)
    - [Remove unreachable nodes from sing-box config file](https://github.com/aceberg/unbox#remove-unreachable-nodes-from-sing-box-config)
3. `keep`: [Keep connection alive, auto switch proxy](https://github.com/aceberg/unbox#keep-connection-alive-and-switch-if-not)
4. `parse`: [Convert file with URLs to sing-box config](https://github.com/aceberg/unbox#convert-file-with-urls-to-sing-box-config)

Supported URLs: `vless://`,`hysteria2://` or `hy2://`,`trojan://`   
Experimental: `anytls://`, `ss://`, `tuic://`



## Screenshot

<details>
  <summary>Expand</summary>

![Screenshot_1](https://raw.githubusercontent.com/aceberg/unbox/main/assets/Screenshot_1.png)   
</details> 

## Install
<details><summary>Expand</summary>

### Docker
Example `keep` command with sing-box Clash API:
```sh
docker run aceberg/unbox keep -a "http://192.168.0.11:9090"
```
Or use `ghcr.io/aceberg/unbox`

### Binary
All available binaries are listed in the [latest](https://github.com/aceberg/unbox/releases/latest) release.    
For `amd64` there is an `apt` [repo](https://github.com/aceberg/ppa).

### Android and Termux
For `arm64` there are `android` and `termux.deb` [files](https://github.com/aceberg/unbox/releases/latest).
</details>

## Deduplicate
<details><summary>Expand</summary>

```sh
unbox conf -d -o sing-box.json
```

| Key | Description | Default |
| --- | ----------- | ------- |
| -d | Deduplicate |  |
| -o | Path to sing-box config file |  |

</details>

## Convert sing-box config outbounds to URLs
<details><summary>Expand</summary>

```sh
unbox conf -i sing-box.json > URLs.txt
```

| Key | Description | Default |
| --- | ----------- | ------- |
| -i | Path to sing-box config file |  |

</details>

## Remove unreachable nodes from sing-box config
<details><summary>Expand</summary>

```sh
unbox conf -a "http://127.0.0.1:9090" -o sing-box.json
```

| Key | Description | Default |
| --- | ----------- | ------- |
| -a | URL of sing-box Clash API |  |
| -as | Clash API secret |  |
| -l | Timeout for proxy delay (latency) check (ms) | 3000 |
| -o | Path to sing-box config file |  |
| -u | URL to test proxies. `https://www.gstatic.com/generate_204` will be used if empty |  |

</details>

## Keep connection alive and switch if not
<details><summary>Expand</summary>

```sh
unbox keep -a "http://127.0.0.1:9090"
```

| Key | Description | Default |
| --- | ----------- | ------- |
| -a | URL of sing-box Clash API |  |
| -as | Clash API secret |  |
| -da | Delay between checks of all proxy servers (seconds). Use 0 to disable | 300 |
| -db | Delay between checks of 3-4 backup servers (seconds). Use 0 to disable | 30 |
| -dm | Delay between checks of the main server (seconds). Use 0 to disable | 5 |
| -ds | Delay between auto switch to a faster proxy attempts (seconds). Use 0 to disable | 300 |
| -l | Timeout for proxy delay (latency) check (ms) | 3000 |
| -u | URL to test proxies. `https://www.gstatic.com/generate_204` will be used if empty |  |


</details>

## Convert file with URLs to sing-box config
<details><summary>Expand</summary>

Here `VLESS.txt` is a file with `vless://`,`hysteria2://`,`trojan://` links. Unbox will ignore anything else in the file, including other protocols and comments.
```sh
unbox parse -f VLESS.txt
```
In this example `tmpl.json` is a [template](https://github.com/aceberg/unbox/blob/main/configs/sing-box.tmpl.json) sing-box config and `sing-box.json` is where unbox will put generated config.
```sh
unbox parse -f VLESS.txt -t tmpl.json -o sing-box.json -j
```

| Key | Description | Default |
| --- | ----------- | ------- |
| -f | Path to file with URLs | VLESS.txt |
| -j | Validate and Indent JSON output |  |
| -n | Rename tags. If used, will rename tags to `tag1`, `tag2`... | |
| -o | Path to output sing-box config file |  |
| -t | Path to template sing-box config. Example [here](https://github.com/aceberg/unbox/blob/main/configs/sing-box.tmpl.json). There are only two variables available in template: `{{ .Unbox_tags }}` and `{{ .Unbox_outbounds }}` |  |
</details>

## Thanks

- [SagerNet/sing-box](https://github.com/SagerNet/sing-box)
- <a href="https://www.flaticon.com/free-icons/unboxing" title="Unboxing icons">Unboxing icons created by Futuer - Flaticon</a>