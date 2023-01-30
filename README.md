# toolkit

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

#### app.work

The toolkit package won't be added in to app/go.mod due to this workspace

```bash
# how to generate this app.work
$ go work init toolkit app
```