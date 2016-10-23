package parser

import (
	"fmt"
)

const (
	// Misc
	itemError itemType = iota
	itemEOF
	itemSpace
	itemNewline
	itemComment

	// Codes
	itemGCode
	itemMCode

	// Params
	itemEParam
	itemFParam
	itemPParam
	itemSParam
	itemXParam
	itemYParam
	itemZParam
)

type codeList map[int]codeMap
type codeMap map[int]Code
type actionFunc func() error

type paramList map[int]codeMap

var eof = rune(0)
var digits = "0987654321"

// Code is the gosuvius representation of a G Code / M Code (Ex: G0)
type Code struct {
	Do actionFunc
}

// Param is the gosuvius representation of a G / M Code parameter (Ex: 'X')
type Param struct {
}

func makeCodeMap(allowed ...rune) codeMap {
	result := codeMap{}
	for i, raw := range allowed {
		fmt.Println(raw)
		result[i+1] = Code{}
	}
	return result
}

// SupportedCodes lists the G and M codes supported by vesuvius
var SupportedCodes = map[rune]codeList{
	'G': codeList{
		0:  makeCodeMap('X', 'Y', 'Z', 'E', 'F', 'S'),      // Go to (fast)
		1:  makeCodeMap('X', 'Y', 'Z', 'E', 'F', 'S'),      // Go to
		2:  makeCodeMap('X', 'Y', 'Z', 'E', 'I', 'J', 'F'), // Clockwise arc
		3:  makeCodeMap('X', 'Y', 'Z', 'E', 'I', 'J', 'F'), // Counter-clockwise arc
		4:  makeCodeMap('P', 'S'),                          // Dwell
		20: makeCodeMap(),                                  // Units = inches
		21: makeCodeMap(),                                  // Units = mm
		28: makeCodeMap('X', 'Y', 'Z'),                     // Go home on all or selected axes
		90: makeCodeMap(),                                  // Absolute coords
		91: makeCodeMap(),                                  // Relative coords
		92: makeCodeMap('X', 'Y', 'Z', 'E'),                // Set position (reset all to 0 w/ no params)
	},
	'M': codeList{
		0: makeCodeMap(),
	},
}

// SupportedParameters lists the parameters supported by vesuvius
var SupportedParameters = map[rune]Param{
	'E': Param{},
	'F': Param{},
	'P': Param{},
	'S': Param{},
	'X': Param{},
	'Y': Param{},
	'Z': Param{},
}

// LeaderMapping lists the X/Y/Z/etc. params that are supported
var LeaderMapping = map[rune]itemType{
	// CODES
	'G': itemGCode,
	'M': itemMCode,

	// PARAMS
	'E': itemEParam,
	'F': itemFParam,
	'P': itemPParam,
	'S': itemSParam,
	'X': itemXParam,
	'Y': itemYParam,
	'Z': itemZParam,
}
