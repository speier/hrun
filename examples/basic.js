const host = 'https://httpbin.org'

// simple GET
GET(`${host}/get`)

// append header and POST data
POST(`${host}/post`, 'Accept: application/vnd.foobar.v2+json', { payload: { data: 'foobar' }})
