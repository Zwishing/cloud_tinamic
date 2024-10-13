// Code generated from CQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package cql // CQLParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CQLParser struct {
	*antlr.BaseParser
}

var CQLParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cqlparserParserInit() {
	staticData := &CQLParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'<'", "'='", "'>'", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "'#'", "'$'", "'_'", "'\"'", "'%'", "'&'", "", "'('", "')'",
		"'['", "']'", "'*'", "'+'", "','", "'-'", "'.'", "'/'", "':'", "';'",
		"'?'", "'|'", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "''''",
	}
	staticData.SymbolicNames = []string{
		"", "ComparisonOperator", "LT", "EQ", "GT", "NEQ", "GTEQ", "LTEQ", "BooleanLiteral",
		"AND", "OR", "NOT", "LIKE", "ILIKE", "BETWEEN", "IS", "NULL", "IN",
		"ArithmeticOperator", "SpatialOperator", "DistanceOperator", "POINT",
		"LINESTRING", "POLYGON", "MULTIPOINT", "MULTILINESTRING", "MULTIPOLYGON",
		"GEOMETRYCOLLECTION", "ENVELOPE", "NumericLiteral", "Identifier", "IdentifierStart",
		"IdentifierPart", "ALPHA", "DIGIT", "OCTOTHORP", "DOLLAR", "UNDERSCORE",
		"DOUBLEQUOTE", "PERCENT", "AMPERSAND", "QUOTE", "LEFTPAREN", "RIGHTPAREN",
		"LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET", "ASTERISK", "PLUS", "COMMA",
		"MINUS", "PERIOD", "SOLIDUS", "COLON", "SEMICOLON", "QUESTIONMARK",
		"VERTICALBAR", "BIT", "HEXIT", "UnsignedNumericLiteral", "SignedNumericLiteral",
		"ExactNumericLiteral", "ApproximateNumericLiteral", "Mantissa", "Exponent",
		"SignedInteger", "UnsignedInteger", "Sign", "TemporalLiteral", "Instant",
		"FullDate", "DateYear", "DateMonth", "DateDay", "UtcTime", "TimeZoneOffset",
		"TimeHour", "TimeMinute", "TimeSecond", "NOW", "WS", "CharacterStringLiteral",
		"QuotedQuote",
	}
	staticData.RuleNames = []string{
		"cqlFilter", "booleanValueExpression", "booleanTerm", "booleanFactor",
		"booleanPrimary", "predicate", "binaryComparisonPredicate", "likePredicate",
		"betweenPredicate", "inPredicate", "isNullPredicate", "scalarExpression",
		"scalarValue", "propertyName", "characterLiteral", "numericLiteral",
		"booleanLiteral", "temporalLiteral", "spatialPredicate", "distancePredicate",
		"geomExpression", "geomLiteral", "point", "pointList", "linestring",
		"polygon", "polygonDef", "multiPoint", "multiLinestring", "multiPolygon",
		"geometryCollection", "envelope", "coordList", "coordinate",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 81, 328, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 5, 1, 78, 8, 1, 10, 1, 12, 1, 81, 9, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 5, 2, 89, 8, 2, 10, 2, 12, 2, 92, 9, 2, 1, 3, 3, 3, 95,
		8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 104, 8, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 113, 8, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 3, 7, 121, 8, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8,
		128, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 137, 8, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 144, 8, 9, 10, 9, 12, 9, 147, 9, 9, 1,
		9, 1, 9, 1, 9, 5, 9, 152, 8, 9, 10, 9, 12, 9, 155, 9, 9, 3, 9, 157, 8,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 164, 8, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 174, 8, 11, 1, 11, 1, 11,
		1, 11, 5, 11, 179, 8, 11, 10, 11, 12, 11, 182, 9, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 189, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 3, 20, 219, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 229, 8, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		26, 1, 26, 5, 26, 248, 8, 26, 10, 26, 12, 26, 251, 9, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 260, 8, 27, 10, 27, 12, 27, 263,
		9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 272, 8,
		28, 10, 28, 12, 28, 275, 9, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 5, 29, 284, 8, 29, 10, 29, 12, 29, 287, 9, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 296, 8, 30, 10, 30, 12, 30, 299,
		9, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 318, 8, 32,
		10, 32, 12, 32, 321, 9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 0,
		3, 2, 4, 22, 34, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 0, 1, 1, 0, 12, 13, 330, 0, 68, 1, 0, 0, 0, 2, 71, 1, 0, 0, 0, 4, 82,
		1, 0, 0, 0, 6, 94, 1, 0, 0, 0, 8, 103, 1, 0, 0, 0, 10, 112, 1, 0, 0, 0,
		12, 114, 1, 0, 0, 0, 14, 118, 1, 0, 0, 0, 16, 125, 1, 0, 0, 0, 18, 134,
		1, 0, 0, 0, 20, 160, 1, 0, 0, 0, 22, 173, 1, 0, 0, 0, 24, 188, 1, 0, 0,
		0, 26, 190, 1, 0, 0, 0, 28, 192, 1, 0, 0, 0, 30, 194, 1, 0, 0, 0, 32, 196,
		1, 0, 0, 0, 34, 198, 1, 0, 0, 0, 36, 200, 1, 0, 0, 0, 38, 207, 1, 0, 0,
		0, 40, 218, 1, 0, 0, 0, 42, 228, 1, 0, 0, 0, 44, 230, 1, 0, 0, 0, 46, 233,
		1, 0, 0, 0, 48, 237, 1, 0, 0, 0, 50, 240, 1, 0, 0, 0, 52, 243, 1, 0, 0,
		0, 54, 254, 1, 0, 0, 0, 56, 266, 1, 0, 0, 0, 58, 278, 1, 0, 0, 0, 60, 290,
		1, 0, 0, 0, 62, 302, 1, 0, 0, 0, 64, 313, 1, 0, 0, 0, 66, 324, 1, 0, 0,
		0, 68, 69, 3, 2, 1, 0, 69, 70, 5, 0, 0, 1, 70, 1, 1, 0, 0, 0, 71, 72, 6,
		1, -1, 0, 72, 73, 3, 4, 2, 0, 73, 79, 1, 0, 0, 0, 74, 75, 10, 1, 0, 0,
		75, 76, 5, 10, 0, 0, 76, 78, 3, 4, 2, 0, 77, 74, 1, 0, 0, 0, 78, 81, 1,
		0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 3, 1, 0, 0, 0, 81,
		79, 1, 0, 0, 0, 82, 83, 6, 2, -1, 0, 83, 84, 3, 6, 3, 0, 84, 90, 1, 0,
		0, 0, 85, 86, 10, 1, 0, 0, 86, 87, 5, 9, 0, 0, 87, 89, 3, 6, 3, 0, 88,
		85, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0,
		0, 91, 5, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 95, 5, 11, 0, 0, 94, 93,
		1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 3, 8, 4, 0,
		97, 7, 1, 0, 0, 0, 98, 104, 3, 10, 5, 0, 99, 100, 5, 42, 0, 0, 100, 101,
		3, 2, 1, 0, 101, 102, 5, 43, 0, 0, 102, 104, 1, 0, 0, 0, 103, 98, 1, 0,
		0, 0, 103, 99, 1, 0, 0, 0, 104, 9, 1, 0, 0, 0, 105, 113, 3, 12, 6, 0, 106,
		113, 3, 14, 7, 0, 107, 113, 3, 16, 8, 0, 108, 113, 3, 20, 10, 0, 109, 113,
		3, 18, 9, 0, 110, 113, 3, 36, 18, 0, 111, 113, 3, 38, 19, 0, 112, 105,
		1, 0, 0, 0, 112, 106, 1, 0, 0, 0, 112, 107, 1, 0, 0, 0, 112, 108, 1, 0,
		0, 0, 112, 109, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 111, 1, 0, 0, 0,
		113, 11, 1, 0, 0, 0, 114, 115, 3, 22, 11, 0, 115, 116, 5, 1, 0, 0, 116,
		117, 3, 22, 11, 0, 117, 13, 1, 0, 0, 0, 118, 120, 3, 26, 13, 0, 119, 121,
		5, 11, 0, 0, 120, 119, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 1, 0,
		0, 0, 122, 123, 7, 0, 0, 0, 123, 124, 3, 28, 14, 0, 124, 15, 1, 0, 0, 0,
		125, 127, 3, 22, 11, 0, 126, 128, 5, 11, 0, 0, 127, 126, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 5, 14, 0, 0, 130, 131,
		3, 22, 11, 0, 131, 132, 5, 9, 0, 0, 132, 133, 3, 22, 11, 0, 133, 17, 1,
		0, 0, 0, 134, 136, 3, 26, 13, 0, 135, 137, 5, 11, 0, 0, 136, 135, 1, 0,
		0, 0, 136, 137, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 5, 17, 0, 0,
		139, 156, 5, 42, 0, 0, 140, 145, 3, 28, 14, 0, 141, 142, 5, 48, 0, 0, 142,
		144, 3, 28, 14, 0, 143, 141, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143,
		1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 157, 1, 0, 0, 0, 147, 145, 1, 0,
		0, 0, 148, 153, 3, 30, 15, 0, 149, 150, 5, 48, 0, 0, 150, 152, 3, 30, 15,
		0, 151, 149, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 140,
		1, 0, 0, 0, 156, 148, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 5, 43,
		0, 0, 159, 19, 1, 0, 0, 0, 160, 161, 3, 26, 13, 0, 161, 163, 5, 15, 0,
		0, 162, 164, 5, 11, 0, 0, 163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		165, 1, 0, 0, 0, 165, 166, 5, 16, 0, 0, 166, 21, 1, 0, 0, 0, 167, 168,
		6, 11, -1, 0, 168, 174, 3, 24, 12, 0, 169, 170, 5, 42, 0, 0, 170, 171,
		3, 22, 11, 0, 171, 172, 5, 43, 0, 0, 172, 174, 1, 0, 0, 0, 173, 167, 1,
		0, 0, 0, 173, 169, 1, 0, 0, 0, 174, 180, 1, 0, 0, 0, 175, 176, 10, 1, 0,
		0, 176, 177, 5, 18, 0, 0, 177, 179, 3, 22, 11, 2, 178, 175, 1, 0, 0, 0,
		179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		23, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 189, 3, 26, 13, 0, 184, 189,
		3, 28, 14, 0, 185, 189, 3, 30, 15, 0, 186, 189, 3, 32, 16, 0, 187, 189,
		3, 34, 17, 0, 188, 183, 1, 0, 0, 0, 188, 184, 1, 0, 0, 0, 188, 185, 1,
		0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 187, 1, 0, 0, 0, 189, 25, 1, 0, 0,
		0, 190, 191, 5, 30, 0, 0, 191, 27, 1, 0, 0, 0, 192, 193, 5, 80, 0, 0, 193,
		29, 1, 0, 0, 0, 194, 195, 5, 29, 0, 0, 195, 31, 1, 0, 0, 0, 196, 197, 5,
		8, 0, 0, 197, 33, 1, 0, 0, 0, 198, 199, 5, 67, 0, 0, 199, 35, 1, 0, 0,
		0, 200, 201, 5, 19, 0, 0, 201, 202, 5, 42, 0, 0, 202, 203, 3, 40, 20, 0,
		203, 204, 5, 48, 0, 0, 204, 205, 3, 40, 20, 0, 205, 206, 5, 43, 0, 0, 206,
		37, 1, 0, 0, 0, 207, 208, 5, 20, 0, 0, 208, 209, 5, 42, 0, 0, 209, 210,
		3, 40, 20, 0, 210, 211, 5, 48, 0, 0, 211, 212, 3, 40, 20, 0, 212, 213,
		5, 48, 0, 0, 213, 214, 5, 29, 0, 0, 214, 215, 5, 43, 0, 0, 215, 39, 1,
		0, 0, 0, 216, 219, 3, 26, 13, 0, 217, 219, 3, 42, 21, 0, 218, 216, 1, 0,
		0, 0, 218, 217, 1, 0, 0, 0, 219, 41, 1, 0, 0, 0, 220, 229, 3, 44, 22, 0,
		221, 229, 3, 48, 24, 0, 222, 229, 3, 50, 25, 0, 223, 229, 3, 54, 27, 0,
		224, 229, 3, 56, 28, 0, 225, 229, 3, 58, 29, 0, 226, 229, 3, 60, 30, 0,
		227, 229, 3, 62, 31, 0, 228, 220, 1, 0, 0, 0, 228, 221, 1, 0, 0, 0, 228,
		222, 1, 0, 0, 0, 228, 223, 1, 0, 0, 0, 228, 224, 1, 0, 0, 0, 228, 225,
		1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 227, 1, 0, 0, 0, 229, 43, 1, 0,
		0, 0, 230, 231, 5, 21, 0, 0, 231, 232, 3, 46, 23, 0, 232, 45, 1, 0, 0,
		0, 233, 234, 5, 42, 0, 0, 234, 235, 3, 66, 33, 0, 235, 236, 5, 43, 0, 0,
		236, 47, 1, 0, 0, 0, 237, 238, 5, 22, 0, 0, 238, 239, 3, 64, 32, 0, 239,
		49, 1, 0, 0, 0, 240, 241, 5, 23, 0, 0, 241, 242, 3, 52, 26, 0, 242, 51,
		1, 0, 0, 0, 243, 244, 5, 42, 0, 0, 244, 249, 3, 64, 32, 0, 245, 246, 5,
		48, 0, 0, 246, 248, 3, 64, 32, 0, 247, 245, 1, 0, 0, 0, 248, 251, 1, 0,
		0, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 252, 1, 0, 0, 0,
		251, 249, 1, 0, 0, 0, 252, 253, 5, 43, 0, 0, 253, 53, 1, 0, 0, 0, 254,
		255, 5, 24, 0, 0, 255, 256, 5, 42, 0, 0, 256, 261, 3, 46, 23, 0, 257, 258,
		5, 48, 0, 0, 258, 260, 3, 46, 23, 0, 259, 257, 1, 0, 0, 0, 260, 263, 1,
		0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 264, 1, 0, 0,
		0, 263, 261, 1, 0, 0, 0, 264, 265, 5, 43, 0, 0, 265, 55, 1, 0, 0, 0, 266,
		267, 5, 25, 0, 0, 267, 268, 5, 42, 0, 0, 268, 273, 3, 64, 32, 0, 269, 270,
		5, 48, 0, 0, 270, 272, 3, 64, 32, 0, 271, 269, 1, 0, 0, 0, 272, 275, 1,
		0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 276, 1, 0, 0,
		0, 275, 273, 1, 0, 0, 0, 276, 277, 5, 43, 0, 0, 277, 57, 1, 0, 0, 0, 278,
		279, 5, 26, 0, 0, 279, 280, 5, 42, 0, 0, 280, 285, 3, 52, 26, 0, 281, 282,
		5, 48, 0, 0, 282, 284, 3, 52, 26, 0, 283, 281, 1, 0, 0, 0, 284, 287, 1,
		0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 288, 1, 0, 0,
		0, 287, 285, 1, 0, 0, 0, 288, 289, 5, 43, 0, 0, 289, 59, 1, 0, 0, 0, 290,
		291, 5, 27, 0, 0, 291, 292, 5, 42, 0, 0, 292, 297, 3, 42, 21, 0, 293, 294,
		5, 48, 0, 0, 294, 296, 3, 42, 21, 0, 295, 293, 1, 0, 0, 0, 296, 299, 1,
		0, 0, 0, 297, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 300, 1, 0, 0,
		0, 299, 297, 1, 0, 0, 0, 300, 301, 5, 43, 0, 0, 301, 61, 1, 0, 0, 0, 302,
		303, 5, 28, 0, 0, 303, 304, 5, 42, 0, 0, 304, 305, 5, 29, 0, 0, 305, 306,
		5, 48, 0, 0, 306, 307, 5, 29, 0, 0, 307, 308, 5, 48, 0, 0, 308, 309, 5,
		29, 0, 0, 309, 310, 5, 48, 0, 0, 310, 311, 5, 29, 0, 0, 311, 312, 5, 43,
		0, 0, 312, 63, 1, 0, 0, 0, 313, 314, 5, 42, 0, 0, 314, 319, 3, 66, 33,
		0, 315, 316, 5, 48, 0, 0, 316, 318, 3, 66, 33, 0, 317, 315, 1, 0, 0, 0,
		318, 321, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320,
		322, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 322, 323, 5, 43, 0, 0, 323, 65,
		1, 0, 0, 0, 324, 325, 5, 29, 0, 0, 325, 326, 5, 29, 0, 0, 326, 67, 1, 0,
		0, 0, 23, 79, 90, 94, 103, 112, 120, 127, 136, 145, 153, 156, 163, 173,
		180, 188, 218, 228, 249, 261, 273, 285, 297, 319,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CQLParserInit initializes any static state used to implement CQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CQLParserInit() {
	staticData := &CQLParserParserStaticData
	staticData.once.Do(cqlparserParserInit)
}

// NewCQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCQLParser(input antlr.TokenStream) *CQLParser {
	CQLParserInit()
	this := new(CQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CQLParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CQLParser.g4"

	return this
}

// CQLParser tokens.
const (
	CQLParserEOF                       = antlr.TokenEOF
	CQLParserComparisonOperator        = 1
	CQLParserLT                        = 2
	CQLParserEQ                        = 3
	CQLParserGT                        = 4
	CQLParserNEQ                       = 5
	CQLParserGTEQ                      = 6
	CQLParserLTEQ                      = 7
	CQLParserBooleanLiteral            = 8
	CQLParserAND                       = 9
	CQLParserOR                        = 10
	CQLParserNOT                       = 11
	CQLParserLIKE                      = 12
	CQLParserILIKE                     = 13
	CQLParserBETWEEN                   = 14
	CQLParserIS                        = 15
	CQLParserNULL                      = 16
	CQLParserIN                        = 17
	CQLParserArithmeticOperator        = 18
	CQLParserSpatialOperator           = 19
	CQLParserDistanceOperator          = 20
	CQLParserPOINT                     = 21
	CQLParserLINESTRING                = 22
	CQLParserPOLYGON                   = 23
	CQLParserMULTIPOINT                = 24
	CQLParserMULTILINESTRING           = 25
	CQLParserMULTIPOLYGON              = 26
	CQLParserGEOMETRYCOLLECTION        = 27
	CQLParserENVELOPE                  = 28
	CQLParserNumericLiteral            = 29
	CQLParserIdentifier                = 30
	CQLParserIdentifierStart           = 31
	CQLParserIdentifierPart            = 32
	CQLParserALPHA                     = 33
	CQLParserDIGIT                     = 34
	CQLParserOCTOTHORP                 = 35
	CQLParserDOLLAR                    = 36
	CQLParserUNDERSCORE                = 37
	CQLParserDOUBLEQUOTE               = 38
	CQLParserPERCENT                   = 39
	CQLParserAMPERSAND                 = 40
	CQLParserQUOTE                     = 41
	CQLParserLEFTPAREN                 = 42
	CQLParserRIGHTPAREN                = 43
	CQLParserLEFTSQUAREBRACKET         = 44
	CQLParserRIGHTSQUAREBRACKET        = 45
	CQLParserASTERISK                  = 46
	CQLParserPLUS                      = 47
	CQLParserCOMMA                     = 48
	CQLParserMINUS                     = 49
	CQLParserPERIOD                    = 50
	CQLParserSOLIDUS                   = 51
	CQLParserCOLON                     = 52
	CQLParserSEMICOLON                 = 53
	CQLParserQUESTIONMARK              = 54
	CQLParserVERTICALBAR               = 55
	CQLParserBIT                       = 56
	CQLParserHEXIT                     = 57
	CQLParserUnsignedNumericLiteral    = 58
	CQLParserSignedNumericLiteral      = 59
	CQLParserExactNumericLiteral       = 60
	CQLParserApproximateNumericLiteral = 61
	CQLParserMantissa                  = 62
	CQLParserExponent                  = 63
	CQLParserSignedInteger             = 64
	CQLParserUnsignedInteger           = 65
	CQLParserSign                      = 66
	CQLParserTemporalLiteral           = 67
	CQLParserInstant                   = 68
	CQLParserFullDate                  = 69
	CQLParserDateYear                  = 70
	CQLParserDateMonth                 = 71
	CQLParserDateDay                   = 72
	CQLParserUtcTime                   = 73
	CQLParserTimeZoneOffset            = 74
	CQLParserTimeHour                  = 75
	CQLParserTimeMinute                = 76
	CQLParserTimeSecond                = 77
	CQLParserNOW                       = 78
	CQLParserWS                        = 79
	CQLParserCharacterStringLiteral    = 80
	CQLParserQuotedQuote               = 81
)

// CQLParser rules.
const (
	CQLParserRULE_cqlFilter                 = 0
	CQLParserRULE_booleanValueExpression    = 1
	CQLParserRULE_booleanTerm               = 2
	CQLParserRULE_booleanFactor             = 3
	CQLParserRULE_booleanPrimary            = 4
	CQLParserRULE_predicate                 = 5
	CQLParserRULE_binaryComparisonPredicate = 6
	CQLParserRULE_likePredicate             = 7
	CQLParserRULE_betweenPredicate          = 8
	CQLParserRULE_inPredicate               = 9
	CQLParserRULE_isNullPredicate           = 10
	CQLParserRULE_scalarExpression          = 11
	CQLParserRULE_scalarValue               = 12
	CQLParserRULE_propertyName              = 13
	CQLParserRULE_characterLiteral          = 14
	CQLParserRULE_numericLiteral            = 15
	CQLParserRULE_booleanLiteral            = 16
	CQLParserRULE_temporalLiteral           = 17
	CQLParserRULE_spatialPredicate          = 18
	CQLParserRULE_distancePredicate         = 19
	CQLParserRULE_geomExpression            = 20
	CQLParserRULE_geomLiteral               = 21
	CQLParserRULE_point                     = 22
	CQLParserRULE_pointList                 = 23
	CQLParserRULE_linestring                = 24
	CQLParserRULE_polygon                   = 25
	CQLParserRULE_polygonDef                = 26
	CQLParserRULE_multiPoint                = 27
	CQLParserRULE_multiLinestring           = 28
	CQLParserRULE_multiPolygon              = 29
	CQLParserRULE_geometryCollection        = 30
	CQLParserRULE_envelope                  = 31
	CQLParserRULE_coordList                 = 32
	CQLParserRULE_coordinate                = 33
)

// ICqlFilterContext is an interface to support dynamic dispatch.
type ICqlFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BooleanValueExpression() IBooleanValueExpressionContext
	EOF() antlr.TerminalNode

	// IsCqlFilterContext differentiates from other interfaces.
	IsCqlFilterContext()
}

type CqlFilterContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCqlFilterContext() *CqlFilterContext {
	var p = new(CqlFilterContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_cqlFilter
	return p
}

func InitEmptyCqlFilterContext(p *CqlFilterContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_cqlFilter
}

func (*CqlFilterContext) IsCqlFilterContext() {}

func NewCqlFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlFilterContext {
	var p = new(CqlFilterContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_cqlFilter

	return p
}

func (s *CqlFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlFilterContext) BooleanValueExpression() IBooleanValueExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanValueExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanValueExpressionContext)
}

func (s *CqlFilterContext) EOF() antlr.TerminalNode {
	return s.GetToken(CQLParserEOF, 0)
}

func (s *CqlFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCqlFilter(s)
	}
}

func (s *CqlFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCqlFilter(s)
	}
}

func (p *CQLParser) CqlFilter() (localctx ICqlFilterContext) {
	localctx = NewCqlFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CQLParserRULE_cqlFilter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.booleanValueExpression(0)
	}
	{
		p.SetState(69)
		p.Match(CQLParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanValueExpressionContext is an interface to support dynamic dispatch.
type IBooleanValueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BooleanTerm() IBooleanTermContext
	BooleanValueExpression() IBooleanValueExpressionContext
	OR() antlr.TerminalNode

	// IsBooleanValueExpressionContext differentiates from other interfaces.
	IsBooleanValueExpressionContext()
}

type BooleanValueExpressionContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanValueExpressionContext() *BooleanValueExpressionContext {
	var p = new(BooleanValueExpressionContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_booleanValueExpression
	return p
}

func InitEmptyBooleanValueExpressionContext(p *BooleanValueExpressionContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_booleanValueExpression
}

func (*BooleanValueExpressionContext) IsBooleanValueExpressionContext() {}

func NewBooleanValueExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueExpressionContext {
	var p = new(BooleanValueExpressionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanValueExpression

	return p
}

func (s *BooleanValueExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanValueExpressionContext) BooleanTerm() IBooleanTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanTermContext)
}

func (s *BooleanValueExpressionContext) BooleanValueExpression() IBooleanValueExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanValueExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanValueExpressionContext)
}

