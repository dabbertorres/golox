// Code generated by "stringer -type=TokenType"; DO NOT EDIT.

package lox

import "fmt"

const _TokenType_name = "UnknownLeftParenRightParenLeftBraceRightBraceCommaDotMinusPlusSemicolonSlashStarBangBangEqualEqualEqualEqualGreaterGreaterEqualLessLessEqualIdentifierStringNumberAndClassElseFalseFunForIfNilOrPrintReturnSuperThisTrueVarWhileEOF"

var _TokenType_index = [...]uint8{0, 7, 16, 26, 35, 45, 50, 53, 58, 62, 71, 76, 80, 84, 93, 98, 108, 115, 127, 131, 140, 150, 156, 162, 165, 170, 174, 179, 182, 185, 187, 190, 192, 197, 203, 208, 212, 216, 219, 224, 227}

func (i TokenType) String() string {
	if i >= TokenType(len(_TokenType_index)-1) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}