package organization

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/austin-rey/go-react-trello-clone/cors"
)

const organizationPath = "organization"

func handleOrganization(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", organizationPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	orgID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
		case http.MethodGet:
			organization := getOrganization(orgID)
			if organization == nil{
				w.WriteHeader(http.StatusNotFound)
			}

			o, err := json.Marshal(organization)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = w.Write(o)
			if err != nil {
				log.Fatal(err)
			}

		case http.MethodPut:
			var organization Organization
			err := json.NewDecoder(r.Body).Decode(&organization)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			
			if organization.OrganizationID != orgID {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = updateOrganization(organization)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		case http.MethodDelete:
			deleteOrganization(orgID)

		case http.MethodOptions:
			return
			
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleOrganizations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			organizations := getAllOrganizations()
			o,err := json.Marshal(organizations)
			if err != nil {
				log.Fatal(err)
			}

			_,err = w.Write(o)
			if err != nil {
				log.Fatal(err)
			}
		case http.MethodPost:
			var organization Organization

			// Decode user to see if it matches user struct
			err := json.NewDecoder(r.Body).Decode(&organization)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = createOrganization(organization)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
	}
}

func handleOrganizationMembers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func SetupRoutes(apiBasePath string) {
	organizationHandler := http.HandlerFunc(handleOrganization)
	organizationMemebersHandler := http.HandlerFunc(handleOrganizationMembers)
	organizationsHandler := http.HandlerFunc(handleOrganizations)
	
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, organizationPath), cors.Middleware(organizationHandler))
	http.Handle(fmt.Sprintf("%s/%s/members", apiBasePath, organizationPath), cors.Middleware(organizationMemebersHandler))
	http.Handle(fmt.Sprintf("%s/%ss", apiBasePath, organizationPath), cors.Middleware(organizationsHandler))
}