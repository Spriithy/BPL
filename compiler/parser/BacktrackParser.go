package parser

/*
	Let us try to parse this grammar, using a Backtracking parser:

	stat
		: expr ';'
		;
	binop : '+' | '-' | '*' | '/' | '^' | '++' | '--' ;
	unaryop : '-' | '++' | '--' ;
	expr
		: assign
		| unaryop expr
		| expr binop expr
		| '(' expr ')'
		| INT
		| ID
		;
	assign
		: ID '=' expr
		| '(' assign ')'
		;
 */

type BkTParser struct {
	
}
