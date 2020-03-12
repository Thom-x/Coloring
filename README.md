Add prefix suffix and color from file or pipe to console

### Build
```bash
go build -o ./bin/decoration
```

### Example
```bash
> cat words | sort | uniq | decoration -p "prefix " -s " suffix" -c "red"
APPLE
BYE
HELLO
ZEBRA
```

### Command line options
```bash
> ./bin/decoration --help
Usage:
  decoration [flags]

Flags:
  -c, --color string    color : black, red, green, yellow, blue, magenta, cyan, white
  -f, --file string     path to the file
  -h, --help            help for decoration
  -p, --prefix string   prefix
  -s, --suffix string    suffix
  -v, --verbose         log verbose output
  
  
> echo "hello from the shell" | ./bin/decoration
  hello from the shell
  
> ./bin/decoration -f /tmp/test 
  hello from file
```