func (s *BooleanValueExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(CQLParserOR, 0)
}

func (s *BooleanValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanValueExpression(s)
	}
}

func (s *BooleanValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanValueExpression(s)
	}
}

func (p *CQLParser) BooleanValueExpression() (localctx IBooleanValueExpressionContext) {
	return p.booleanValueExpression(0)
}

func (p *CQLParser) booleanValueExpression(_p int) (localctx IBooleanValueExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewBooleanValueExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanValueExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CQLParserRULE_booleanValueExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.booleanTerm(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBooleanValueExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CQLParserRULE_booleanValueExpression)
			p.SetState(74)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(75)
				p.Match(CQLParserOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(76)
				p.booleanTerm(0)
			}

		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanTermContext is an interface to support dynamic dispatch.
type IBooleanTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BooleanFactor() IBooleanFactorContext
	BooleanTerm() IBooleanTermContext
	AND() antlr.TerminalNode

	// IsBooleanTermContext differentiates from other interfaces.
	IsBooleanTermContext()
}

type BooleanTermContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanTermContext() *BooleanTermContext {
	var p = new(BooleanTermContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_booleanTerm
	return p
}

func InitEmptyBooleanTermContext(p *BooleanTermContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_booleanTerm
}

func (*BooleanTermContext) IsBooleanTermContext() {}

func NewBooleanTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanTermContext {
	var p = new(BooleanTermContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanTerm

	return p
}

func (s *BooleanTermContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanTermContext) BooleanFactor() IBooleanFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanFactorContext)
}

func (s *BooleanTermContext) BooleanTerm() IBooleanTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanTermContext)
}

func (s *BooleanTermContext) AND() antlr.TerminalNode {
	return s.GetToken(CQLParserAND, 0)
}

func (s *BooleanTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanTerm(s)
	}
}

func (s *BooleanTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanTerm(s)
	}
}

func (p *CQLParser) BooleanTerm() (localctx IBooleanTermContext) {
	return p.booleanTerm(0)
}

func (p *CQLParser) booleanTerm(_p int) (localctx IBooleanTermContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewBooleanTermContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanTermContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, CQLParserRULE_booleanTerm, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.BooleanFactor()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBooleanTermContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CQLParserRULE_booleanTerm)
			p.SetState(85)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(86)
				p.Match(CQLParserAND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(87)
				p.BooleanFactor()
			}

		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanFactorContext is an interface to support dynamic dispatch.
type IBooleanFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BooleanPrimary() IBooleanPrimaryContext
	NOT() antlr.TerminalNode

	// IsBooleanFactorContext differentiates from other interfaces.
	IsBooleanFactorContext()
}

type BooleanFactorContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanFactorContext() *BooleanFactorContext {
	var p = new(BooleanFactorContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_booleanFactor
	return p
}

func InitEmptyBooleanFactorContext(p *BooleanFactorContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_booleanFactor
}

func (*BooleanFactorContext) IsBooleanFactorContext() {}

func NewBooleanFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanFactorContext {
	var p = new(BooleanFactorContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanFactor

	return p
}

func (s *BooleanFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanFactorContext) BooleanPrimary() IBooleanPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanPrimaryContext)
}

func (s *BooleanFactorContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *BooleanFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanFactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanFactor(s)
	}
}

func (s *BooleanFactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanFactor(s)
	}
}

