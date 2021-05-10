package colorhelper

import (
	"fmt"
	"image/color"
	"reflect"
	"strings"
)

func TweakColorTemplateFunc(args ...interface{}) (interface{}, error) {
	if len(args) >= 3 && len(args)%2 == 1 {
		var sourceColorArgument string
		if arg, ok := args[1].(string); ok {
			sourceColorArgument = arg
		} else {
			return nil, fmt.Errorf("bad argument #1 (expected string, got %v)", reflect.TypeOf(args[1]))
		}

		sourceColor, err := ParseCSSColorRepresentation(sourceColorArgument)
		if err != nil {
			return nil, fmt.Errorf("cannot parse source color: %v", err)
		}

		for index := 1; index < len(args)-1; index += 2 {
			var colorParameterName string
			if arg, ok := args[index].(string); ok {
				colorParameterName = arg
			} else {
				return nil, fmt.Errorf("bad argument #%v (expected string, got %v)", index+2, reflect.TypeOf(args[index]))
			}

			switch strings.ToLower(colorParameterName) {
			case "r", "red":
				switch colorParameterValue := args[index+1].(type) {
				case int64:
					if colorParameterValue < 0 || colorParameterValue > 255 {
						return nil, fmt.Errorf("bad argument #%v (invalid red value %v)", index+2, colorParameterValue)
					}
					sourceColorRGBA := color.RGBAModel.Convert(sourceColor).(color.RGBA)
					sourceColorRGBA.R = uint8(colorParameterValue)
					return MakeCSSColorRepresentation(sourceColorRGBA, AnyRepresentation), nil
				case float64:
					if colorParameterValue < 0.0 || colorParameterValue > 1.0 {
						return nil, fmt.Errorf("bad argument #%v (invalid red value %v)", index+2, colorParameterValue)
					}
					sourceColorSRGBA := SRGBAModel.Convert(sourceColor).(SRGBA)
					sourceColorSRGBA.R = colorParameterValue
					return MakeCSSColorRepresentation(sourceColorSRGBA, AnyRepresentation), nil
				default:
					return nil, fmt.Errorf("bad argument #%v (expected int64 or float64, got %v)", index+2, reflect.TypeOf(args[index+1]))
				}
			case "g", "green":
				switch colorParameterValue := args[index+1].(type) {
				case int64:
					if colorParameterValue < 0 || colorParameterValue > 255 {
						return nil, fmt.Errorf("bad argument #%v (invalid green value %v)", index+2, colorParameterValue)
					}
					sourceColorRGBA := color.RGBAModel.Convert(sourceColor).(color.RGBA)
					sourceColorRGBA.G = uint8(colorParameterValue)
					return MakeCSSColorRepresentation(sourceColorRGBA, AnyRepresentation), nil
				case float64:
					if colorParameterValue < 0.0 || colorParameterValue > 1.0 {
						return nil, fmt.Errorf("bad argument #%v (invalid red value %v)", index+2, colorParameterValue)
					}
					sourceColorSRGBA := SRGBAModel.Convert(sourceColor).(SRGBA)
					sourceColorSRGBA.G = colorParameterValue
					return MakeCSSColorRepresentation(sourceColorSRGBA, AnyRepresentation), nil
				default:
					return nil, fmt.Errorf("bad argument #%v (expected int64 or float64, got %v)", index+2, reflect.TypeOf(args[index+1]))
				}
			case "b", "blue":
				switch colorParameterValue := args[index+1].(type) {
				case int64:
					if colorParameterValue < 0 || colorParameterValue > 255 {
						return nil, fmt.Errorf("bad argument #%v (invalid blue value %v)", index+2, colorParameterValue)
					}
					sourceColorRGBA := color.RGBAModel.Convert(sourceColor).(color.RGBA)
					sourceColorRGBA.B = uint8(colorParameterValue)
					return MakeCSSColorRepresentation(sourceColorRGBA, AnyRepresentation), nil
				case float64:
					if colorParameterValue < 0.0 || colorParameterValue > 1.0 {
						return nil, fmt.Errorf("bad argument #%v (invalid blue value %v)", index+2, colorParameterValue)
					}
					sourceColorSRGBA := SRGBAModel.Convert(sourceColor).(SRGBA)
					sourceColorSRGBA.B = colorParameterValue
					return MakeCSSColorRepresentation(sourceColorSRGBA, AnyRepresentation), nil
				default:
					return nil, fmt.Errorf("bad argument #%v (expected int64 or float64, got %v)", index+2, reflect.TypeOf(args[index+1]))
				}
			case "h", "hue":
				if colorParameterValue, ok := args[index+1].(int64); ok {
					if colorParameterValue < 0 || colorParameterValue > 360 {
						return nil, fmt.Errorf("bad argument #%v (invalid hue value %v)", index+2, colorParameterValue)
					}
					sourceColorHSLA := HSLAModel.Convert(sourceColor).(HSLA)
					sourceColorHSLA.H = float64(colorParameterValue)
					return MakeCSSColorRepresentation(sourceColorHSLA, AnyRepresentation), nil
				} else {
					return nil, fmt.Errorf("bad argument #%v (expected int64, got %v)", index+2, reflect.TypeOf(args[index+1]))
				}
			case "s", "saturation":
				if colorParameterValue, ok := args[index+1].(float64); ok {
					if colorParameterValue < 0.0 || colorParameterValue > 1.0 {
						return nil, fmt.Errorf("bad argument #%v (invalid saturation value %v)", index+2, colorParameterValue)
					}
					sourceColorHSLA := HSLAModel.Convert(sourceColor).(HSLA)
					sourceColorHSLA.S = colorParameterValue
					return MakeCSSColorRepresentation(sourceColorHSLA, AnyRepresentation), nil
				} else {
					return nil, fmt.Errorf("bad argument #%v (expected int64, got %v)", index+2, reflect.TypeOf(args[index+1]))
				}
			case "l", "lightness":
				if colorParameterValue, ok := args[index+1].(float64); ok {
					if colorParameterValue < 0.0 || colorParameterValue > 1.0 {
						return nil, fmt.Errorf("bad argument #%v (invalid lightness value %v)", index+2, colorParameterValue)
					}
					sourceColorHSLA := HSLAModel.Convert(sourceColor).(HSLA)
					sourceColorHSLA.S = colorParameterValue
					return MakeCSSColorRepresentation(sourceColorHSLA, AnyRepresentation), nil
				} else {
					return nil, fmt.Errorf("bad argument #%v (expected int64, got %v)", index+2, reflect.TypeOf(args[index+1]))
				}
			case "a", "alpha":
				if colorParameterValue, ok := args[index+1].(float64); ok {
					if colorParameterValue < 0.0 || colorParameterValue > 1.0 {
						return nil, fmt.Errorf("bad argument #%v (invalid lightness value %v)", index+2, colorParameterValue)
					}
					sourceColorSRGBA := SRGBAModel.Convert(sourceColor).(SRGBA)
					sourceColorSRGBA.A = colorParameterValue
					return MakeCSSColorRepresentation(sourceColorSRGBA, AnyRepresentation), nil
				} else {
					return nil, fmt.Errorf("bad argument #%v (expected int64, got %v)", index+2, reflect.TypeOf(args[index+1]))
				}
			}
		}
	}
	return nil, fmt.Errorf("invalid number of arguments given")
}

