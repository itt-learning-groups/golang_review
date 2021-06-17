# Ch. 11 Self Quiz

1. How do you know you are "done" reading data when using an `io.Reader`?

2. Create a simple but original example of a use of `io.ReadCloser` to read a local text file.

3. What can happen if you use `defer` to close a resource when opening a resource in a loop?

4. Early in Ch. 11, Bodner mentions that, "Any time you need to add additional methods to a type so that it can meet an interface, use this embedded type pattern." Give an original example of what Bodner is talking about.

5. What is the difference between the way `time.Duration` and `time.Time` are used?

6. Give an example of the use of `time.ParseDuration`.

7. How do you correctly check if two times are "the same"? Give an example.

8. How do you use `time.Parse`? What do you get?

9. How do you print a custom-formatted time string from a `time.Duration` instance?

10. What's the best way to get the `time.Duration` between two instants of time? Give an example.

11. Bodner mentions that "The most commonly used date and time formats have been given their own constants in the `time` package. What is he talking about?

12. How does the use-case for `json.Marshal`/`json.Unmarshal` differ from the use-case for `json.Decoder`/`json.Encoder`?

13. What is a common "gotcha" involving the use of "omitempty" in a JSON struct tag with a zero-value field value?

14. Why does Bodner note that `net/http` supports HTTP/2? What is that, and why would we care?

15. OK. But my project uses gRPC clients and servers. So my project isn't using HTTP, right?

16. What is a "mux" like `http.ServeMux` or the `gorilla/mux` Router? How does a `net/http` Server use a "mux" instance?

17. True/false/maybe: A REST server needs a "mux".

18. True/false/maybe: A gRPC server needs a "mux".

19. A lot of our Kount One gRPC-server projects do have a "mux" like the `http.ServeMux`. As an example, identify the "mux" being used here: <https://gitlab.gs.kount.com/kount/kntpo/k1-pbe-user-management/-/blob/master/cmd/k1-pbe-user-management-server/main.go>. What is the purpose of this "mux" and how does it relate to the gRPC server?

20. Describe the "middleware pattern".
