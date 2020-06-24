# Ping microservice

A small microservice to ping a given url and return the status code.

## Usage - JS

```js
var formdata = new FormData();
formdata.append("url", "https://google.com.au");

var requestOptions = {
  method: 'POST',
  body: formdata
};

fetch("https://next-ping.vercel.app/api/ping", requestOptions)
  .then(response => response.json())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

```json
{
  "code": "200",
  "url": "https://google.com.au"
}
```
