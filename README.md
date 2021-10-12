# Go Acorn Mail

#### Librerías
- [Acorn Email Framework](http://docs.thememountain.com/acorn/)

#### Instalación
```go
go get github.com/arskang/gomail-acorn-template
```

#### Ejemplo
```go
acorn := acornmail.NewAcornEmailComponents()
row := acorn.NewRow([]*acorntypes.Col{
    {
        Content: "Hola mundo",
    },
})
html := acorn.GetBoilerPlate(row)
fmt.Println(html)
```

#### Métodos
```go
import (
    acornmail "github.com/arskang/gomail-acorn-template"
)
```

- *GetHTMLString*: Obtener un html con parámetros dinámicos
```go
html, err = acornmail.GetHTMLString("<div>{{.Title}}</div>", map[string]interface{}{
    "Title": "Hola mundo",
})
if err != nil {
    panic(err)
}
fmt.Println(html)
```

- *NewAcornEmailComponents*: Obtener el tipo **acornEmail** para poder construir un template
```go
acorn := acornmail.NewAcornEmailComponents()
```

##### Components

- **GetBoilerPlate**
```go
boilertemplate := acorn.GetBoilerPlate("content")
fmt.Println(boilertemplate)
```

- **Row**
```go
w := "1/4"
row := acorn.NewRow([]*acorntypes.Col{
    {
        Content: "Content",
        Width: &w,
    }
})
fmt.Println(row)
```

- **Alerts**
```go
content := "Aceptar"
hexColor := "#008f38"
outlined := nil // *bool
alert := acorn.NewAlert(content, hexColor, outlined)
fmt.Println(alert)
```

#### Types
```go
import "github.com/arskang/gomail-acorn-template/acorntypes"
```

- *Col*
```go
w := "1/4"
col := acorntypes.Col {
    Content: "Content", // string
    Width: &w, // *string
}
```

#### Utilidades

- *Ancho de columnas*

```go
var width *string

// Grid a 4 columnas
width = "1/4"
width = "1/2"
width = "3/4"

// Grid a 3 columnas
width = "1/3"
width = "2/3"

// Todo el ancho
width = nil
width = "1"
width = "Cualquier otra cosa"
```