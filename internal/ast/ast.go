package ast 

type Node interface {
  children []Node
}


// --------------------- EXPRESSIONS ---------------------------------------------
// All expression nodes implement the Expr interface
// Example for statement nodes: call expression   
type Expression interface {
	Node
	exprNode()
}

type BinaryOperationExpression struct {
  Left      Node
  Rigth     Node
  Operator String
} 
// BINARY OPERATION EXPRESSION EXAMPLES:
// ADD: a + b
// SUB: a - b
// MUL: a * b
// SUB: a / b
// MOD: a % b
// AND: a & b
// OR: a | b
// XOR: a ^ b
// AND: a && b
// OR: a || b
// Equality: a == b
// Inequality: a != b
// Greater than: a > b
// Less than: a < b
// Greater than or equal to: a >= b
// Less than or equal to: a <= b
  
type UnaryOperationExpression  struct {
  Operator String
  Operand Node
}
// UNARY OPERATION EXPRESSION EXAMPLES
// Negation: -a
// Logical NOT: !a
// Bitwise NOT: ~a

type FunctionCallExpression struct {
	FunctionName string
	Arguments    []Node
}

type ConditionalExpression struct { // ( contition ? trueExpression : falseExpression )
  Condition Node
  TrueExpression Node
  FalseExpression Node
}

type AssignmentExpression struct { // variable = value
	Variable string
	Value    Node
}

type VariableReferenceExpression struct { // variableName
  VariableName string
}

type LiteralExpression struct {
  Tok token.Token
}
// LITERAL EXPRESSION EXAMPLES:
// Numeric literals: 42, 3.14
// String literals: "hello", 'world'
// Boolean literals: true, false
// Null literal: null


// --------------------- STATEMENTS ---------------------------------------------
// All statement nodes implement the Stmt interface
// Example for statement nodes: if statement,  
type Statement interface {
	Node
	stmtNode()
}



// --------------------- DECLARATIONS ---------------------------------------------
// All declaration nodes implement the Decl interface
// Example for declaration nodes: variable declaration, function declaration 
type Declaration interface {
	Node
	declNode()
}

// --------------------- LITERALS ---------------------------------------------
// All Literal nodes implement the Literal interface
// Example for literal nodes:  
type Literal interface {
	Node
	litNode()
}


