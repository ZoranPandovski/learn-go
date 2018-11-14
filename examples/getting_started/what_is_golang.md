# What is Golang ?
Go language is a programming language initially developed at Google in the year 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. It is a statically-typed language having syntax similar to that of C. It provides garbage collection, type safety, dynamic-typing capability, many advanced built-in types such as variable length arrays and key-value maps. It also provides a rich standard library. The Go programming language was launched in November 2009 and is used in some of the Google's production systems.

## Some advantages of using Go:

#### Compiled
* There is no VM. It compiles directly to the machine code (if we exclude Go’s intermediary assembly) which is fast, fast and fast (did I say fast?).
* Fast compilation. The programming language design is built for fast compilation in mind from the beginning.
* Compiles cross-platform to OS X, Linux, Windows, 👉 and many others.
* Creates only one executable file output after the compilation without any dependencies, so that you can upload it anywhere which Go supports and just run it. Or just compile it there after you upload the code. No dependency hell.
#### Safe
* Strong and static typed.
* Garbage collected. It cleans up your dirt after you and integrates the whole garbage collection system into your executable binary.
* Reliable. You can really create a very reliable software with Go. Because the inherent language design prevents you from doing awful stuff with it. 👉 For example: It has pointers but they’re mostly not dangerous as in C because the memory is being managed by Go and pointer arithmetic is not advised by default.
👉However, this reliability is only for the compilation part, in runtime, nasty things can happen, if you want maximum run-time reliability, you may prefer, for example, Rust instead.
#### Paradigms
* It’s an imperative language which is an advantage and a disadvantage to some people.
* Supports a different kind of object-oriented programming (OOP). I come from many OOP languages like Java, C#, Ruby, but, Go gets the best practices from OOP and lets you program differently, in a Go way.
* Go wants you to compose things not inherit like in other OOP langs.
* Supports interfaces (as in OOP). This helps to compose things. Does “Polymorphism” ring a bell? Go does not force you to mark your types as “implements this and that interface”, it inferres this from the functionality the type supports. This increases flexibility and composability.
* Go lets you attach functions to any type. This flexibility lets you compose your programs from smaller things. When a type implements the functions of an interface it’s said that the type satisfies that interface and can be used in places that desires that interface.
* Supports functional programming (FP). For example, Go supports anonymous functions, closures and first-class functions.
#### Concurrent
* Built-in concurrency. There are no heavy threads but channels.
* Ability to program and structure your programs in a synchronous manner but actually it’s asynchronus. Channels hide that complexity and lets you structure your programs in a more maintainable way.
#### Standard Library
* Almost all of the things built into its standard library (which is the library that comes with Go by default) like http fetching, json parsing, and encryption. So, this makes you faster and prevents fragmentation in the ecosystem (most of the time).

## References
* [About Go Language-An Overview](https://blog.learngoprogramming.com/about-go-language-an-overview-f0bee143597c)
* [Tutorials point](https://www.tutorialspoint.com/go/go_overview.htm)