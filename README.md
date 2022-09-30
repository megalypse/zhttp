# zhttp

It's a simple library that have as its life's purpose making YOUR life easier when doing HTTP requests.

## Usage
### GET Requests
```go
import (
	"fmt"

	"github.com/megalypse/zhttp"
	"github.com/megalypse/zhttp/zmodels"
)

func main() {
	response := zhttp.Get[StarWarsPerson](zmodels.ZRequest[zmodels.Void]{
		Url: "https://swapi.dev/api/people/1",
	})

	fmt.Printf("%+v\n", response.Content)
}

type StarWarsPerson struct {
	Name      string
	Height    string
	Mass      string
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
}
```
Every function from `zhttp` package takes a type argument, type that represents the response body you expect to get from the request.
If you're not expecting to get a response, just assign `zmodels.Void` to it, or anything else you want, the same is valid if you don't
need to send a body, just like we did in this example.

### POST Requests
```go
import (
	"fmt"

	"github.com/megalypse/zhttp"
	"github.com/megalypse/zhttp/zmodels"
)

func main() {
	response := zhttp.Post[Output](zmodels.ZRequest[Input]{
		Url: "https://jsonplaceholder.typicode.com/posts",
		Body: Input{
			Title:  "Test post",
			Body:   "This is the body for the post",
			UserId: 777,
		},
	})

	fmt.Printf("%+v\n", response.Content)
}

type Input struct {
	Title  string
	Body   string
	UserId int
}

type Output struct {
	Id     int
	Title  string `json:"title"`
	Body   string
	UserId int
}
```
If you want to send a body through the request, just create the type, instantiate the `ZRequest` type parameter with the body type, 
and give the `Body` field from `ZRequest` a value from this same type. In the example above we used the `Input` struct as our data model.

The rules you learned until now applies to every other request function from `zhttp` package.

### URL params
`zhttp` supports URL params replacement, both through `{}` and `:`, as the example above demonstrates:
```go
zhttp.Get[zmodels.Void](zmodels.ZRequest[zmodels.Void]{
		Url: "https://mock-url.com/people/{personId}/:addressId",
		UrlParams: map[string]string{
			"personId": "1",
			"addressId": "888",
		},
})
```
The final URL in this example before the request is made would be: `https://mock-url.com/people/1/888`

### Query Params
```go
zhttp.Get[zmodels.Void](zmodels.ZRequest[zmodels.Void]{
		Url: "https://mock-url.com/people",
		QueryParams: map[string][]string{
			"allowedStates": {"MG", "RJ", "PR"},
			"sort": {"YES"},
		},
})
```
The final URL generated from this example would be: `https://mock-url.com/people?allowedStates%3DMGallowedStates%3DRJallowedStates%3DPR%26sort%3DYES`(URL encoded)

### ZClient
If you want the possibility to have default headers and a default host URL for every request, `ZClient` provides it. Let's remake our first example
with the Star Wars API, but using a client this time.
```go
import (
	"fmt"

	"github.com/megalypse/zhttp/zclient"
	"github.com/megalypse/zhttp/zmodels"
)

func main() {
	client := zmodels.ZClient{
		ContextUrl: "https://swapi.dev/api/",
	}

	response := zclient.Get[StarWarsPerson](&client, zmodels.ZRequest[zmodels.Void]{
		Url: "/people/1",
	})

	fmt.Printf("%+v\n", response.Content)
}
```
By using `ZClient` you dismiss the need of having to provide the host URL for every request, and all that's left as a requirement is the endpoint
URI.
