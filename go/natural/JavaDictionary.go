/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package natural

import (
	"github.com/emirpasic/gods/sets/hashset"
)

var (
	JavaDictionary = []string{
		"abstract",
		"continue",
		"for",
		"new",
		"switch",
		"assert",
		"default",
		"goto",
		"package",
		"synchronized",
		"boolean",
		"do",
		"if",
		"private",
		"this",
		"break",
		"double",
		"implements",
		"protected",
		"throw",
		"byte",
		"else",
		"import",
		"public",
		"throws",
		"case",
		"enum",
		"instanceof",
		"return",
		"transient",
		"catch",
		"extends",
		"int",
		"short",
		"try",
		"char",
		"final",
		"interface",
		"static",
		"void",
		"class",
		"finally",
		"long",
		"strictfp",
		"volatile",
		"const",
		"float",
		"native",
		"super",
		"while",
		//new
		"done",
		"properties",
		"entity",
		"param",
		"generic",
		"debug",
		"integer",
		"html",
		"print",
		"http",
		"jdbc",
		"odbc",
		"close",
		"stack",
		"system",
		"object",
		"false",
		"true",
		"locale",
		"decimal",
		"util",
		"context",
		"user",
		"login",
		"list",
		"link",
		"linked",
		"argument",
		"builder",
		"append",
		"temp",
		"result",
		"results",
		"hash",
		"error",
		"errors",
		"message",
		"messages",
		"value",
		"size",
		"module",
		"modules",
		"success",
		"warn",
		"warning",
		"stored",
		"modified",
		"node",
		"tree",
		"root",
	}
)

func GetJavaDictionary() hashset.Set {

	dictSet := hashset.New()

	//--- scan words into dictionary slice
	for _, entry := range JavaDictionary {
		dictSet.Add(entry)
	}

	return *dictSet
}
