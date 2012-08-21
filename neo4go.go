package neo4go

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Error struct {
	Message     string
	Exception   string
	Stackstrace []string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Exception)
}

type Neo4j struct {
	URL string
}

type Node struct {
	Id                  uint
	Relationships       string
	RelationshipsOut    string `json:"outgoing_relationships"`
	RelationshipsIn     string `json:"incoming_relationships"`
	RelationshipsAll    string `json:"all_relationships"`
	RelationshipsCreate string `json:"create_relationship"`
	Data                map[string]interface{}
	Traverse            string
	Property            string
	Properties          string
	Self                string
	Extensions          map[string]interface{}
	Start               string // relationships & traverse // returns both obj & string
	End                 string // relationships & traverse // returns both obj & string
	Type                string // relationships & traverse
	Indexed             string // index related
	Length              string // traverse framework
}

/*
 {   
    "outgoing_relationships" : "http://localhost:7474/db/data/node/0/relationships/out",   
    "data" : {     "name" : "foobar"   },   
        "traverse" : "http://localhost:7474/db/data/node/0/traverse/{returnType}",   
        "all_typed_relationships" : "http://localhost:7474/db/data/node/0/relationships/all/{-list|&|types}",   
        "self" : "http://localhost:7474/db/data/node/0",   "property" : "http://localhost:7474/db/data/node/0/properties/{key}",   
        "properties" : "http://localhost:7474/db/data/node/0/properties",   
        "outgoing_typed_relationships" : "http://localhost:7474/db/data/node/0/relationships/out/{-list|&|types}",   
        "incoming_relationships" : "http://localhost:7474/db/data/node/0/relationships/in",   
        "extensions" : {   },   
        "create_relationship" : "http://localhost:7474/db/data/node/0/relationships",   
        "paged_traverse" : "http://localhost:7474/db/data/node/0/paged/traverse/{returnType}{?pageSize,leaseTime}",   
        "all_relationships" : "http://localhost:7474/db/data/node/0/relationships/all",   
        "incoming_typed_relationships" : "http://localhost:7474/db/data/node/0/relationships/in/{-list|&|types}" }
*/

type Property map[string]string

func NewNeo4j(uri string) (*Neo4j, error) {
	neo := new(Neo4j)
	neo.URL = uri

	return neo, nil
}

func (neo *Neo4j) GetNode(id uint) (*Node, error) {
	url := neo.URL + "/db/data/node/%d"
	url = fmt.Sprintf(url, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Read all data from response
	body, _ := ioutil.ReadAll(resp.Body)
	node := new(Node)
	if resp.StatusCode == 200 {
		err = json.Unmarshal(body, node)
		node.Id = id
		return node, err
	} else {
		err := new(Error)
		_ = json.Unmarshal(body, err)
		log.Printf("%s\n", string(body))
	}
	return node, err
}

// CreateNode creates a new node and returns it
func (neo *Neo4j) CreateNode(data map[string]string) (*Node, error) {
	_, err := json.Marshal(data)
	//node := new(Node)
	if err != nil {
		return nil, errors.New("Unable to Marshal Json data")
	}
	return nil, nil
}
