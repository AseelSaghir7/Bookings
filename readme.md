# Bookings

The repository for Building Modern Web Applications with Go


>Starting from **35 sharing data with templates...**

- for sending data from handler to templates we will use a type which will hold template data.
- Import cycle not allowed here. We cannot have packages that depend upon each others
- The problem here is that I have packages that are importing each other and you're not allowed to do that and go. You can do it in some languages like C++. You're allowed to well, you have to put some serious hacks in place to make it work. But you can do that. You can't do it. And go goes a little more strict and says, no, sorry, can't have packages that depend upon each other like that.And it refuses to compile.And the irritating thing about it is no idea.I have ever tried ever warned you about the import cycle while you're putting all the time into writing code.

```go
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
```

- This code should be in models folder instead of handlers otherwise import cycle not allowed error occors. because render package uses that.