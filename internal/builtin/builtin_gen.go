// This file was generated by codegen-builtins, DO NOT EDIT!
// See: /tools/codegen-builtins
//
// Source: ../../tools/gendata/builtin.go.txt
// Skipped: ComplexType,Type,Type1,IntegerType,FloatType

package builtin

import "go.lsp.dev/protocol"

// buckets contains list of completions for builtin Gno functions grouped by a first letter.
var buckets = map[rune][]protocol.CompletionItem{
	int32(100): []protocol.CompletionItem{protocol.CompletionItem{
		Detail: "func delete(m map[Type]Type1, key Type)",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The delete built-in function deletes the element with the specified key (m\\[key]) from the map. If m is nil or there is no such element, delete is a no-op.",
		},
		InsertText:       "delete()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "delete",
	}},
	int32(101): []protocol.CompletionItem{protocol.CompletionItem{
		Detail: "type error interface {\n\tError() string\n}",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The error built-in interface type is the conventional interface for representing an error condition, with the nil value representing no error.",
		},
		InsertText:       "error",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindInterface,
		Label:            "error",
	}},
	int32(102): []protocol.CompletionItem{protocol.CompletionItem{
		InsertText:       "false",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindConstant,
		Label:            "false",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "float32 is the set of all IEEE-754 32-bit floating-point numbers.",
		},
		InsertText:       "float32",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "float32",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "float64 is the set of all IEEE-754 64-bit floating-point numbers.",
		},
		InsertText:       "float64",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "float64",
	}},
	int32(105): []protocol.CompletionItem{protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "int8 is the set of all signed 8-bit integers. Range: -128 through 127.",
		},
		InsertText:       "int8",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "int8",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "int16 is the set of all signed 16-bit integers. Range: -32768 through 32767.",
		},
		InsertText:       "int16",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "int16",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "int32 is the set of all signed 32-bit integers. Range: -2147483648 through 2147483647.",
		},
		InsertText:       "int32",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "int32",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "int64 is the set of all signed 64-bit integers. Range: -9223372036854775808 through 9223372036854775807.",
		},
		InsertText:       "int64",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "int64",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "int is a signed integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, int32.",
		},
		InsertText:       "int",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "int",
	}, protocol.CompletionItem{
		Detail:           "const iota = 0",
		InsertText:       "iota",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindConstant,
		Label:            "iota",
	}},
	int32(108): []protocol.CompletionItem{protocol.CompletionItem{
		Detail: "func len(v Type) int",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The len built-in function returns the length of v, according to its type:\n\n\tArray: the number of elements in v.\n\tPointer to array: the number of elements in *v (even if v is nil).\n\tSlice, or map: the number of elements in v; if v is nil, len(v) is zero.\n\tString: the number of bytes in v.\n\nFor some arguments, such as a string literal or a simple array expression, the result can be a constant.",
		},
		InsertText:       "len()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "len",
	}},
	int32(109): []protocol.CompletionItem{protocol.CompletionItem{
		Detail: "func make(t Type, size ...IntegerType) Type",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. The specification of the result depends on the type:\n\n\tSlice: The size specifies the length. The capacity of the slice is\n\tequal to its length. A second integer argument may be provided to\n\tspecify a different capacity; it must be no smaller than the\n\tlength. For example, make([]int, 0, 10) allocates an underlying array\n\tof size 10 and returns a slice of length 0 and capacity 10 that is\n\tbacked by this underlying array.\n\tMap: An empty map is allocated with enough space to hold the\n\tspecified number of elements. The size may be omitted, in which case\n\ta small starting size is allocated.",
		},
		InsertText:       "make()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "make",
	}},
	int32(110): []protocol.CompletionItem{protocol.CompletionItem{
		Detail: "func new(Type) *Type",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The new built-in function allocates memory. The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.",
		},
		InsertText:       "new()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "new",
	}, protocol.CompletionItem{
		Detail:           "var nil Type",
		InsertText:       "nil",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindVariable,
		Label:            "nil",
	}},
	int32(112): []protocol.CompletionItem{protocol.CompletionItem{
		Detail: "func panic(v interface{})",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The panic built-in function stops normal execution of the program. Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller.\n\nAt that point, the program is terminated with a non-zero exit code. This termination sequence is called panicking and can be controlled by the built-in function recover.",
		},
		InsertText:       "panic()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "panic",
	}, protocol.CompletionItem{
		Detail: "func print(args ...Type)",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The print built-in function formats its arguments in an implementation-specific way and writes the result to standard error. Print is useful for bootstrapping and debugging; it is not guaranteed to stay in the language.",
		},
		InsertText:       "print()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "print",
	}, protocol.CompletionItem{
		Detail: "func println(args ...Type)",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The println built-in function formats its arguments in an implementation-specific way and writes the result to standard error. Spaces are always added between arguments and a newline is appended. Println is useful for bootstrapping and debugging; it is not guaranteed to stay in the language.",
		},
		InsertText:       "println()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "println",
	}},
	int32(114): []protocol.CompletionItem{protocol.CompletionItem{
		Detail: "func recover() interface{}",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The recover built-in function allows a program to manage behavior of a panicking goroutine. Executing a call to recover inside a deferred function (but not any function called by it) stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic. If recover is called outside the deferred function it will not stop a panicking sequence. In this case, or when the program is not panicking, recover returns nil.",
		},
		InsertText:       "recover()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "recover",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.",
		},
		InsertText:       "rune",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "rune",
	}},
	int32(115): []protocol.CompletionItem{protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "string is the set of all strings of 8-bit bytes, conventionally but not necessarily representing UTF-8-encoded text. A string may be empty, but not nil. Values of string type are immutable.",
		},
		InsertText:       "string",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "string",
	}},
	int32(116): []protocol.CompletionItem{protocol.CompletionItem{
		InsertText:       "true",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindConstant,
		Label:            "true",
	}},
	int32(117): []protocol.CompletionItem{protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255.",
		},
		InsertText:       "uint8",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "uint8",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "uint16 is the set of all unsigned 16-bit integers. Range: 0 through 65535.",
		},
		InsertText:       "uint16",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "uint16",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "uint32 is the set of all unsigned 32-bit integers. Range: 0 through 4294967295.",
		},
		InsertText:       "uint32",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "uint32",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "uint64 is the set of all unsigned 64-bit integers. Range: 0 through 18446744073709551615.",
		},
		InsertText:       "uint64",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "uint64",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "uint is an unsigned integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, uint32.",
		},
		InsertText:       "uint",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "uint",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "uintptr is an integer type that is large enough to hold the bit pattern of any pointer.",
		},
		InsertText:       "uintptr",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "uintptr",
	}},
	int32(97): []protocol.CompletionItem{protocol.CompletionItem{
		Detail: "func append(slice []Type, elems ...Type) []Type",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself:\n\n\tslice = append(slice, elem1, elem2)\n\tslice = append(slice, anotherSlice...)\n\nAs a special case, it is legal to append a string to a byte slice, like this:\n\n\tslice = append([]byte(\"hello \"), \"world\"...)",
		},
		InsertText:       "append()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "append",
	}},
	int32(98): []protocol.CompletionItem{protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "bool is the set of boolean values, true and false.",
		},
		InsertText:       "bool",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "bool",
	}, protocol.CompletionItem{
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "byte is an alias for uint8 and is equivalent to uint8 in all ways. It is used, by convention, to distinguish byte values from 8-bit unsigned integer values.",
		},
		InsertText:       "byte",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindClass,
		Label:            "byte",
	}},
	int32(99): []protocol.CompletionItem{protocol.CompletionItem{
		Detail: "func copy(dst, src []Type) int",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The copy built-in function copies elements from a source slice into a destination slice. (As a special case, it also will copy bytes from a string to a slice of bytes.) The source and destination may overlap. Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).",
		},
		InsertText:       "copy()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "copy",
	}, protocol.CompletionItem{
		Detail: "func cap(v Type) int",
		Documentation: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: "The cap built-in function returns the capacity of v, according to its type:\n\n\tArray: the number of elements in v (same as len(v)).\n\tPointer to array: the number of elements in *v (same as len(v)).\n\tSlice: the maximum length the slice can reach when resliced;\n\tif v is nil, cap(v) is zero.\n\nFor some arguments, such as a simple array expression, the result can be a constant.",
		},
		InsertText:       "cap()",
		InsertTextFormat: protocol.InsertTextFormatPlainText,
		Kind:             protocol.CompletionItemKindFunction,
		Label:            "cap",
	}},
}