func (p *CQLParser) BooleanFactor() (localctx IBooleanFactorContext) {
	localctx = NewBooleanFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CQLParserRULE_booleanFactor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(93)
			p.Match(CQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(96)
		p.BooleanPrimary()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanPrimaryContext is an interface to support dynamic dispatch.
type IBooleanPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Predicate() IPredicateContext
	LEFTPAREN() antlr.TerminalNode
	BooleanValueExpression() IBooleanValueExpressionContext
	RIGHTPAREN() antlr.TerminalNode

	// IsBooleanPrimaryContext differentiates from other interfaces.
	IsBooleanPrimaryContext()
}

type BooleanPrimaryContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanPrimaryContext() *BooleanPrimaryContext {
	var p = new(BooleanPrimaryContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_booleanPrimary
	return p
}

func InitEmptyBooleanPrimaryContext(p *BooleanPrimaryContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_booleanPrimary
}

func (*BooleanPrimaryContext) IsBooleanPrimaryContext() {}

func NewBooleanPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanPrimaryContext {
	var p = new(BooleanPrimaryContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanPrimary

	return p
}

func (s *BooleanPrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanPrimaryContext) Predicate() IPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *BooleanPrimaryContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *BooleanPrimaryContext) BooleanValueExpression() IBooleanValueExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanValueExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanValueExpressionContext)
}

func (s *BooleanPrimaryContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *BooleanPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanPrimary(s)
	}
}

func (s *BooleanPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanPrimary(s)
	}
}

func (p *CQLParser) BooleanPrimary() (localctx IBooleanPrimaryContext) {
	localctx = NewBooleanPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CQLParserRULE_booleanPrimary)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Predicate()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(CQLParserLEFTPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.booleanValueExpression(0)
		}
		{
			p.SetState(101)
			p.Match(CQLParserRIGHTPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BinaryComparisonPredicate() IBinaryComparisonPredicateContext
	LikePredicate() ILikePredicateContext
	BetweenPredicate() IBetweenPredicateContext
	IsNullPredicate() IIsNullPredicateContext
	InPredicate() IInPredicateContext
	SpatialPredicate() ISpatialPredicateContext
	DistancePredicate() IDistancePredicateContext

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_predicate
	return p
}

func InitEmptyPredicateContext(p *PredicateContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_predicate
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) BinaryComparisonPredicate() IBinaryComparisonPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryComparisonPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryComparisonPredicateContext)
}

func (s *PredicateContext) LikePredicate() ILikePredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILikePredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILikePredicateContext)
}

func (s *PredicateContext) BetweenPredicate() IBetweenPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBetweenPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBetweenPredicateContext)
}

func (s *PredicateContext) IsNullPredicate() IIsNullPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsNullPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsNullPredicateContext)
}

func (s *PredicateContext) InPredicate() IInPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInPredicateContext)
}

func (s *PredicateContext) SpatialPredicate() ISpatialPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpatialPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpatialPredicateContext)
}

func (s *PredicateContext) DistancePredicate() IDistancePredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistancePredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistancePredicateContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *CQLParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CQLParserRULE_predicate)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.BinaryComparisonPredicate()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.LikePredicate()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(107)
			p.BetweenPredicate()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)
			p.IsNullPredicate()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(109)
			p.InPredicate()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(110)
			p.SpatialPredicate()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(111)
			p.DistancePredicate()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryComparisonPredicateContext is an interface to support dynamic dispatch.
type IBinaryComparisonPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllScalarExpression() []IScalarExpressionContext
	ScalarExpression(i int) IScalarExpressionContext
	ComparisonOperator() antlr.TerminalNode

	// IsBinaryComparisonPredicateContext differentiates from other interfaces.
	IsBinaryComparisonPredicateContext()
}

type BinaryComparisonPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBinaryComparisonPredicateContext() *BinaryComparisonPredicateContext {
	var p = new(BinaryComparisonPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_binaryComparisonPredicate
	return p
}

func InitEmptyBinaryComparisonPredicateContext(p *BinaryComparisonPredicateContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_binaryComparisonPredicate
}

func (*BinaryComparisonPredicateContext) IsBinaryComparisonPredicateContext() {}

func NewBinaryComparisonPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryComparisonPredicateContext {
	var p = new(BinaryComparisonPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_binaryComparisonPredicate

	return p
}

func (s *BinaryComparisonPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryComparisonPredicateContext) AllScalarExpression() []IScalarExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalarExpressionContext); ok {
			len++
		}
	}

	tst := make([]IScalarExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalarExpressionContext); ok {
			tst[i] = t.(IScalarExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryComparisonPredicateContext) ScalarExpression(i int) IScalarExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarExpressionContext)
}

func (s *BinaryComparisonPredicateContext) ComparisonOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserComparisonOperator, 0)
}

func (s *BinaryComparisonPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryComparisonPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryComparisonPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBinaryComparisonPredicate(s)
	}
}

func (s *BinaryComparisonPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBinaryComparisonPredicate(s)
	}
}

func (p *CQLParser) BinaryComparisonPredicate() (localctx IBinaryComparisonPredicateContext) {
	localctx = NewBinaryComparisonPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CQLParserRULE_binaryComparisonPredicate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.scalarExpression(0)
	}
	{
		p.SetState(115)
		p.Match(CQLParserComparisonOperator)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.scalarExpression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILikePredicateContext is an interface to support dynamic dispatch.
type ILikePredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyName() IPropertyNameContext
	CharacterLiteral() ICharacterLiteralContext
	LIKE() antlr.TerminalNode
	ILIKE() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsLikePredicateContext differentiates from other interfaces.
	IsLikePredicateContext()
}

type LikePredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyLikePredicateContext() *LikePredicateContext {
	var p = new(LikePredicateContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_likePredicate
	return p
}

func InitEmptyLikePredicateContext(p *LikePredicateContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_likePredicate
}

func (*LikePredicateContext) IsLikePredicateContext() {}

func NewLikePredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikePredicateContext {
	var p = new(LikePredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_likePredicate

	return p
}

func (s *LikePredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *LikePredicateContext) PropertyName() IPropertyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *LikePredicateContext) CharacterLiteral() ICharacterLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharacterLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *LikePredicateContext) LIKE() antlr.TerminalNode {
	return s.GetToken(CQLParserLIKE, 0)
}

func (s *LikePredicateContext) ILIKE() antlr.TerminalNode {
	return s.GetToken(CQLParserILIKE, 0)
}

func (s *LikePredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *LikePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikePredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterLikePredicate(s)
	}
}

func (s *LikePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitLikePredicate(s)
	}
}

func (p *CQLParser) LikePredicate() (localctx ILikePredicateContext) {
	localctx = NewLikePredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CQLParserRULE_likePredicate)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.PropertyName()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(119)
			p.Match(CQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(122)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CQLParserLIKE || _la == CQLParserILIKE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(123)
		p.CharacterLiteral()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBetweenPredicateContext is an interface to support dynamic dispatch.
type IBetweenPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllScalarExpression() []IScalarExpressionContext
	ScalarExpression(i int) IScalarExpressionContext
	BETWEEN() antlr.TerminalNode
	AND() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsBetweenPredicateContext differentiates from other interfaces.
	IsBetweenPredicateContext()
}

type BetweenPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBetweenPredicateContext() *BetweenPredicateContext {
	var p = new(BetweenPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_betweenPredicate
	return p
}

func InitEmptyBetweenPredicateContext(p *BetweenPredicateContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_betweenPredicate
}

func (*BetweenPredicateContext) IsBetweenPredicateContext() {}

func NewBetweenPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BetweenPredicateContext {
	var p = new(BetweenPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_betweenPredicate

	return p
}

func (s *BetweenPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *BetweenPredicateContext) AllScalarExpression() []IScalarExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalarExpressionContext); ok {
			len++
		}
	}

	tst := make([]IScalarExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalarExpressionContext); ok {
			tst[i] = t.(IScalarExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BetweenPredicateContext) ScalarExpression(i int) IScalarExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarExpressionContext)
}

func (s *BetweenPredicateContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(CQLParserBETWEEN, 0)
}

func (s *BetweenPredicateContext) AND() antlr.TerminalNode {
	return s.GetToken(CQLParserAND, 0)
}

func (s *BetweenPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *BetweenPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BetweenPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBetweenPredicate(s)
	}
}

func (s *BetweenPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBetweenPredicate(s)
	}
}

func (p *CQLParser) BetweenPredicate() (localctx IBetweenPredicateContext) {
	localctx = NewBetweenPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CQLParserRULE_betweenPredicate)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.scalarExpression(0)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(126)
			p.Match(CQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(129)
		p.Match(CQLParserBETWEEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.scalarExpression(0)
	}
	{
		p.SetState(131)
		p.Match(CQLParserAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.scalarExpression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInPredicateContext is an interface to support dynamic dispatch.
type IInPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyName() IPropertyNameContext
	IN() antlr.TerminalNode
	LEFTPAREN() antlr.TerminalNode
	RIGHTPAREN() antlr.TerminalNode
	AllCharacterLiteral() []ICharacterLiteralContext
	CharacterLiteral(i int) ICharacterLiteralContext
	AllNumericLiteral() []INumericLiteralContext
	NumericLiteral(i int) INumericLiteralContext
	NOT() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsInPredicateContext differentiates from other interfaces.
	IsInPredicateContext()
}

type InPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyInPredicateContext() *InPredicateContext {
	var p = new(InPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_inPredicate
	return p
}

func InitEmptyInPredicateContext(p *InPredicateContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_inPredicate
}

func (*InPredicateContext) IsInPredicateContext() {}

func NewInPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InPredicateContext {
	var p = new(InPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_inPredicate

	return p
}

func (s *InPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *InPredicateContext) PropertyName() IPropertyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *InPredicateContext) IN() antlr.TerminalNode {
	return s.GetToken(CQLParserIN, 0)
}

func (s *InPredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *InPredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *InPredicateContext) AllCharacterLiteral() []ICharacterLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICharacterLiteralContext); ok {
			len++
		}
	}

	tst := make([]ICharacterLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICharacterLiteralContext); ok {
			tst[i] = t.(ICharacterLiteralContext)
			i++
		}
	}

	return tst
}

