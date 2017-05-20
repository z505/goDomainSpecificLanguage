// Embed a toy language (jokeLang) DSL inside go code, however the DSL has some
// disadvantages like having to use quotes around certain syntax, and commas.
// However a text editor could easily convert snippets of text to the correct
// syntax if programmer got tired of comma/quote syntax for the DSL.
// How it works:
// variadic function accepts a list of parameters, these parameters are the
// programming language. But it's much more than that, as one can embed the DSL
// and mix Go code using a go func, and the line between interpreter and
// compiler becomes blurry. This code is to be considered language abuse and
// may cause brain damage.
//
// jokeLang is an incomplete non finished example of a DSL, one could create
// hundreds of other languages. However English tokens seem better suited
// than special characters that need to be quoted, so algol/cobol style
// languages may be easier to embed than others if avoiding using too many
// quotations. However since one can convert any snipped of code to have
// quotations around certain identifiers, really any language could be
// embedded, sometimes painful to use without an automatic "add the quotes"
// converter.
// NO UNICODE allowed in DSL.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

const IS_GO_CODE = `THIS_IS_GOLANG_CODE`

// Tokens for DSL
const START = `TOKEN_START`
const FINISH = `TOKEN_FINISH`
const RUN = `TOKEN_RUN`
const PROC = `TOKEN_PROC`
const DONE = `TOKEN_DONE`
const CALL_GOLANG_FUNC = `TOKEN_CALL_GOLANG_FUNC`
const VAR = `TOKEN_VAR`
const dslFinished = `TOKEN_dslFinished`

const dslCompErr = `DSL compilinterptr error: `

// operators
const ASSIGN = `<--`
// (there is no plus sign, minus, multiply, because it's a toy language)

type VarInfo struct {
	name string
	typ string
}

/*
type statementKind int

const (
	skPROC statementKind = 1 + iota   // procedure
	skASSIGN  // variable assignment
)
*/

type Assignment struct {
	variable string
	value string
}

type Statement struct {
	line string
}


func errorStatus(s string) {
	fmt.Println(s)
}

func dslTest(lang ...interface{}) {
var (
	i int
	s string
	isInt bool
	isString bool
)

    // Deal with parameters
    for idx,p := range lang {

        s, isString = p.(string)
		if isString {
			fmt.Println(s)
		}
		fmt.Println(strconv.Itoa(i))

        i, isInt = p.(int)
		if isInt {
			fmt.Println(strconv.Itoa(i))
		}
		fmt.Println(strconv.Itoa(idx))

	}

}

// returns whether or not procedure header syntax was correct
// returns proc name as a string
func eatProcHeader(s string) (ok bool, procName string) {
var i int
	ok = false
	foundBracket := false
	loop:
	for i < len(s) {
		switch string(s[i]) {
		case `(`:
			foundBracket = true
			// as it is a toy language, only ExampleProc() style custom procedures are allowed
			// not ExampleProc(1,"string", 2, 3) with parameters. That could be implemented later.
			if string(s[i+1]) == `)` {
				ok = true
				break loop
			} else {
				errorStatus(dslCompErr+` problem in procedure syntax, missing closing ')'`)
				return
			}
		default:
			// concatenate string result (procedure name)
			procName += string(s[i])
		}
		i++
	}
	if !foundBracket {
		errorStatus(`Error in procedure syntax, missing opening '('`)
	}
	return
}

func parseCode() {

}

// returns whether or not "string" (sub) exists in "stringhere" (s) at position zero
func leftStr(sub, s string) (found bool) {
	found = strings.Index(s, sub) == 0
	return
}

// the only two internal functions available in jokeLang, exposed to use by DSL
const INTERN_ADD = `add(`
const INTERN_SUBTRACT = `subtract(`

// returns function name and whether it is a valid internal function
func isInternalFunc(item interface{}) (string, bool) {
	s, isString := item.(string)
	if isString {
		// add(x,y) function
		if leftStr(INTERN_ADD, s) {
			return INTERN_ADD, true
		}
		// subtract(x,y) function
		if leftStr(INTERN_SUBTRACT, s) {
			return INTERN_SUBTRACT, true
		}
	}
	return "", false
}

// todo: type check, place internal add/subtract function inline,
// parse function parameters to add/subtract
func parseAssignVal(item interface{}, typ string) (string, bool /* success */) {
//	s, isString := item.(string)
//	if isString {
		// determine if it's an internal standard function handled by go
		// TODO
		// funcName, is := isInternalFunc(s)
//		if is {
		    // TODO
			// params := getFuncParams(funcName, s)
			// TODO:
			// callInternalFunc(funcName, )
//		} else { // determine if it's a "string" or other value
//
//		}
//	}

	// TODO
	// i, isInt := item.(int)
	// if isInt {
		// return , true
	// }

	return "NOT IMPLEMENTED YET", false
}

