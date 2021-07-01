# Golang Review (June 30, 2021)

## Goals

1. Each of us should...

* Achieve clarity about...

  * what Go `context` ***is***, what it is ***for***, and how it is idiomtically ***used***
  * how "parent" and "child"/derived contexts are ***implemented***; highlight how this influences our `context` use
  * what *metadata* means and how context *values* should be used

* Without looking up examples or docs, be able to...

  * implement an idiomatic context cancellation
  * implement an idiomatic timeout...
    * *across* an API boundary
    * (*within* an API boundary) <-- Extra credit for now: We'll be looking at specific examples of this when we return to ch.10

* Be confident in our individual ability/commitment to ***teach ourselves*** the majority of the material this meeting covers by having/developing...

  * a reliably creative (vs passive) way to engage with the material
  * a reliable time/attention budget for it
  * skill identifying and prioritizing new materials on our own

* Be confident that we understand what is...

  * really important
  * not important

  ...about unit testing, and ***why***

2. Have a team discussion about how we can, should, and will commit to reform our *modus operandi* for testing in day-to-day work

   *(Modus Operandi: A particular way or method of doing something, especially one that is characteristic or well-established.)*

3. Time allowing...

* Do a walk-through of implementing a feature using...
  * a test-driven style
  * a disciplined strategy for atomic (well, let's say atomic-ish) merge requests

* We'll likely not have time to get far into ch.13. But if so, for this week each of us should be able to...

  * create a table test from scratch
  * understand when/how you should exit early from a test
  * understand the tradeoffs of testing within vs outside the same package as the production code
  * understand how to identify, avoid, and teach other team members about approx half a dozen or so common mistakes/pitfalls when writing unit tests

## "Lab"

### Context (*Learning Go*, ch.12)

* What Go `context` ***is***, what it is ***for***, and how it is idiomtically ***used***

  * Take a look at the `context` [package](https://golang.org/src/context/context.go)
    * In one sentence, answer the question "What is a context?" Write your answer in the meeting chat window.

  * Review the four methods of the `context` interface.
    * In a moment, we'll each attempt to recall and describe these methods from memory. Take a moment to memorize the four methods.
      * When you're ready, hide the doc from yourself and type the four method names into the chat window.

    * Prepare to describe each of these methods from memory using your understanding of what they ***do*** and how you would/should ***use*** each.
      * Let's have a discussion to help each of us get clarity on each method.
        * You may find it useful to browse through the rest of the `context` package as work on this.
        * Find, then share, an example of each interface method being used in code... somewhere.
        * *(You may find it helpful to leave `Deadline` until you've worked on the other 3)*
      * When you're ready, hide the doc from yourself and write a description of each method in the chat window.
        * Do not copy & re-write the method descriptions from the godoc. Use your own understanding of what the method is for, how it would be used, and how it supports the intended purpose of Go `context`.
  
  * Question: What is special about `context` and API boundaries? What does "API" mean here?

  * Find or create, then share, an example of `context` used to implement cancellation or timeout across (an) API boundary(ies)
    * You can find examples in blogs, etc., by Googling -- but finding an example or two in Kount code would be even better

  * Find or create, then share, an example of implementing a goroutine cancellation or timeout *without* using `context`
    * Do you have to use `context` to do a cancellation or timeout?
    * When *should* you use `context` to do a cancellation or timeout?

* How "parent" and "child"/derived contexts are ***implemented***; highlight how this influences our `context` use

  * Take a look at this short *Medium* article (especially the pictures): [Go: Context and Cancellation by Propagation](https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c), by Vincent Blanchon
    * (You may want to refer back to the `context` package to follow what Blanchon has to say)
  * Use one of the diagrams and accompanying code sample from Blanchon's article to describe...
    * what is a "parent" vs "child"/derived context
    * the one-direction orientation of context cancellation in a family of `context`s
    * what happens to sibling/concurrent contexts when the parent context is cancelled
  * How does the `cancelCtx` type help us explain why "any time you create a cancellable context, you *must* call the cancel function".
    * What does calling a cancel function ***do***?
    * What if you created your conext using WithTimeout or WithDeadline rather than WithCancel. Must you still call a cancel function?

* What *metadata* means and how context *values* should be used
  * What is *metadata* in "context" of Go `context`?
  * What constitutes the sort of metadata that *should* be stored in `context` values?
    * Find, then share, an example of the good use of context metadata in Kount code.
  * What constitutes the sort of metadata that *should* ***not*** be stored in `context` values?
    * Find, then share, an example of request data in Kount code that would ***not*** be appropriate as a `context` value.
  * What is a good rule of thumb for the proper & idiomatic use of context values in Go?
  * Futher food for thought: Check out [How to correctly use context.Context in Go 1.7](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39), by Jack Lindamood

* Homework:
  * Without looking up examples or docs, be able to...

    * implement an idiomatic context cancellation
    * implement an idiomatic timeout...
      * *across* an API boundary