func (s *InPredicateContext) CharacterLiteral(i int) ICharacterLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharacterLiteralContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *InPredicateContext) AllNumericLiteral() []INumericLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumericLiteralContext); ok {
			len++
		}
	}

	tst := make([]INumericLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumericLiteralContext); ok {
			tst[i] = t.(INumericLiteralContext)
			i++
		}
	}

	return tst
}

func (s *InPredicateContext) NumericLiteral(i int) INumericLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *InPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *InPredicateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *InPredicateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *InPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterInPredicate(s)
	}
}

func (s *InPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitInPredicate(s)
	}
}

func (p *CQLParser) InPredicate() (localctx IInPredicateContext) {
	localctx = NewInPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CQLParserRULE_inPredicate)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.PropertyName()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(135)
			p.Match(CQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(138)
		p.Match(CQLParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CQLParserCharacterStringLiteral:
		{
			p.SetState(140)
			p.CharacterLiteral()
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CQLParserCOMMA {
			{
				p.SetState(141)
				p.Match(CQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(142)
				p.CharacterLiteral()
			}

			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case CQLParserNumericLiteral:
		{
			p.SetState(148)
			p.NumericLiteral()
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CQLParserCOMMA {
			{
				p.SetState(149)
				p.Match(CQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(150)
				p.NumericLiteral()
			}

			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(158)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIsNullPredicateContext is an interface to support dynamic dispatch.
type IIsNullPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyName() IPropertyNameContext
	IS() antlr.TerminalNode
	NULL() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsIsNullPredicateContext differentiates from other interfaces.
	IsIsNullPredicateContext()
}

type IsNullPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyIsNullPredicateContext() *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_isNullPredicate
	return p
}

func InitEmptyIsNullPredicateContext(p *IsNullPredicateContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_isNullPredicate
}

func (*IsNullPredicateContext) IsIsNullPredicateContext() {}

func NewIsNullPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_isNullPredicate

	return p
}

func (s *IsNullPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *IsNullPredicateContext) PropertyName() IPropertyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *IsNullPredicateContext) IS() antlr.TerminalNode {
	return s.GetToken(CQLParserIS, 0)
}

func (s *IsNullPredicateContext) NULL() antlr.TerminalNode {
	return s.GetToken(CQLParserNULL, 0)
}

func (s *IsNullPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *IsNullPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsNullPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterIsNullPredicate(s)
	}
}

func (s *IsNullPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitIsNullPredicate(s)
	}
}

func (p *CQLParser) IsNullPredicate() (localctx IIsNullPredicateContext) {
	localctx = NewIsNullPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CQLParserRULE_isNullPredicate)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.PropertyName()
	}
	{
		p.SetState(161)
		p.Match(CQLParserIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(162)
			p.Match(CQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(165)
		p.Match(CQLParserNULL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScalarExpressionContext is an interface to support dynamic dispatch.
type IScalarExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ScalarValue() IScalarValueContext
	LEFTPAREN() antlr.TerminalNode
	AllScalarExpression() []IScalarExpressionContext
	ScalarExpression(i int) IScalarExpressionContext
	RIGHTPAREN() antlr.TerminalNode
	ArithmeticOperator() antlr.TerminalNode

	// IsScalarExpressionContext differentiates from other interfaces.
	IsScalarExpressionContext()
}

type ScalarExpressionContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyScalarExpressionContext() *ScalarExpressionContext {
	var p = new(ScalarExpressionContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_scalarExpression
	return p
}

func InitEmptyScalarExpressionContext(p *ScalarExpressionContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_scalarExpression
}

func (*ScalarExpressionContext) IsScalarExpressionContext() {}

func NewScalarExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarExpressionContext {
	var p = new(ScalarExpressionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_scalarExpression

	return p
}

func (s *ScalarExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarExpressionContext) ScalarValue() IScalarValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarValueContext)
}

func (s *ScalarExpressionContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *ScalarExpressionContext) AllScalarExpression() []IScalarExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalarExpressionContext); ok {
			len++
		}
	}

	tst := make([]IScalarExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalarExpressionContext); ok {
			tst[i] = t.(IScalarExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ScalarExpressionContext) ScalarExpression(i int) IScalarExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarExpressionContext)
}

func (s *ScalarExpressionContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *ScalarExpressionContext) ArithmeticOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserArithmeticOperator, 0)
}

func (s *ScalarExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterScalarExpression(s)
	}
}

func (s *ScalarExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitScalarExpression(s)
	}
}

func (p *CQLParser) ScalarExpression() (localctx IScalarExpressionContext) {
	return p.scalarExpression(0)
}

func (p *CQLParser) scalarExpression(_p int) (localctx IScalarExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewScalarExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IScalarExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, CQLParserRULE_scalarExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CQLParserBooleanLiteral, CQLParserNumericLiteral, CQLParserIdentifier, CQLParserTemporalLiteral, CQLParserCharacterStringLiteral:
		{
			p.SetState(168)
			p.ScalarValue()
		}

	case CQLParserLEFTPAREN:
		{
			p.SetState(169)
			p.Match(CQLParserLEFTPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.scalarExpression(0)
		}
		{
			p.SetState(171)
			p.Match(CQLParserRIGHTPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewScalarExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CQLParserRULE_scalarExpression)
			p.SetState(175)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(176)
				p.Match(CQLParserArithmeticOperator)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(177)
				p.scalarExpression(2)
			}

		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScalarValueContext is an interface to support dynamic dispatch.
type IScalarValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyName() IPropertyNameContext
	CharacterLiteral() ICharacterLiteralContext
	NumericLiteral() INumericLiteralContext
	BooleanLiteral() IBooleanLiteralContext
	TemporalLiteral() ITemporalLiteralContext

	// IsScalarValueContext differentiates from other interfaces.
	IsScalarValueContext()
}

type ScalarValueContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyScalarValueContext() *ScalarValueContext {
	var p = new(ScalarValueContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_scalarValue
	return p
}

func InitEmptyScalarValueContext(p *ScalarValueContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_scalarValue
}

func (*ScalarValueContext) IsScalarValueContext() {}

func NewScalarValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarValueContext {
	var p = new(ScalarValueContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_scalarValue

	return p
}

func (s *ScalarValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarValueContext) PropertyName() IPropertyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *ScalarValueContext) CharacterLiteral() ICharacterLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharacterLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *ScalarValueContext) NumericLiteral() INumericLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *ScalarValueContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ScalarValueContext) TemporalLiteral() ITemporalLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITemporalLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITemporalLiteralContext)
}

func (s *ScalarValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterScalarValue(s)
	}
}

func (s *ScalarValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitScalarValue(s)
	}
}

func (p *CQLParser) ScalarValue() (localctx IScalarValueContext) {
	localctx = NewScalarValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CQLParserRULE_scalarValue)
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CQLParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.PropertyName()
		}

	case CQLParserCharacterStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.CharacterLiteral()
		}

	case CQLParserNumericLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(185)
			p.NumericLiteral()
		}

	case CQLParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(186)
			p.BooleanLiteral()
		}

	case CQLParserTemporalLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(187)
			p.TemporalLiteral()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_propertyName
	return p
}

func InitEmptyPropertyNameContext(p *PropertyNameContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_propertyName
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CQLParserIdentifier, 0)
}

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (p *CQLParser) PropertyName() (localctx IPropertyNameContext) {
	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CQLParserRULE_propertyName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(CQLParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharacterLiteralContext is an interface to support dynamic dispatch.
type ICharacterLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CharacterStringLiteral() antlr.TerminalNode

	// IsCharacterLiteralContext differentiates from other interfaces.
	IsCharacterLiteralContext()
}

type CharacterLiteralContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCharacterLiteralContext() *CharacterLiteralContext {
	var p = new(CharacterLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_characterLiteral
	return p
}

func InitEmptyCharacterLiteralContext(p *CharacterLiteralContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_characterLiteral
}

func (*CharacterLiteralContext) IsCharacterLiteralContext() {}

func NewCharacterLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterLiteralContext {
	var p = new(CharacterLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_characterLiteral

	return p
}

func (s *CharacterLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterLiteralContext) CharacterStringLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserCharacterStringLiteral, 0)
}

func (s *CharacterLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCharacterLiteral(s)
	}
}

func (s *CharacterLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCharacterLiteral(s)
	}
}

