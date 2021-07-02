# Ch. 12 Self Quiz

1. Difference between Background and TODO context?
2. Initialisation of app what kind of context we can use? 
3. Difference between WithTimeOut and WithDeadLine context?
4. What the use of WithValue context?
5. What the role of Done() function?
6. What will be the output?
func main() {

	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Hello")
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Saurabh")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

7. what indicates cancel() function?
8. In which sitiuation we can use TODO context?