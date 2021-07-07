# Ch. 12 Self Quiz (Key)

1. What is metadata (for a network request/response)?

2. How can you tell if some data that is part of a request should be metadata or not? Give an example of an important piece of request data that should be considered metadata. Give an example of an important piece of request data that should *not* be considered metadata. Give an example of a piece of request data that is sort of "borderline": It could be metadata or not depending on its use.

3. What does Bodner mean when he says, "By default, you should prefer to pass data through explicit parameters"? Or when he says, "Go functions have explicit parameters and you shouldn't use the context as a way to sneak values past the API"?

4. What's the relationship between metadata and `context.Context`?

5. Describe the special use-case for `context.TODO` (vs `context.Background`).

6. How is context "treated as an immutable instance"?

7. Describe the need for context cancellation.

8. How is manual cancellation different from automatic timeout?

9. What are the two ways to create a context that will time out?

10. Why *"must"* you call the cancel function if you create a `context.WithCancel`?

11. How do you support context cancellation "in your own code"? What is meant by "in your own code" here?

12. How can you use the context `Err` method?

13. How is the timer you can create with an automatically timed-out context different from a timer you can create from the  `time.After` function?

14. Give an example of how you store or retrieve a value in a context.

15. Why is there a pattern for using specially-created, one-off types for the keys of key-value pairs stored on a context?

16. Describe the pattern for building "an API to place a value into the conext and to read the value from the context."

17. Bodner's suggestion for implementing a custom logger that retrieves a GUID stored in request metadata seems like a heck of a good idea. Should we all implement this in our projects tomorrow?
