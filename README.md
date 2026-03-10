# NewGin

`newgin` is a lightweight wrapper around the popular [Gin](https://github.com/gin-gonic/gin) web framework for Go. It provides a custom `Context` with built-in helper methods for standardized API responses, simplifying the development of RESTful APIs.

## Features

*   **Gin Wrapper**: Fully compatible with Gin's high-performance HTTP web framework.
*   **Custom Context**: Extends `gin.Context` to provide additional functionality.
*   **Standardized Responses**: Built-in `Success` and `Error` methods for consistent JSON API responses.
*   **Easy Migration**: API design mimics Gin's, making it easy to adopt if you're already familiar with Gin.

## Installation

```bash
go get github.com/lane-fullstack/newgin
```

## Usage

### Initialization

Initialize the engine just like you would with Gin:

```go
package main

import (
	"github.com/lane-fullstack/newgin"
)

func main() {
	// Create a new engine instance with default middleware (Logger and Recovery)
	r := newgin.Default()

	// Or create a blank engine
	// r := newgin.New()

    // ... define routes ...

	r.Run(":8080")
}
```

### Defining Routes

Use the custom `ctx.Context` in your handlers to access the helper methods:

```go
package main

import (
	"github.com/lane-fullstack/newgin"
	"github.com/lane-fullstack/newgin/ctx"
)

func main() {
	r := newgin.Default()

	r.GET("/ping", func(c *ctx.Context) {
		// Return a success response with data
		c.Success(map[string]string{
			"message": "pong",
		})
	})

    r.GET("/users", func(c *ctx.Context) {
        users := []string{"alice", "bob"}
        // Return a success response with list data and total count
        c.Success(users, 100)
    })

	r.GET("/error", func(c *ctx.Context) {
		// Return an error response
		c.Error("Something went wrong")
	})

	r.Run(":8080")
}
```

### Response Format

**Success Response:**

```json
{
    "code": 0,
    "message": "success",
    "data": { ... }
}
```

**Success Response (with list and total):**

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "list": [ ... ],
        "total": 100
    }
}
```

**Error Response:**

```json
{
    "code": 1,
    "msg": "Something went wrong"
}
```
