# ducksay-go

*Because the world desperately needed a Go port of a Rust tool that generates a duck*.

`ducksay-go` is a port of [sboult's `ducksay`](https://github.com/sboult/ducksay) (which was written in Rust), but I also added some more stuff that isn't in the original repo.

### Why Go? 
Because I hate Rust. *(sorry)*

## What is this?

Have you ever looked at your HTML, Markdown, or source code and thought, *"This is great, but it lacks a duck named Waddles speaking to me in the comments"*? 

Look no further. `ducksay` generates a beautiful ASCII representation of Waddles wrapped in HTML comments, allowing you to embed secret avian wisdom directly into your codebase.

### Example

```html
<!--        _
        .__(.)< (I am Waddles)
         \___)
 ~~~~~~~~~~~~~~~~~~-->
```

Or, if you prefer Twitter's layout standards (where Waddles has a slightly larger eye for some reason, come on Elon...):

```html
<!--         _
        .__( . )< (Follow me for more hot takes)
         \___)
 ~~~~~~~~~~~~~~~~~~-->
```

## Installation

Assuming you have Go installed, you can build it from source:

```bash
go build -o ducksay ducksay.go
```

Or install it directly to your `$GOPATH/bin`:

```bash
go install
```

Or just [get the binary from Releases](https://github.com/dmx3377/ducksay-go/releases) if you don't want to install Go on your machine.

## Usage

Make Waddles say the default message:
```bash
./ducksay
```

Make Waddles say whatever you want:
```bash
./ducksay "Quack is the new track."
```

Wrap the text at a specific column width (default is 40):
```bash
./ducksay -width 20 "This is a very long message that Waddles will politely wrap for you so it does not overflow your HTML comments."
```

Use the Twitter/X-compatible style (larger eye space):
```bash
./ducksay -twitter "I am tweeting!"
```

### Comment Styles

Waddles can be wrapped in comments for various programming languages using the `-style` flag:

```bash
# Line comments (Go, JS, C++, Rust, etc.)
./ducksay -style go "Waddles in Go!"

# Hash comments (Python, Ruby, Bash, YAML, etc.)
./ducksay -style py "Waddles in Python!"

# SQL line comments (SQL, Lua, Haskell)
./ducksay -style sql "Waddles in SQL!"

# No comments (raw terminal mode)
./ducksay -style none "Raw Waddles!"
```

### Color Styling

Add premium colors using `-color-duck` and `-color-bubble` flags:

```bash
# Color Waddles yellow and the bubble cyan
./ducksay -color-duck yellow -color-bubble cyan "Colorful Waddles!"

# Use green and standard blue
./ducksay -style go -color-duck green -color-bubble blue "Developer Waddles!"

# Use custom hex codes
./ducksay -color-duck #FF4500 -color-bubble #FFD700 "Hex custom colours!"
```

Supported colours include standard names (`red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`), `devgreen`, or any custom hex color code in `#RRGGBB` format.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Disclaimer
[Waddles is an AWS/Amazon thing](https://x.com/awsdevelopers/status/2032135107977322989). **I don't own Waddles.** This is just a project I made for fun.

