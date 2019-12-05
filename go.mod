module github.com/jmontupet/gbgl

go 1.13

require (
	github.com/go-gl/gl v0.0.0-20190320180904-bf2b1f2f34d7
	github.com/go-gl/glfw v0.0.0-20190409004039-e6da0acd62b1
	github.com/go-gl/mathgl v0.0.0-20190713194549-592312d8590a
	github.com/jmontupet/gbcore v0.0.0-20191205034804-61d529614f73
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b // indirect
)

replace github.com/jmontupet/gbcore => ../go-gb/