func PickBestTextColorTemplateFunc(args ...interface{}) (interface{}, error) {
	if len(args) >= 1 {
		var backgroundColorArgument string
		if arg, ok := args[1].(string); ok {
			backgroundColorArgument = arg
		} else {
			return nil, fmt.Errorf("bad argument #1 (expected string, got %v)", reflect.TypeOf(args[1]))
		}

		backgroundColor, err := ParseCSSColorRepresentation(backgroundColorArgument)
		if err != nil {
			return nil, fmt.Errorf("cannot parse background color: %v", err)
		}

		if len(args) == 1 {
			return MakeCSSColorRepresentation(PickBestTextColor(backgroundColor), AnyRepresentation), nil
		} else {
			var textColors []color.Color
			for index := 1; index < len(args); index++ {
				var textColorArgument string
				if arg, ok := args[index].(string); ok {
					textColorArgument = arg
				} else {
					return nil, fmt.Errorf("bad argument #%d (expected string, got %v)", index+1, reflect.TypeOf(arg[index]))
				}

				if textColor, err := ParseCSSColorRepresentation(textColorArgument); err == nil {
					textColors = append(textColors, textColor)
				} else {
					return nil, fmt.Errorf("cannot parse text color #%v: %v", index+1, err)
				}
			}

			return MakeCSSColorRepresentation(PickBestTextColor(backgroundColor, textColors...), AnyRepresentation), nil
		}
	}
	return nil, fmt.Errorf("invalid number of arguments given")
}
