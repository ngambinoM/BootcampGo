package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)
var (
	ErrCannotpen = errors.New("cannot open de file")
	ErrCannotRead = errors.New("cannot  read de file")
	ErrNotParseToInt= errors.New("cannot parse string to int")
	ErrNotParseToFloat= errors.New("cannot parse string to Float32")
	ErrNotFoundAny = errors.New("there are not results")
	ErrNotValidInput=errors.New("input is not a valid data")
)
type Ticket struct {
	id int
	name string
	email string
	destiny string
	timeFlight time.Time
	price float32
}

func ExtractData () ([]Ticket ,error){
	var ticketsData = []Ticket{}
	file, err := os.Open("tickets.csv")
	if err != nil {
        return nil, fmt.Errorf("error: %w", ErrCannotpen)
    }
    defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ','
	r.FieldsPerRecord = 6
	

	data, err := r.ReadAll()
	if err != nil {
        return nil, fmt.Errorf("error: %w", ErrCannotRead)
    }
	for _, line := range data {
        // Obtener valores de cada campo
        id := line[0]
        name := line[1]
        email := line[2]
        destiny := line[3]
        hourFlight := line[4]
        price :=  line[5]
		
	i, err := strconv.Atoi(id)
    	if err != nil {
        	return nil, fmt.Errorf("error: %w", ErrNotParseToInt)
        	
    	}
	hora, err := time.Parse("15:04", hourFlight)
	if err != nil {
		return nil, fmt.Errorf("error: %w", ErrNotParseToFloat)
		
	}
	priceT, err := strconv.ParseFloat(price, 32)
		if err != nil {
			return nil, fmt.Errorf("error: %w", ErrNotParseToFloat)
			
		}
		ticketsData = append(ticketsData, Ticket{i,name,email,destiny,hora,float32(priceT)})

	}
	
		return ticketsData, nil	
}

// ejemplo 1
func GetTotalTickets(destiny string, tickets []Ticket) (int,error) {
	counter :=0
		for _, ticket := range tickets {
			if ticket.destiny==destiny {
				counter ++
			}
		}
		if counter==0 {
			return 0,ErrNotFoundAny
		}
		return counter,nil
}

// ejemplo 2
func GetCountByPeriod(timeFlight string, tickets []Ticket) (int , error) {
	counter :=0
	var a , b string
	switch timeFlight {
	case "earlymorning":
		a="0:00"; b="6:59"
	case "morning":
		a="7:00"; b="12:59"
	case "afternoon":
		a="13:00"; b="19:59"
	case "night":
		a="20:00"; b="23:59"	
	default: 
			return 0, ErrNotValidInput
	}
	timeLow, _ := time.Parse("15:04", a)
	timeHight, _ := time.Parse("15:04", b)
		for _, ticket := range tickets {
			if ticket.timeFlight.After(timeLow) && ticket.timeFlight.After(timeHight){
				counter ++
			}
		}
		return counter,nil
}

// ejemplo 3
func AverageDestination(destination string, tickets []Ticket) (float32,error) {
	total,err := GetTotalTickets(destination,tickets)
	if err!=nil {
		return 0, ErrNotFoundAny
	}
	return float32(total*100)/float32(len(tickets)),nil
}