func (p *CQLParser) CharacterLiteral() (localctx ICharacterLiteralContext) {
	localctx = NewCharacterLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CQLParserRULE_characterLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(CQLParserCharacterStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumericLiteral() antlr.TerminalNode

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_numericLiteral
	return p
}

func InitEmptyNumericLiteralContext(p *NumericLiteralContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_numericLiteral
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, 0)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (p *CQLParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CQLParserRULE_numericLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(CQLParserNumericLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BooleanLiteral() antlr.TerminalNode

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_booleanLiteral
	return p
}

func InitEmptyBooleanLiteralContext(p *BooleanLiteralContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_booleanLiteral
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserBooleanLiteral, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *CQLParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CQLParserRULE_booleanLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(CQLParserBooleanLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITemporalLiteralContext is an interface to support dynamic dispatch.
type ITemporalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TemporalLiteral() antlr.TerminalNode

	// IsTemporalLiteralContext differentiates from other interfaces.
	IsTemporalLiteralContext()
}

type TemporalLiteralContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyTemporalLiteralContext() *TemporalLiteralContext {
	var p = new(TemporalLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_temporalLiteral
	return p
}

func InitEmptyTemporalLiteralContext(p *TemporalLiteralContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_temporalLiteral
}

func (*TemporalLiteralContext) IsTemporalLiteralContext() {}

func NewTemporalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemporalLiteralContext {
	var p = new(TemporalLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_temporalLiteral

	return p
}

func (s *TemporalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *TemporalLiteralContext) TemporalLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserTemporalLiteral, 0)
}

func (s *TemporalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemporalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemporalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterTemporalLiteral(s)
	}
}

func (s *TemporalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitTemporalLiteral(s)
	}
}

func (p *CQLParser) TemporalLiteral() (localctx ITemporalLiteralContext) {
	localctx = NewTemporalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CQLParserRULE_temporalLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(CQLParserTemporalLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISpatialPredicateContext is an interface to support dynamic dispatch.
type ISpatialPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SpatialOperator() antlr.TerminalNode
	LEFTPAREN() antlr.TerminalNode
	AllGeomExpression() []IGeomExpressionContext
	GeomExpression(i int) IGeomExpressionContext
	COMMA() antlr.TerminalNode
	RIGHTPAREN() antlr.TerminalNode

	// IsSpatialPredicateContext differentiates from other interfaces.
	IsSpatialPredicateContext()
}

type SpatialPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptySpatialPredicateContext() *SpatialPredicateContext {
	var p = new(SpatialPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_spatialPredicate
	return p
}

func InitEmptySpatialPredicateContext(p *SpatialPredicateContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_spatialPredicate
}

func (*SpatialPredicateContext) IsSpatialPredicateContext() {}

func NewSpatialPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpatialPredicateContext {
	var p = new(SpatialPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_spatialPredicate

	return p
}

func (s *SpatialPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *SpatialPredicateContext) SpatialOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserSpatialOperator, 0)
}

func (s *SpatialPredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *SpatialPredicateContext) AllGeomExpression() []IGeomExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGeomExpressionContext); ok {
			len++
		}
	}

	tst := make([]IGeomExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGeomExpressionContext); ok {
			tst[i] = t.(IGeomExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SpatialPredicateContext) GeomExpression(i int) IGeomExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeomExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeomExpressionContext)
}

func (s *SpatialPredicateContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, 0)
}

func (s *SpatialPredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *SpatialPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpatialPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpatialPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterSpatialPredicate(s)
	}
}

func (s *SpatialPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitSpatialPredicate(s)
	}
}

func (p *CQLParser) SpatialPredicate() (localctx ISpatialPredicateContext) {
	localctx = NewSpatialPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CQLParserRULE_spatialPredicate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(CQLParserSpatialOperator)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.GeomExpression()
	}
	{
		p.SetState(203)
		p.Match(CQLParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.GeomExpression()
	}
	{
		p.SetState(205)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDistancePredicateContext is an interface to support dynamic dispatch.
type IDistancePredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DistanceOperator() antlr.TerminalNode
	LEFTPAREN() antlr.TerminalNode
	AllGeomExpression() []IGeomExpressionContext
	GeomExpression(i int) IGeomExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	NumericLiteral() antlr.TerminalNode
	RIGHTPAREN() antlr.TerminalNode

	// IsDistancePredicateContext differentiates from other interfaces.
	IsDistancePredicateContext()
}

type DistancePredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyDistancePredicateContext() *DistancePredicateContext {
	var p = new(DistancePredicateContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_distancePredicate
	return p
}

func InitEmptyDistancePredicateContext(p *DistancePredicateContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_distancePredicate
}

func (*DistancePredicateContext) IsDistancePredicateContext() {}

func NewDistancePredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistancePredicateContext {
	var p = new(DistancePredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_distancePredicate

	return p
}

func (s *DistancePredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *DistancePredicateContext) DistanceOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserDistanceOperator, 0)
}

func (s *DistancePredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *DistancePredicateContext) AllGeomExpression() []IGeomExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGeomExpressionContext); ok {
			len++
		}
	}

	tst := make([]IGeomExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGeomExpressionContext); ok {
			tst[i] = t.(IGeomExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DistancePredicateContext) GeomExpression(i int) IGeomExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeomExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeomExpressionContext)
}

func (s *DistancePredicateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *DistancePredicateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *DistancePredicateContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, 0)
}

func (s *DistancePredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *DistancePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistancePredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistancePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterDistancePredicate(s)
	}
}

func (s *DistancePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitDistancePredicate(s)
	}
}

func (p *CQLParser) DistancePredicate() (localctx IDistancePredicateContext) {
	localctx = NewDistancePredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CQLParserRULE_distancePredicate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(CQLParserDistanceOperator)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.GeomExpression()
	}
	{
		p.SetState(210)
		p.Match(CQLParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.GeomExpression()
	}
	{
		p.SetState(212)
		p.Match(CQLParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(CQLParserNumericLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGeomExpressionContext is an interface to support dynamic dispatch.
type IGeomExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyName() IPropertyNameContext
	GeomLiteral() IGeomLiteralContext

	// IsGeomExpressionContext differentiates from other interfaces.
	IsGeomExpressionContext()
}

type GeomExpressionContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyGeomExpressionContext() *GeomExpressionContext {
	var p = new(GeomExpressionContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_geomExpression
	return p
}

func InitEmptyGeomExpressionContext(p *GeomExpressionContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_geomExpression
}

func (*GeomExpressionContext) IsGeomExpressionContext() {}

func NewGeomExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeomExpressionContext {
	var p = new(GeomExpressionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_geomExpression

	return p
}

func (s *GeomExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *GeomExpressionContext) PropertyName() IPropertyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *GeomExpressionContext) GeomLiteral() IGeomLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeomLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeomLiteralContext)
}

func (s *GeomExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeomExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeomExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterGeomExpression(s)
	}
}

func (s *GeomExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitGeomExpression(s)
	}
}

func (p *CQLParser) GeomExpression() (localctx IGeomExpressionContext) {
	localctx = NewGeomExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CQLParserRULE_geomExpression)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CQLParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.PropertyName()
		}

	case CQLParserPOINT, CQLParserLINESTRING, CQLParserPOLYGON, CQLParserMULTIPOINT, CQLParserMULTILINESTRING, CQLParserMULTIPOLYGON, CQLParserGEOMETRYCOLLECTION, CQLParserENVELOPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.GeomLiteral()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGeomLiteralContext is an interface to support dynamic dispatch.
type IGeomLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Point() IPointContext
	Linestring() ILinestringContext
	Polygon() IPolygonContext
	MultiPoint() IMultiPointContext
	MultiLinestring() IMultiLinestringContext
	MultiPolygon() IMultiPolygonContext
	GeometryCollection() IGeometryCollectionContext
	Envelope() IEnvelopeContext

	// IsGeomLiteralContext differentiates from other interfaces.
	IsGeomLiteralContext()
}

type GeomLiteralContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyGeomLiteralContext() *GeomLiteralContext {
	var p = new(GeomLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_geomLiteral
	return p
}

func InitEmptyGeomLiteralContext(p *GeomLiteralContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_geomLiteral
}

func (*GeomLiteralContext) IsGeomLiteralContext() {}

func NewGeomLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeomLiteralContext {
	var p = new(GeomLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_geomLiteral

	return p
}

func (s *GeomLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *GeomLiteralContext) Point() IPointContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *GeomLiteralContext) Linestring() ILinestringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILinestringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILinestringContext)
}

func (s *GeomLiteralContext) Polygon() IPolygonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPolygonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPolygonContext)
}

func (s *GeomLiteralContext) MultiPoint() IMultiPointContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiPointContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiPointContext)
}

func (s *GeomLiteralContext) MultiLinestring() IMultiLinestringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiLinestringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiLinestringContext)
}

func (s *GeomLiteralContext) MultiPolygon() IMultiPolygonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiPolygonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiPolygonContext)
}

func (s *GeomLiteralContext) GeometryCollection() IGeometryCollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeometryCollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeometryCollectionContext)
}

func (s *GeomLiteralContext) Envelope() IEnvelopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvelopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvelopeContext)
}

func (s *GeomLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeomLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeomLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterGeomLiteral(s)
	}
}

func (s *GeomLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitGeomLiteral(s)
	}
}

func (p *CQLParser) GeomLiteral() (localctx IGeomLiteralContext) {
	localctx = NewGeomLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CQLParserRULE_geomLiteral)
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CQLParserPOINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Point()
		}

	case CQLParserLINESTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.Linestring()
		}

	case CQLParserPOLYGON:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(222)
			p.Polygon()
		}

	case CQLParserMULTIPOINT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(223)
			p.MultiPoint()
		}

	case CQLParserMULTILINESTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(224)
			p.MultiLinestring()
		}

	case CQLParserMULTIPOLYGON:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(225)
			p.MultiPolygon()
		}

	case CQLParserGEOMETRYCOLLECTION:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(226)
			p.GeometryCollection()
		}

	case CQLParserENVELOPE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(227)
			p.Envelope()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPointContext is an interface to support dynamic dispatch.
type IPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	POINT() antlr.TerminalNode
	PointList() IPointListContext

	// IsPointContext differentiates from other interfaces.
	IsPointContext()
}

type PointContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPointContext() *PointContext {
	var p = new(PointContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_point
	return p
}

func InitEmptyPointContext(p *PointContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_point
}

func (*PointContext) IsPointContext() {}

func NewPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointContext {
	var p = new(PointContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_point

	return p
}

func (s *PointContext) GetParser() antlr.Parser { return s.parser }

func (s *PointContext) POINT() antlr.TerminalNode {
	return s.GetToken(CQLParserPOINT, 0)
}

func (s *PointContext) PointList() IPointListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPointListContext)
}

func (s *PointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPoint(s)
	}
}

func (s *PointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPoint(s)
	}
}

func (p *CQLParser) Point() (localctx IPointContext) {
	localctx = NewPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CQLParserRULE_point)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(CQLParserPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.PointList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPointListContext is an interface to support dynamic dispatch.
type IPointListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFTPAREN() antlr.TerminalNode
	Coordinate() ICoordinateContext
	RIGHTPAREN() antlr.TerminalNode

	// IsPointListContext differentiates from other interfaces.
	IsPointListContext()
}

type PointListContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPointListContext() *PointListContext {
	var p = new(PointListContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_pointList
	return p
}

func InitEmptyPointListContext(p *PointListContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_pointList
}

func (*PointListContext) IsPointListContext() {}

func NewPointListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointListContext {
	var p = new(PointListContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_pointList

	return p
}

func (s *PointListContext) GetParser() antlr.Parser { return s.parser }

func (s *PointListContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *PointListContext) Coordinate() ICoordinateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoordinateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *PointListContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *PointListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPointList(s)
	}
}

func (s *PointListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPointList(s)
	}
}

func (p *CQLParser) PointList() (localctx IPointListContext) {
	localctx = NewPointListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CQLParserRULE_pointList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.Coordinate()
	}
	{
		p.SetState(235)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILinestringContext is an interface to support dynamic dispatch.
type ILinestringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LINESTRING() antlr.TerminalNode
	CoordList() ICoordListContext

	// IsLinestringContext differentiates from other interfaces.
	IsLinestringContext()
}

type LinestringContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyLinestringContext() *LinestringContext {
	var p = new(LinestringContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_linestring
	return p
}

func InitEmptyLinestringContext(p *LinestringContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_linestring
}

func (*LinestringContext) IsLinestringContext() {}

func NewLinestringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinestringContext {
	var p = new(LinestringContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_linestring

	return p
}

func (s *LinestringContext) GetParser() antlr.Parser { return s.parser }

func (s *LinestringContext) LINESTRING() antlr.TerminalNode {
	return s.GetToken(CQLParserLINESTRING, 0)
}

func (s *LinestringContext) CoordList() ICoordListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoordListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoordListContext)
}

func (s *LinestringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinestringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinestringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterLinestring(s)
	}
}

func (s *LinestringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitLinestring(s)
	}
}

func (p *CQLParser) Linestring() (localctx ILinestringContext) {
	localctx = NewLinestringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CQLParserRULE_linestring)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(CQLParserLINESTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.CoordList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPolygonContext is an interface to support dynamic dispatch.
type IPolygonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	POLYGON() antlr.TerminalNode
	PolygonDef() IPolygonDefContext

	// IsPolygonContext differentiates from other interfaces.
	IsPolygonContext()
}

type PolygonContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPolygonContext() *PolygonContext {
	var p = new(PolygonContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_polygon
	return p
}

func InitEmptyPolygonContext(p *PolygonContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_polygon
}

func (*PolygonContext) IsPolygonContext() {}

func NewPolygonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonContext {
	var p = new(PolygonContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_polygon

	return p
}

func (s *PolygonContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonContext) POLYGON() antlr.TerminalNode {
	return s.GetToken(CQLParserPOLYGON, 0)
}

func (s *PolygonContext) PolygonDef() IPolygonDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPolygonDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPolygonDefContext)
}

func (s *PolygonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPolygon(s)
	}
}

func (s *PolygonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPolygon(s)
	}
}

func (p *CQLParser) Polygon() (localctx IPolygonContext) {
	localctx = NewPolygonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CQLParserRULE_polygon)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(CQLParserPOLYGON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.PolygonDef()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPolygonDefContext is an interface to support dynamic dispatch.
type IPolygonDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFTPAREN() antlr.TerminalNode
	AllCoordList() []ICoordListContext
	CoordList(i int) ICoordListContext
	RIGHTPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPolygonDefContext differentiates from other interfaces.
	IsPolygonDefContext()
}

type PolygonDefContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPolygonDefContext() *PolygonDefContext {
	var p = new(PolygonDefContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_polygonDef
	return p
}

func InitEmptyPolygonDefContext(p *PolygonDefContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_polygonDef
}

func (*PolygonDefContext) IsPolygonDefContext() {}

func NewPolygonDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonDefContext {
	var p = new(PolygonDefContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_polygonDef

	return p
}

func (s *PolygonDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonDefContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *PolygonDefContext) AllCoordList() []ICoordListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICoordListContext); ok {
			len++
		}
	}

	tst := make([]ICoordListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICoordListContext); ok {
			tst[i] = t.(ICoordListContext)
			i++
		}
	}

	return tst
}

func (s *PolygonDefContext) CoordList(i int) ICoordListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoordListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoordListContext)
}

func (s *PolygonDefContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *PolygonDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *PolygonDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *PolygonDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPolygonDef(s)
	}
}

func (s *PolygonDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPolygonDef(s)
	}
}

func (p *CQLParser) PolygonDef() (localctx IPolygonDefContext) {
	localctx = NewPolygonDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CQLParserRULE_polygonDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.CoordList()
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(245)
			p.Match(CQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.CoordList()
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(252)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiPointContext is an interface to support dynamic dispatch.
type IMultiPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULTIPOINT() antlr.TerminalNode
	LEFTPAREN() antlr.TerminalNode
	AllPointList() []IPointListContext
	PointList(i int) IPointListContext
	RIGHTPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMultiPointContext differentiates from other interfaces.
	IsMultiPointContext()
}

type MultiPointContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyMultiPointContext() *MultiPointContext {
	var p = new(MultiPointContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_multiPoint
	return p
}

func InitEmptyMultiPointContext(p *MultiPointContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_multiPoint
}

func (*MultiPointContext) IsMultiPointContext() {}

func NewMultiPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPointContext {
	var p = new(MultiPointContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_multiPoint

	return p
}

func (s *MultiPointContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPointContext) MULTIPOINT() antlr.TerminalNode {
	return s.GetToken(CQLParserMULTIPOINT, 0)
}

func (s *MultiPointContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *MultiPointContext) AllPointList() []IPointListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPointListContext); ok {
			len++
		}
	}

	tst := make([]IPointListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPointListContext); ok {
			tst[i] = t.(IPointListContext)
			i++
		}
	}

	return tst
}

func (s *MultiPointContext) PointList(i int) IPointListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPointListContext)
}

func (s *MultiPointContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *MultiPointContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *MultiPointContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *MultiPointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterMultiPoint(s)
	}
}

func (s *MultiPointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitMultiPoint(s)
	}
}

func (p *CQLParser) MultiPoint() (localctx IMultiPointContext) {
	localctx = NewMultiPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CQLParserRULE_multiPoint)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(CQLParserMULTIPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.PointList()
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(257)
			p.Match(CQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.PointList()
		}

		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(264)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiLinestringContext is an interface to support dynamic dispatch.
type IMultiLinestringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULTILINESTRING() antlr.TerminalNode
	LEFTPAREN() antlr.TerminalNode
	AllCoordList() []ICoordListContext
	CoordList(i int) ICoordListContext
	RIGHTPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMultiLinestringContext differentiates from other interfaces.
	IsMultiLinestringContext()
}

type MultiLinestringContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyMultiLinestringContext() *MultiLinestringContext {
	var p = new(MultiLinestringContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_multiLinestring
	return p
}

func InitEmptyMultiLinestringContext(p *MultiLinestringContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_multiLinestring
}

func (*MultiLinestringContext) IsMultiLinestringContext() {}

func NewMultiLinestringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiLinestringContext {
	var p = new(MultiLinestringContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_multiLinestring

	return p
}

func (s *MultiLinestringContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiLinestringContext) MULTILINESTRING() antlr.TerminalNode {
	return s.GetToken(CQLParserMULTILINESTRING, 0)
}

func (s *MultiLinestringContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *MultiLinestringContext) AllCoordList() []ICoordListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICoordListContext); ok {
			len++
		}
	}

	tst := make([]ICoordListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICoordListContext); ok {
			tst[i] = t.(ICoordListContext)
			i++
		}
	}

	return tst
}

func (s *MultiLinestringContext) CoordList(i int) ICoordListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoordListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoordListContext)
}

func (s *MultiLinestringContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *MultiLinestringContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *MultiLinestringContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *MultiLinestringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiLinestringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiLinestringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterMultiLinestring(s)
	}
}

func (s *MultiLinestringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitMultiLinestring(s)
	}
}

