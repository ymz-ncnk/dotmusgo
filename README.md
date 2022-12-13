# DotMusGo
Creates MusGo, which by default puts "a__" at the beginning of the each 
generated file name. This way all generated files can be grouped by name.

# How to use
Create MusGo with help of:
```go
	musGo, err := dotmusgo.New()
	if err != nil {
		t.Fatal(err)
	}
	// use musGo as usual
	... 
```