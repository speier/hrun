const host = 'https://httpbin.org'

let data = { name: 'Foo Logo', cv: '@/tmp/foo.png' }
POST(`${host}/post`, 'Content-Type: multipart/form-data', 'Accept: multipart/form-data', data)
