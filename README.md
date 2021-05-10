# colorhelper

`colorhelper` provides additional golang utilities for colors.

## Features

* sRGB type compatible with the go `color.Color` and `color.Model` interfaces.
* HSLA type compatible with the go `color.Color` and `color.Model` interfaces.
* Function to parse legal CSS color values into `color.Color` values.
* Function to generate legal CSS color values from `color.Color` values.
* Function to choose the best text color for a given background color.
* Template function to tweak a CSS color value.
* Template function to to choose the best text color for a given background color.

## Examples

### Parsing CSS color values

```go
c1, err := colorhelper.ParseCSSColorRepresentation("#FF0000")
if err != nil {
    panic(err)
}

c2, err := colorhelper.ParseCSSColorRepresentation("hsla(0, 100%, 50%, 0.5)")
if err != nil {
    panic(err)
}

fmt.Printf("%[1]T%[1]v\n", c1) // color.RGBA{255, 0, 0, 255}
fmt.Printf("%[1]T%[1]v\n", c2) // colorhelper.HSLA{0, 1, 0.5, 0.5}
```

### Generating CSS Color Values

```go
c := color.RGBA{255, 0, 0, 127}
fmt.Println(colorhelper.MakeCSSColorRepresentation(c, colorhelper.HSLAFunctionRepresentation)) // hsla(0, 100%, 50%, 0.5)
fmt.Println(colorhelper.MakeCSSColorRepresentation(c, colorhelper.AnyRepresentation)) // rgba(255, 0, 0, 0.5)
```

### Tweak Color Template Function

```
{{ tweakcolor "#000000" "r" 255 }} => "#FF0000"
{{ tweakcolor "hsl(0, 100%, 0%)" "lightness" 0.5 }} => "hsl(0, 100%, 50%)"
```
