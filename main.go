// go utility commands
// go mod tidy
// go mod init home.udp1024.com/web-service-gin
// go init .
// go get github.com/gin-gonic/gin
// go get .
// go run .

package main

import (
	"encoding/json"
	"net/http/httputil"
	"strconv"

	//"io"
	"log"
	"net/http"
	"os"

	//"strings"

	"github.com/gin-gonic/gin"
)

type card struct {
	ID          string `json:"id"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

var cards []card

func main() {
	var err error
	cards, err = ReadCards()
	if err != nil {
		log.Fatalf("Failed to read cards in main(): %v", err)
		return
	}
	router := gin.Default()
	router.GET("/cards", getCards)
	router.GET("/cards/:id", getCardByID)
	router.POST("/cards", createCard)
	router.PUT("cards/:id", updateCard)
	router.DELETE("cards/:id", deleteCard)

	router.Run("localhost:8080")
}

// createCard function creates a new card from data sent in http request
// the new card is assigned the next sequential id in the array
func createCard(c *gin.Context) {
	var newCard card

	dump, erro := httputil.DumpRequest(c.Request, true)
	if erro != nil {
		// handle error
		c.JSON(http.StatusBadRequest, gin.H{"status ": http.StatusBadRequest, "message ": erro.Error()})
		return
	}
	log.Println(string(dump))

	err := c.BindJSON(&newCard)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	//newCardID := strconv.Itoa(len(cards) + 1)
	newCardID := getNextID(cards)
	newCard.ID = newCardID
	log.Println("New Card ID will be: ", newCard.ID)

	cards = append(cards, newCard)
	log.Println("New card appended to the slice")

	// Write the cards slice back to the file
	err = writeCardsToFile(cards, "json/data.json")
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}
	log.Println("File written.")
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": newCard})
}

// getCards responds with the list of all cards as JSON.
func getCards(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cards)
}

// getCardByID locates the card whose ID value matches the id
// parameter sent by the client, then returns that card as a response.
func getCardByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of cards, looking for
	// a card whose ID value matches the parameter.
	for _, a := range cards {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "card not found"})
}

// updateCard locates the card whose id matches the id in the data then updates the card with provided data
func updateCard(c *gin.Context) {
	var updatedCard card

	// dump, erro := httputil.DumpRequest(c.Request, true)
	// if erro != nil {
	// 	// handle error
	// 	c.JSON(http.StatusBadRequest, gin.H{"status ": http.StatusBadRequest, "message ": erro.Error()})
	// 	return
	// }
	// log.Println(string(dump))

	// bind the json body of request to structure card
	err := c.BindJSON(&updatedCard)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	// log.Println("called updateCard function for card.ID ", updatedCard.ID)

	for index, icard := range cards {
		// log.Println("enumerating ...")
		// log.Println("value of current card id ", icard.ID)
		if icard.ID == updatedCard.ID {
			//parse the request for updated card data
			// log.Println("match found ...")

			//update the card with new data
			cards[index] = updatedCard

			// Write the cards slice back to the file
			err := writeCardsToFile(cards, "json/data.json")
			if err != nil {
				// log.Println("Error writing to file:", err)
				return
			}

			// return updated card
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": updatedCard})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Card not found"})
}

func deleteCard(c *gin.Context) {
	id := c.Param("id")

	// log.Println("Deleting card id ", id)

	for index, icard := range cards {
		// log.Println("Card ID ", icard.ID, " - seeking ", id)
		if icard.ID == id {
			// log.Println("match found. Deleting ...")
			cards = append(cards[:index], cards[index+1:]...)

			// Write the cards slice back to the file
			err := writeCardsToFile(cards, "json/data.json")
			if err != nil {
				// log.Println("Error writing to file:", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Card deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Card could not be deleted, not found!"})
}

// function writeCardsToFile writes the cards to the specified disk file
func writeCardsToFile(cards []card, filePath string) error {
	// Open the file
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new JSON encoder that writes to the file
	encoder := json.NewEncoder(file)

	// Write the cards slice to the file in JSON format
	err = encoder.Encode(cards)
	if err != nil {
		return err
	}

	return nil
}

// ReadCards reads the json/data.json file into a slice of cards
func ReadCards() ([]card, error) {
	data, err := os.ReadFile("json/data.json")
	if err != nil {
		return nil, err
	}

	var cards []card
	err = json.Unmarshal(data, &cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

// function getNextID scans the cards to find the highest id, adds 1 to it and returns the resulting number as a string to the caller
func getNextID(cards []card) string {
	maxID := 0

	for _, card := range cards {
		// Convert the ID string to an integer
		id, err := strconv.Atoi(card.ID)
		if err != nil {
			continue
		}

		// Check if this ID is larger than the current maxID
		if id > maxID {
			maxID = id
		}
	}

	// Increment maxID and convert it back to a string
	nextID := maxID + 1
	return strconv.Itoa(nextID)
}
