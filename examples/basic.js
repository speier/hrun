const host = 'https://httpbin.org'

// simple GET
GET(`${host}/get?foo=bar`)

// append header and POST data
POST(`${host}/post`, 'Accept: application/vnd.xxx.v2+json', { foo: { bar: 'baz' }})

// expect status code
expect(200, GET(`${host}/status/400`).statusCode)
