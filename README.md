# ducksay-go

*Because the world desperately needed a Go port of a Rust tool that generates a duck*.

`ducksay-go` is a direct, rewrite-it-in-Go port of [sboult's `ducksay`](https://github.com/sboult/ducksay) (which was written in Rust). 

### Why Go? 
Because I hate Rust. *(sorry)*

## What is this?

Have you ever looked at your HTML, Markdown, or source code and thought, *"This is great, but it lacks a duck named Waddles speaking to me in the comments"*? 

Look no further. `ducksay` generates a beautiful ASCII representation of Waddles wrapped in HTML comments, allowing you to embed secret avian wisdom directly into your codebase.

### Example

```html
<!--       _
        .__(.)< (I am Waddles)
         \___)
 ~~~~~~~~~~~~~~~~~~-->
```

Or, if you prefer Twitter's layout standards (where Waddles has a slightly larger eye for some reason, come on Elon...):

```html
<!--      _
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

Or just get the binary from Releases if you don't want to install Go on your machine.

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

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
