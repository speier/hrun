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

`-e`: set env var, key=value or filename (repeatable)

For example:

```sh
$ hrun -e=1 -e=2 -e bar.env foo.js
```

Flag `-e` is repeatable and can be key=value or a filename to load env vars, will be available on `env` object from the scripts.

## License

[MIT](LICENSE)
