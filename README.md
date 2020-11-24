# hrun

HTTP request runner with a simple DSL based on JavaScript.

## Install

Download the [latest release](https://github.com/speier/hrun/releases)

or install using GoBinaries:

```sh
$ curl -sf https://gobinaries.com/speier/hrun | sh
```

## Usage example

Create a script:

```js
const host = 'https://httpbin.org'

GET(`${host}/get`)
POST(`${host}/post`, 'Accept: application/vnd.foobar.v2+json', { payload: { data: 'foobar' }})
```

Run the script:

```sh
$ hrun -f examples/basic.js
```

## API

All HTTP method is a function with an URL, headers and payload arguments:

```js
METHOD(URL, [headers], [payload])
```

`URL`: required

`[headers]`: optional

`[payload]`: optional

## CLI

The following flags are available:

`-f`: filename

`-s`: set param

For example:

```sh
$ hrun -f foo.js -s a=1 b=2
```

Filename `-f` is required, `-s` is optional to set key value params, these params alongside the host environment variables will be available for the script on the `env` object.

## License

[MIT](LICENSE)