func (p *CQLParser) MultiLinestring() (localctx IMultiLinestringContext) {
	localctx = NewMultiLinestringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CQLParserRULE_multiLinestring)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(CQLParserMULTILINESTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.CoordList()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(269)
			p.Match(CQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.CoordList()
		}

		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(276)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiPolygonContext is an interface to support dynamic dispatch.
type IMultiPolygonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULTIPOLYGON() antlr.TerminalNode
	LEFTPAREN() antlr.TerminalNode
	AllPolygonDef() []IPolygonDefContext
	PolygonDef(i int) IPolygonDefContext
	RIGHTPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMultiPolygonContext differentiates from other interfaces.
	IsMultiPolygonContext()
}

type MultiPolygonContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyMultiPolygonContext() *MultiPolygonContext {
	var p = new(MultiPolygonContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_multiPolygon
	return p
}

func InitEmptyMultiPolygonContext(p *MultiPolygonContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_multiPolygon
}

func (*MultiPolygonContext) IsMultiPolygonContext() {}

func NewMultiPolygonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPolygonContext {
	var p = new(MultiPolygonContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_multiPolygon

	return p
}

func (s *MultiPolygonContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPolygonContext) MULTIPOLYGON() antlr.TerminalNode {
	return s.GetToken(CQLParserMULTIPOLYGON, 0)
}

func (s *MultiPolygonContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *MultiPolygonContext) AllPolygonDef() []IPolygonDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPolygonDefContext); ok {
			len++
		}
	}

	tst := make([]IPolygonDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPolygonDefContext); ok {
			tst[i] = t.(IPolygonDefContext)
			i++
		}
	}

	return tst
}

func (s *MultiPolygonContext) PolygonDef(i int) IPolygonDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPolygonDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPolygonDefContext)
}

func (s *MultiPolygonContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *MultiPolygonContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *MultiPolygonContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *MultiPolygonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPolygonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPolygonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterMultiPolygon(s)
	}
}

func (s *MultiPolygonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitMultiPolygon(s)
	}
}

func (p *CQLParser) MultiPolygon() (localctx IMultiPolygonContext) {
	localctx = NewMultiPolygonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CQLParserRULE_multiPolygon)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(CQLParserMULTIPOLYGON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.PolygonDef()
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(281)
			p.Match(CQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.PolygonDef()
		}

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(288)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGeometryCollectionContext is an interface to support dynamic dispatch.
type IGeometryCollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GEOMETRYCOLLECTION() antlr.TerminalNode
	LEFTPAREN() antlr.TerminalNode
	AllGeomLiteral() []IGeomLiteralContext
	GeomLiteral(i int) IGeomLiteralContext
	RIGHTPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsGeometryCollectionContext differentiates from other interfaces.
	IsGeometryCollectionContext()
}

type GeometryCollectionContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyGeometryCollectionContext() *GeometryCollectionContext {
	var p = new(GeometryCollectionContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_geometryCollection
	return p
}

func InitEmptyGeometryCollectionContext(p *GeometryCollectionContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_geometryCollection
}

func (*GeometryCollectionContext) IsGeometryCollectionContext() {}

func NewGeometryCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeometryCollectionContext {
	var p = new(GeometryCollectionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_geometryCollection

	return p
}

func (s *GeometryCollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *GeometryCollectionContext) GEOMETRYCOLLECTION() antlr.TerminalNode {
	return s.GetToken(CQLParserGEOMETRYCOLLECTION, 0)
}

func (s *GeometryCollectionContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *GeometryCollectionContext) AllGeomLiteral() []IGeomLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGeomLiteralContext); ok {
			len++
		}
	}

	tst := make([]IGeomLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGeomLiteralContext); ok {
			tst[i] = t.(IGeomLiteralContext)
			i++
		}
	}

	return tst
}

func (s *GeometryCollectionContext) GeomLiteral(i int) IGeomLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeomLiteralContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeomLiteralContext)
}

func (s *GeometryCollectionContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *GeometryCollectionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *GeometryCollectionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *GeometryCollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeometryCollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeometryCollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterGeometryCollection(s)
	}
}

func (s *GeometryCollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitGeometryCollection(s)
	}
}

func (p *CQLParser) GeometryCollection() (localctx IGeometryCollectionContext) {
	localctx = NewGeometryCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CQLParserRULE_geometryCollection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(CQLParserGEOMETRYCOLLECTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.GeomLiteral()
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(293)
			p.Match(CQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.GeomLiteral()
		}

		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(300)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnvelopeContext is an interface to support dynamic dispatch.
type IEnvelopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENVELOPE() antlr.TerminalNode
	LEFTPAREN() antlr.TerminalNode
	AllNumericLiteral() []antlr.TerminalNode
	NumericLiteral(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	RIGHTPAREN() antlr.TerminalNode

	// IsEnvelopeContext differentiates from other interfaces.
	IsEnvelopeContext()
}

type EnvelopeContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyEnvelopeContext() *EnvelopeContext {
	var p = new(EnvelopeContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_envelope
	return p
}

func InitEmptyEnvelopeContext(p *EnvelopeContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_envelope
}

func (*EnvelopeContext) IsEnvelopeContext() {}

func NewEnvelopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvelopeContext {
	var p = new(EnvelopeContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_envelope

	return p
}

func (s *EnvelopeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvelopeContext) ENVELOPE() antlr.TerminalNode {
	return s.GetToken(CQLParserENVELOPE, 0)
}

func (s *EnvelopeContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *EnvelopeContext) AllNumericLiteral() []antlr.TerminalNode {
	return s.GetTokens(CQLParserNumericLiteral)
}

func (s *EnvelopeContext) NumericLiteral(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, i)
}

func (s *EnvelopeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *EnvelopeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *EnvelopeContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *EnvelopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvelopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvelopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterEnvelope(s)
	}
}

func (s *EnvelopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitEnvelope(s)
	}
}

func (p *CQLParser) Envelope() (localctx IEnvelopeContext) {
	localctx = NewEnvelopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CQLParserRULE_envelope)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(CQLParserENVELOPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.Match(CQLParserNumericLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Match(CQLParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(CQLParserNumericLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(CQLParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Match(CQLParserNumericLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.Match(CQLParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.Match(CQLParserNumericLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICoordListContext is an interface to support dynamic dispatch.
type ICoordListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFTPAREN() antlr.TerminalNode
	AllCoordinate() []ICoordinateContext
	Coordinate(i int) ICoordinateContext
	RIGHTPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCoordListContext differentiates from other interfaces.
	IsCoordListContext()
}

type CoordListContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCoordListContext() *CoordListContext {
	var p = new(CoordListContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_coordList
	return p
}

func InitEmptyCoordListContext(p *CoordListContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_coordList
}

func (*CoordListContext) IsCoordListContext() {}

func NewCoordListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordListContext {
	var p = new(CoordListContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_coordList

	return p
}

func (s *CoordListContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordListContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *CoordListContext) AllCoordinate() []ICoordinateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICoordinateContext); ok {
			len++
		}
	}

	tst := make([]ICoordinateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICoordinateContext); ok {
			tst[i] = t.(ICoordinateContext)
			i++
		}
	}

	return tst
}

func (s *CoordListContext) Coordinate(i int) ICoordinateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoordinateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *CoordListContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *CoordListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *CoordListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *CoordListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoordListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCoordList(s)
	}
}

func (s *CoordListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCoordList(s)
	}
}

func (p *CQLParser) CoordList() (localctx ICoordListContext) {
	localctx = NewCoordListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CQLParserRULE_coordList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Match(CQLParserLEFTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Coordinate()
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(315)
			p.Match(CQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Coordinate()
		}

		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(322)
		p.Match(CQLParserRIGHTPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICoordinateContext is an interface to support dynamic dispatch.
type ICoordinateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumericLiteral() []antlr.TerminalNode
	NumericLiteral(i int) antlr.TerminalNode

	// IsCoordinateContext differentiates from other interfaces.
	IsCoordinateContext()
}

type CoordinateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCoordinateContext() *CoordinateContext {
	var p = new(CoordinateContext)
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_coordinate
	return p
}

func InitEmptyCoordinateContext(p *CoordinateContext) {
	p.CqlContext = NewCqlContext(nil, -1) // Jim super
	p.RuleIndex = CQLParserRULE_coordinate
}

func (*CoordinateContext) IsCoordinateContext() {}

func NewCoordinateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordinateContext {
	var p = new(CoordinateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = CQLParserRULE_coordinate

	return p
}

func (s *CoordinateContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordinateContext) AllNumericLiteral() []antlr.TerminalNode {
	return s.GetTokens(CQLParserNumericLiteral)
}

func (s *CoordinateContext) NumericLiteral(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, i)
}

func (s *CoordinateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoordinateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCoordinate(s)
	}
}

func (s *CoordinateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCoordinate(s)
	}
}

func (p *CQLParser) Coordinate() (localctx ICoordinateContext) {
	localctx = NewCoordinateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CQLParserRULE_coordinate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(CQLParserNumericLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.Match(CQLParserNumericLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *CQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *BooleanValueExpressionContext = nil
		if localctx != nil {
			t = localctx.(*BooleanValueExpressionContext)
		}
		return p.BooleanValueExpression_Sempred(t, predIndex)

	case 2:
		var t *BooleanTermContext = nil
		if localctx != nil {
			t = localctx.(*BooleanTermContext)
		}
		return p.BooleanTerm_Sempred(t, predIndex)

	case 11:
		var t *ScalarExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ScalarExpressionContext)
		}
		return p.ScalarExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CQLParser) BooleanValueExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CQLParser) BooleanTerm_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CQLParser) ScalarExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
