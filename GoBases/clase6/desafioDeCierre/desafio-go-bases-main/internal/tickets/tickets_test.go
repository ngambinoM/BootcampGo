package tickets

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	tickets := []Ticket{
		{
			id:         1,
			name:       "Juan",
			email:      "juan@example.com",
			destiny:    "New York",
			timeFlight: time.Now(),
			price:      100.0,
		},
		{
			id:         2,
			name:       "Maria",
			email:      "maria@example.com",
			destiny:    "Argentina",
			timeFlight: time.Now().Add(24 * time.Hour),
			price:      200.0,
		},
		{
			id:         3,
			name:       "Pedro",
			email:      "pedro@example.com",
			destiny:    "Argentina",
			timeFlight: time.Now().Add(48 * time.Hour),
			price:      300.0,
		},
	}
	//resultadoEsperado:=2
	_,err:=GetTotalTickets("rgentina",tickets)
	//assert.Equal(t,resultadoEsperado,resultadActual)
	assert.Error(t,err)
}
