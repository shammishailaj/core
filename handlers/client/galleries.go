package clienthandlers

import (
	"net/http"

	"github.com/backpulse/core/database"
	"github.com/backpulse/core/utils"
	"github.com/gorilla/mux"
)

// GetGallery : return specific gallery
func GetGallery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	gallery, err := database.GetGallery(id)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, "not_found", nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "success", gallery)
	return
}

// GetGalleries : return array of galleries
func GetGalleries(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	site, err := database.GetSiteByName(name)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, "not_found", nil)
		return
	}

	galleries, err := database.GetGalleries(site.ID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, "not_found", nil)
		return
	}

	for i := range galleries {
		photos, _ := database.GetGalleryPhotos(galleries[i].ID)
		galleries[i].Photos = photos
	}

	utils.RespondWithJSON(w, http.StatusOK, "success", galleries)
	return
}

// GetHomeGallery : return home gallery of site
func GetDefaultGallery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	site, err := database.GetSiteByName(name)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, "not_found", nil)
		return
	}
	gallery, err := database.GetDefaultGallery(site.ID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, "not_found", nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "success", gallery)
	return

}