// find variable in slice, return variable's type and if it was found or not
func varDeclared(varName string, vars []VarInfo) (typ string, found bool) {
	found = false
	for i := range vars {
		if vars[i].name == varName {
			typ = vars[i].typ
			found = true
			break
        }
    }
    return typ, found
}

//
// idx is the position in variadic function to start parsing at
// lang is the parameters from variadic function
func eatProcBody(idx int, vars []VarInfo, lang []interface{}) (ok bool, codes []Statement) {
	ok = false
	i := idx + 1 // string after VAR is the variable name and type
	if i < 0 || i >= len(lang) {
		errorStatus(dslCompErr+`eatProcBody function index out of bounds`)
		return
	}
	for {
		param := lang[i]
		s, isString := param.(string)
		if !isString {
			errorStatus(dslCompErr+`first code statement (or item on left of assignment) must be a string, variadic param: ` + strconv.Itoa(i))
			break
		}
		// reached the end of code block
		if s == FINISH {
			ok = true
			break
		}
		up1 := i+1
//		nextParam := lang[up1]
		nextStr, isString := param.(string)
		// TODO
		var parsed bool
		var val string
		if nextStr == ASSIGN {
			varName := s
			typ, declared := varDeclared(varName, vars)
			if declared {
				// TODO
				val, parsed = parseAssignVal(lang[up1], typ) // parse token to the right of assignment op
			} else {
				errorStatus(dslCompErr+`variable not declared before use: `+varName)
				break
			}
		}

		if parsed {
			// add code to slice TODO
			// codes = append(codes, )
    		i++ // skip val to right of assignment
		} else {
			errorStatus(dslCompErr+`couldn't parse code statement, variadic param: ` + strconv.Itoa(i) + ` '`+s+`'`+" Val: "+val)
		}
		i++
	}
	return
}


func status(msg string) {
	fmt.Println(`Status: ` + msg)
}


func registerProc(name string) {
	status(`Procedure registered: ` + name)
}

func registerLocalVars(vars []VarInfo) {
    for _,v := range vars {
    	status(`Variable registered: ` + v.name + ` with a type of: ` + v.typ)
    }
    status(`Registered ` + strconv.Itoa(len(vars)) + ` total local vars`)
}

// parses var declaration i.e. "x int" and returns VarInfo , variable name and type, split by space
// returns false if problem parsing
func parseVar(s string) (v VarInfo, ok bool) {
	ok = false
	if len(s) < 1 {
		return
	}

	gotVarName := false
	i := 0
	for {
		// todo: only accept a..Z, 0..9  characters and space, report error if any other invalid char
		// fmt.Println(`DEBUG: i: `+strconv.Itoa(i))
		// fmt.Println(`DEBUG: s: `+s)
		// space found, this splits var name from type
		if string(s[i]) == ` ` && len(v.name) > 0 {
        	gotVarName = true
        	i++ // skip past space!
        	// check if done
        	if i >= len(s) {
        		errorStatus(dslCompErr+`no type defined after VAR name`)
            	break
        	}
		}

 		// concat char by char until we find space, extract var name before space
		if !gotVarName {
   			v.name += string(s[i])
		} else { // extract type, after space
			// no spaces allowed in types
			if string(s[i]) == ` ` {
				errorStatus(dslCompErr+`no spaces allowed in var type specification`)
            	break
			} else { // concat type
				v.typ += string(s[i])
			}
		}

		i++
		if i >= len(s) {
			if len(v.typ) > 0 {
				ok = true
			}
			break
		}
	}
	if !gotVarName { // if var name was not obtained, return default not ok
		return
	}
	// ...
	return
}

// find all variable declarations in procedure up until START of code block (where var declarations end)
// sets ok to false, if end of the var declarations not found
// idx is passed in which is the DSL variadic param number the parser is currently on, so CPU is not wasted reparsing from position zero
// idx is similar to a pointer to the parser position
func eatVars(idx int, lang []interface{}) (ok bool, vars []VarInfo) {
	ok = false
	i := idx + 1 // string after VAR is the variable name and type
	if i < 0 || i >= len(lang) {
		errorStatus(dslCompErr+`eatVars function index out of bounds`)
		return
	}
	for {
    	param := lang[i]
        s, isString := param.(string)
		if !isString {
			errorStatus(dslCompErr+`variable declaration must be a string, variadic param: ` + strconv.Itoa(i))
			break
		}
		// reached the START reserved word, end of VAR declarations
		if s == START && len(vars) > 0 {
			ok = true
			break
		}

		v, parsed := parseVar(s)

		if parsed {
			// add var to slice
    		vars = append(vars, v)
		} else {
			errorStatus(dslCompErr+`couldn't parse var declarations, variadic param: ` + strconv.Itoa(i) + ` '`+s+`'`)
		}
		i++
	}
	return
}


type goFunc func()

func callGo(f goFunc) goFunc /* yourself */ {
	return f
}

