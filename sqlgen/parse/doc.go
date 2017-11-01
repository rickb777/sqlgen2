// Package parse handles the traversing of the Go source code abstracat syntax tree
// to derive a Node tree, which is in turn converted elsewhere to a schema.Table
// and []schema.Field.
package parse

// https://golang.org/ref/spec

// Declaration   = ConstDecl | TypeDecl | VarDecl .
// TopLevelDecl  = Declaration | FunctionDecl | MethodDecl .

// SourceFile       = PackageClause ";" { ImportDecl ";" } { TopLevelDecl ";" } .

// PackageClause  = "package" PackageName .
// PackageName    = identifier .

// ImportDecl       = "import" ( ImportSpec | "(" { ImportSpec ";" } ")" ) .
// ImportSpec       = [ "." | PackageName ] ImportPath .
// ImportPath       = string_lit .

// Type      = TypeName | TypeLit | "(" Type ")" .
// TypeName  = identifier | QualifiedIdent .
// TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType | SliceType | MapType | ChannelType .

// ArrayType   = "[" ArrayLength "]" ElementType .
// ArrayLength = Expression .
// ElementType = Type .

// SliceType = "[" "]" ElementType .

// StructType    = "struct" "{" { FieldDecl ";" } "}" .
// FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
// EmbeddedField = [ "*" ] TypeName .
// Tag           = string_lit .

// IdentifierList = identifier { "," identifier } .

// PointerType = "*" BaseType .
// BaseType    = Type .

// MapType     = "map" "[" KeyType "]" ElementType .
// KeyType     = Type .

// FunctionType not relevant here.

// InterfaceType not relevant here.

// ChannelType not relevant here.

// TypeDecl = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
// TypeSpec = AliasDecl | TypeDef .

// AliasDecl = identifier "=" Type .

// TypeDef = identifier Type .

// Operand     = Literal | OperandName | MethodExpr | "(" Expression ")" .
// Literal     = BasicLit | CompositeLit | FunctionLit .
// BasicLit    = int_lit | float_lit | imaginary_lit | rune_lit | string_lit .
// OperandName = identifier | QualifiedIdent.

// QualifiedIdent = PackageName "." identifier .
