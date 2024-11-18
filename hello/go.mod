module example.com/hello

go 1.21.3

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	example.com/inputter v0.0.0-00010101000000-000000000000
)

replace example.com/inputter => ../inputter