func someGoLangFunc() {
	fmt.Println(`This is a golang function (procedure) being called`)
	// do whatever code you want here, side effects galore..
	if 2 + 2 == 5 {
		fmt.Println(`If this line doesn't print, return your computer to the ` +
					`store immediately for a refund after January 1, 1984`)
	}
}

// TODO
func processCode(code []Statement) (ok bool) {
	return
}

// domain specific language variadic function trick for GoLang.
// parses variadic function paramaters as a programming language of tokens, etc.
//								"Avoid tricks like the plague"
//								"DSL's may cause brain damage"
//								"Drug Specific Language"
//								"Lisp. No."
func dsl(lang ...interface{}) {
var s string
// var i int
var (
	vars []VarInfo
	isInt bool
	isString bool
	// what zone are we inside, status flag
	inProcHeader = false
	inProcBody = false
	doneProcHeader = false
	inVar = false
	procName string
	procBodyOk = false
	varsOk = false
	code []Statement
)

	idx := 0

	// parse the dsl language
	loop:
	for {
    	if idx >= len(lang) {
			break
		}
    	param := lang[idx]
        s, isString = param.(string)
        ok := false
		if isString {
			switch s {
			case START:
				inProcBody = true
				if !doneProcHeader {
					errorStatus(dslCompErr+`you have not defined a procedure header before code block START`)
                	break loop
				}
				inVar = false
				procBodyOk, code = eatProcBody(idx, vars, lang)

				if procBodyOk {
					processCode(code)
				} else {
					errorStatus(dslCompErr+`problem parsing PROC body`)
				}
            case FINISH:
				inProcBody = false
				vars = []VarInfo{} // clear variable data at end of procedure code block
			case RUN:
			//
			case PROC:
				up1 := idx+1 // next item to parse after PROC reserved word
				if up1 <= len(lang) {
					param := lang[up1]
					s, isString = param.(string)
			    } else {
					errorStatus(dslCompErr+`problem parsing PROC, no text after`)
					break loop
			    }

				ok, procName = eatProcHeader(s)
				if ok {
					registerProc(procName)
					doneProcHeader = true
					inProcHeader = false
				}
			case DONE:
			//
			case VAR:
				if doneProcHeader {
					inVar = true
					varsOk, vars = eatVars(idx, lang)
					if varsOk {
                		registerLocalVars(vars)
                		inVar = false
                	} else {
	                	errorStatus(dslCompErr+`problem extracting vars`)
						break loop
                	}
				}

			case dslFinished:

			default:
				if inProcHeader {
				}
				if inProcBody {
				}
				if inVar {
				}
				errorStatus(dslCompErr+`invalid token: `+s)

			}
		}

        // i, isInt = param.(int)
		if isInt {
        	// switch i {
        	// case
        	//}
		}

		// call GoLang code from within DSL code section using more tricks
		fthis, isGoFunc := param.(goFunc)
		if isGoFunc {
			fthis()
		}


		idx++
		ok = false
	}
}

func goLangExample(x,y,z int) (dummy string) {

	// now write your code
	fmt.Println(x, ": x value ", y,": y value ", z,": z value ")
	if `42` == `true` {
		fmt.Println("Life is pointful")
	} else {
		fmt.Println("Life is senseless")
	}
	// you must return this function result to inform that your function is go code
	return IS_GO_CODE
}


func goHello () (dummy string) {
	// now write your code
	fmt.Println("hello")
	// ...

	// you must return this function result
	return IS_GO_CODE
}

// THIS is how you embed a Domain Specific language inside GoLang causing
// potential Lisp like brain damage

func jokeLang() {
dsl(
	PROC,`Example()`,
	VAR,`x int`,
	    `y int`,
	    `s string`,
	START,
		`x`,`<--`,50,
		`y`,`<--`,100,
		`total`,`<--`,`add(x, y)`,
		`stdout(total)`,
		`s`,`<--`,`s became a string`,
	FINISH,

	PROC,`AnotherExample()`,
	VAR, `a int`,
	     `b int`,
	START,
		`a`,`<--`,10,
		`b`,`<--`,20,
		`total`,`<--`,`add(a, b)`,
		`stdout(total)`,
	FINISH,

	RUN,

	  `Example()`,

	  `# this is a comment`,
	  `# Or an inside joke`,
	  `# My voice does not have a Lispppft.`,

	  `# Let's run the example procedure twice more`,

      `Example()`,
	  `Example()`,

	  `# Comment: it should therefore print the example 3 times total`,

	  `# now let's call another example procedure`,
	  `AnotherExample()`,

	   // but what about if we want to call some GoLang Code?
	   // this calls someGoLangFunc() inside the DSL run section we are in
	   callGo(someGoLangFunc),
	   // there, we just mixed DSL code with Go Code calls

	  `# Now we're back to the DSL code again.. call the example one more time...`,
	  `Example()`,

	 DONE,


dslFinished)
}

func main() {
	jokeLang()
}
