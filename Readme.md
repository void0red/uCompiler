# uCompiler
> a simple interpreter for a simple drawing language

## Usage
```shell
Usage of .\bin\release.exe:
  -i string
        source file to parse
  -o string
        output file to save (default "out.png")
```

## Feature
1. Drawing language tokens
    - normal: CONST_ID
    - arguments: T
    - function: FUNC
    - reserved:	ORIGIN SCALE ROT IS TO STEP DRAW FOR FROM COLOR MAP
    - operator:	PLUS MINUS MUL DIV POWER
    - separator: COMMA SEMICOLON L_BRACKET R_BRACKET
    - special: EOF ERROR
2. Case insensitive
3. Use strict `;` to end the statement
4. Support Statement
    - MapStatement
        - used to set the image bound
        - example: `map is (800, 800);`
        - default bound is `(800, 600)`
        - warning: the statement can only write at the first line, otherwise it won't work
    - RotStatement
        - used to set the rot angle
        - example: `rot is 10;`
        - default rot angle is `0`
    - ScaleStatement
        - used to set the scale size
        - example: `scale is (200, 200);`
        - default scale size is `(1, 1)`
    - OriginStatement
        - used to set the origin point
        - example: `origin is (400, 400);`
        - default origin is `(0, 0)`
    - ColorStatement
        - used to set the draw RGBA color
        - example: `color is (255, 0, 0, 255);`
        - default color is `(255, 255, 255, 255)`
    - ForStatement
        - used to draw the point
        - example: `for t from 0 to 400 step 0.0001 draw(t, t);`
        - warning: statements can't be nested, because only a variable `T` can be identified :-P
5. function support
    - sin, cos, tan
    - abs, ln, exp, abs
    
## Build
1. add the project path to your $GOPATH
2. go build src
