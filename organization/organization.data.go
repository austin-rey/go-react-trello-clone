package organization

import (
	"fmt"
	"sort"
	"sync"
)

var organizationMap = struct {
	sync.RWMutex
	o map[int]Organization
}{o: make(map[int]Organization)}

func getOrganization(organizationId int) (*Organization) {
	organizationMap.RLock()
	defer organizationMap.RUnlock()
	if organization, ok:= organizationMap.o[organizationId]; ok {
		return &organization
	}
	return nil
}

// func getOrganizationMembers(organizationId int) []user.User {
	
// }

func updateOrganization(organization Organization)(int, error) {
	oldOrg := getOrganization(organization.OrganizationID)
	oldOrganizationId := organization.OrganizationID

	if oldOrg == nil {
		return 0, fmt.Errorf("org id [%d] doesn't exist", oldOrg.OrganizationID)
	}
	organizationMap.Lock()
	organizationMap.o[oldOrganizationId] = organization
	organizationMap.Unlock()
	return oldOrganizationId, nil
}

func deleteOrganization(organizationId int) {
	organizationMap.Lock()
	defer organizationMap.Unlock()
	delete(organizationMap.o, organizationId)
}

func createOrganization(organization Organization)(int, error) {
	nextOrgId := getNextOrganizationID()
	organization.OrganizationID = nextOrgId
	organizationMap.Lock()
	organizationMap.o[nextOrgId] = organization
	organizationMap.Unlock()
	return nextOrgId, nil
}

func getAllOrganizations()[]Organization {
	organizationMap.RLock()
	organizations := make([]Organization,0,len(organizationMap.o))
	for _, value := range organizationMap.o {
		organizations = append(organizations, value)
	}
	organizationMap.RUnlock()
	return organizations
}

// Utility Functions -------------------------------------------------

func getOrganizationIds() []int {
	organizationMap.RLock()
	organizationIds := []int{}
	for key := range organizationMap.o {
		organizationIds = append(organizationIds, key)
	}
	organizationMap.RUnlock()
	sort.Ints(organizationIds)
	return organizationIds
}

func getNextOrganizationID() int {
	organizationIds := getOrganizationIds()
	fmt.Println(organizationIds)

	if len(organizationIds) == 0 {
		return 1
	}

	return organizationIds[len(organizationIds)-1] + 1
}