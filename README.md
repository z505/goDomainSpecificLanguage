# goDomainSpecificLanguage
domain specific language in Go (possible abuse)

Use a variadic function to embed a mini programming language inside go, called JOKE LANG or Joke Language

Just a small sample, not complete, and is a silly little language.

This is what the language looks like, which a Go compiler will compile as go code

```

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

```

Warnings: 
* avoid tricks like the plague  
* DSL = drug specific language. 
* Lisp = troublesome
