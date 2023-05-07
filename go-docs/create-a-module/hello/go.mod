module create-a-module/hello

go 1.20

replace create-a-module/greetings => ../greetings

require create-a-module/greetings v0.0.0-00010101000000-000000000000
