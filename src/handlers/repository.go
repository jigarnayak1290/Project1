package handlers

import (
	"Project1/data"
	"database/sql"
	"net/http"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() (*Repository, error) {
	// Initialize the database connection
	err := data.InitializeDB()
	if err != nil {
		return nil, err
	}

	// Create a new UserRepository instance with the database connection
	repo := &Repository{
		db: data.GetDB(),
	}
	return repo, nil
}

func (p *Repository) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getVessal(rw, r)
		return
	}

	// if r.Method == http.MethodPost {
	// 	p.addProduct(rw, r)
	// 	return
	// }

	// // put via id
	// if r.Method == http.MethodPut {
	// 	p.l.Println("PUT", r.URL)

	// 	//Check for ID in URI
	// 	//reg := regexp.MustCompile(`/([0-9]+)`)
	// 	// reg := regexp.MustCompile(`NACCS_Code=(\d+)`)
	// 	// g := reg.FindAllStringSubmatch(r.URL.Path, -1)

	// 	queryParams := r.URL.Query()
	// 	// Get the query parameters
	// 	//queryParams := parsedURL.Query()

	// 	if len(queryParams) != 1 {
	// 		p.l.Println("Must have 1 parameter to be updated in URL, your parameter count is ", len(queryParams))
	// 		http.Error(rw, "Parameter count is mismatch", http.StatusBadRequest)
	// 		return
	// 	}

	// 	// Check if a parameter exists
	// 	var receivedParams []string
	// 	for param := range queryParams {
	// 		receivedParams = append(receivedParams, param)
	// 	}

	// 	firstParam := receivedParams[0]
	// 	if firstParam != "NACCS_Code" {
	// 		p.l.Println("Must have NACCS_Code in parameter, your parameter name is ", receivedParams[0])
	// 		http.Error(rw, "Incorrect parameter name ", http.StatusBadRequest)
	// 		return
	// 	}

	// 	// Access individual parameter values
	// 	NACCS_Code := queryParams.Get("NACCS_Code")
	// 	//NACCS_Code, _ := strconv.Atoi(param1)
	// 	// if len(g) != 1 {
	// 	// 	p.l.Println("Invalid URL ID is not one -> ", len(g))
	// 	// 	http.Error(rw, "Invalid URI", http.StatusBadRequest)
	// 	// 	return
	// 	// }

	// 	// if len(g[0]) != 2 {
	// 	// 	p.l.Println("Invalid URL more than capture group")
	// 	// 	http.Error(rw, "Invalid URI", http.StatusBadRequest)
	// 	// 	return
	// 	// }

	// 	// idString := g[0][1]
	// 	// id, err := strconv.Atoi(idString)
	// 	// if err != nil {
	// 	// 	p.l.Println("Invalid URL unable to convert to number", idString)
	// 	// 	http.Error(rw, "Invalid URI", http.StatusBadRequest)
	// 	// 	return
	// 	// }

	// 	//p.updateProducts(id, rw, r)
	// 	p.UpdateProductByNACCSCode(NACCS_Code, rw, r)
	// 	return
	// 	//p.l.Println("Got Id", id)
	// }

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// Implement the repository methods...
func (p *Repository) getVessal(rw http.ResponseWriter, r *http.Request) {
	//p.l.Println("Handle GET Products")
	//lp := data.getVesselByNaccsCode("123")
	lp := data.getVesselByNaccsCode("123")
	//d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	//rw.Write(d)
}
