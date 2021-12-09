package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlgallego99/OSTfind/internal/cancion"
)

// Guardar temporalmente las OSTs en una variable global
var osts []*cancion.BandaSonora

type Cancion_msg struct {
	Titulo     string `json:"titulo"`
	Compositor string `json:"compositor"`
	Genero     string `json:"genero"`
}

type Canciones_msg struct {
	Nombre    string        `json:"nombre"`
	Canciones []Cancion_msg `json:"canciones"`
}

func newOST(c *gin.Context) {
	var ost *cancion.BandaSonora
	var canciones []*cancion.Cancion_info
	var err error

	// Leer cuerpo de la petición
	cancionesmsg := new(Canciones_msg)
	err = c.BindJSON(cancionesmsg)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	obra := c.Param("obra")
	switch obra {
	case "videojuego":
		ost, err = cancion.NewVideojuegoOST(cancionesmsg.Nombre, make([]*cancion.Cancion_info, 0))

	case "serie":
		ost, err = cancion.NewSerieOST(cancionesmsg.Nombre, 1, 1, make([]*cancion.Cancion_info, 0))

	case "pelicula":
		ost, err = cancion.NewPeliculaOST(cancionesmsg.Nombre, make([]*cancion.Cancion_info, 0))

	default:
		err = errors.New("no se reconoce el tipo de OST")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Añadir canciones de la ost
	for _, cmsg := range cancionesmsg.Canciones {
		can, err := cancion.NewCancion(cmsg.Titulo, cmsg.Compositor, cancion.StringToGenero[cmsg.Genero])

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		canciones = append(canciones, can)
	}

	ost.ActualizarOST(canciones)
	osts = append(osts, ost)

	c.JSON(http.StatusOK, gin.H{
		"message": "OST creada",
		"ost": gin.H{
			"id":        ost.Id,
			"nombre":    ost.Obra.Titulo(),
			"canciones": ost.Canciones,
		},
	})
}

func getOST(c *gin.Context) {
	var err error

	obra := c.Param("obra")
	ostName := c.Param("ost")

	switch obra {
	case "videojuego", "serie", "pelicula":
		err = nil

	default:
		err = errors.New("no se reconoce el tipo de OST")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for _, ost := range osts {
		if ost.Obra.Titulo() == ostName {
			err = nil

			c.JSON(http.StatusOK, gin.H{
				"message": "OST encontrada",
				"ost": gin.H{
					"nombre":    ost.Obra.Titulo(),
					"canciones": ost.Canciones,
				},
			})

			return
		} else {
			err = errors.New("no existe esa OST")
		}
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
