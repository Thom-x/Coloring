# Decoration

Execute a program and add prefix suffix and color to its log output

![Go](https://github.com/Thom-x/Coloring/workflows/Go/badge.svg)
![Release](https://github.com/Thom-x/Decoration/workflows/Release/badge.svg)

---

### Build
```bash
go build -o ./bin/decoration
```

### Example
```bash
> decoration -e myprogram -a "args args args" -p "[prefix] " -s "[suffix]" -c yellow
  [prefix] output [suffix]
```

### Command line options
```bash
> ./bin/decoration --help
Execute a program and add prefix suffix and color to its log output

Usage:
  decoration [flags]

Flags:
  -a, --args string      arguments of the program
  -c, --color string     color : black, red, green, yellow, blue, magenta, cyan, white
  -h, --help             help for decoration
  -p, --prefix string    prefix
  -e, --program string   program to execute
  -s, --suffix string    suffix
  
> decoration -e myprogram -a "args args args" -p "[prefix] " -s "[suffix]" -c yellow
  [prefix] output [suffix]
```
