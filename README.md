# Toolkit

This is the course [`Building a module in Go`](https://www.udemy.com/course/building-a-module-in-go-golang/learn/lecture/32938420#overview)

## How to init this project

### Project structure

```
.
+-- app
|   +-- go.mod
+-- toolkit
|   +-- go.mod
+-- go.work
```

#### app

Our application which leverages tools from toolkit

#### toolkit

Tools that we want to publish and reuse

The included tools are:

- [ ] Read JSON
- [ ] Write JSON
- [ ] Produce a JSON encoded error response
- [X] Upload a file to a specified directory
- [ ] Download a static file
- [X] Get a random string of length n
- [ ] Post JSON to a remote service 
- [X] Create a directory, including all parent directories, if it does not already exist
- [X] Create a URL safe slug from a string

Installation

`go get -u github.com/tsawler/toolbox`

#### app.work

The toolkit package won't be added in to app/go.mod due to this workspace

```bash
# how to generate this go.work
$ go work init toolkit app
```