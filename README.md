# unbox

Converts a list of `vless://` links to a sing-box config file.

## Quick start
Here `VLESS.txt` is a file with `vless://` links. Unbox will ignore anything else in the file, including other protocols and comments.
```sh
unbox -f VLESS.txt
```
In this example `sing-box.tmpl.json` is a [template](configs/sing-box.tmpl.json) sing-box config and `sing-box.json` is where unbox will put generated config.
```sh
unbox -f VLESS.txt -t sing-box.tmpl.json -o sing-box.json
```

## Options

| Key | Description | Default |
| --- | ----------- | ------- |
| -f | Path to file with `vless://` links | VLESS.txt |
| -n | Rename tags (yes/no). If `yes`, with rename tags to `tag1`, `tag2`... | no |
| -o | Path to output file |  |
| -t | Path to template sing-box config. Example [here](configs/sing-box.tmpl.json). There are only two variables available in template: `{{ .Unbox_tags }}` and `{{ .Unbox_outbounds }}` |  |

## Install