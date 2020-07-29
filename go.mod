module github.com/jmontupet/gbgl

go 1.14

require (
	github.com/go-gl/gl v0.0.0-20190320180904-bf2b1f2f34d7
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20200707082815-5321531c36a2
	github.com/go-gl/mathgl v0.0.0-20190713194549-592312d8590a
	github.com/jmontupet/gbcore v0.0.0-20191205034804-61d529614f73
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b // indirect
)

replace github.com/jmontupet/gbcore => ../gbcore/
