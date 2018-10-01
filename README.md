# JSON Utility (jut)

A simple tool for pretty printing JSON

## Installation

The package's API is not fleshed out, so until further notice, the only value is in installing the command ```jut``` with:

```go get github.com/the-rileyj/jsonutility/jut```

## Usage

After installation, assuming your GOPATH is setup, you should be able to use ```jut```:

```bash
$ echo "{\"wow\":[1,2,3]}" | jut
```

Outputs:

```bash
{
  "wow": [
    1,
    2,
    3
  ]
}
```

Adjust output spacing:

```bash
$ echo "{\"wow\":[1,2,3]}" | jut -s 4
{
    "wow": [
        1,
        2,
        3
    ]
}
$ 
```

More to come...