package main

import (
	"encoding/json"
	"net/http"
	"retro/config"
	"retro/retroachievements"

	"go.uber.org/zap"
)

func getGameInformation(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "getGameInformation"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	game, err := retroachievements.GetGameInformation(logger, body.Name, body.Key)
	if err != nil {
		logger.Error("Error getting game information",
			zap.String("Username", body.Name),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(game); err != nil {
		logger.Error("Error writing payload",
			zap.String("function", "getGameInformation"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func setGameOrder(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Game  string   `json:"game"`
		Order []string `json:"order"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "setGameOrder"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	config.SetGameOrder(body.Game, body.Order)
}

func getMetadata(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Mode    string `json:"mode"`
		Refresh int    `json"refresh"`
	}{
		Mode:    config.GetGameMode(),
		Refresh: config.GetRefresh(),
	}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Error("Error writing payload",
			zap.String("function", "getMetadata"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func setGameMode(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Mode string `json:"mode"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "setGameMode"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	config.SetGameMode(body.Mode)
}

func setRefresh(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Refresh int `json:"refresh"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "setRefresh"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := config.SetRefresh(body.Refresh)
	if err != nil {
		logger.Error("Error setting refresh rate",
			zap.Int("refresh", body.Refresh),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully set refresh rate",
		zap.Int("refresh", body.Refresh),
	)
}

func getLayoutList(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Loggerayouts map[string]string `json:"layouts"`
	}{
		Loggerayouts: config.GetLayouts(),
	}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Error("Error writing payload",
			zap.String("function", "getLayoutList"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func addLayout(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name   string `json:"name"`
		Layout string `json:"layout"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "addLayout"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := config.AddLayout(body.Name, body.Layout); err != nil {
		logger.Error("Error adding layout",
			zap.String("layout", body.Name),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully added layout",
		zap.String("layout", body.Name),
	)
}

func removeLayout(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name string `json:"name"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "removeLayout"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := config.RemoveLayout(body.Name); err != nil {
		logger.Error("Error removing layout",
			zap.String("layout", body.Name),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully removed layout",
		zap.String("layout", body.Name),
	)
}

func updateLayout(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name   string `json:"name"`
		Layout string `json:"layout"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "updateLayout"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := config.UpdateLayout(body.Name, body.Layout); err != nil {
		logger.Error("Error updating layout",
			zap.String("layout", body.Name),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Don't send an info here 'cause that'd spam a lot
}

func getActiveLayout(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Active string `json:"layout"`
	}{
		Active: config.GetActiveLayout(),
	}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Error("Error writing payload",
			zap.String("function", "getActiveLayout"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func setActiveLayout(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name string `json:"name"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "setActiveLayout"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := config.SetActiveLayout(body.Name); err != nil {
		logger.Error("Error setting active layout",
			zap.String("layout", body.Name),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully set active layout",
		zap.String("layout", body.Name),
	)
}

func getProfiles(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Profiles map[string]config.Profile `json:"profiles"`
	}{
		Profiles: config.GetProfiles(),
	}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Error("Error writing payload",
			zap.String("function", "getProfiles"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func addProfile(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "addProfile"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := config.AddProfile(body.Name, body.Key); err != nil {
		logger.Error("Error adding profile",
			zap.String("profile", body.Name),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully added profile",
		zap.String("profile", body.Name),
	)
}

func removeProfile(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name string `json:"name"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "removeProfile"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := config.RemoveProfile(body.Name); err != nil {
		logger.Error("Error removing profile",
			zap.String("profile", body.Name),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully removed profile",
		zap.String("profile", body.Name),
	)
}

func getActiveProfile(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Active string `json:"profile"`
	}{
		Active: config.GetActiveProfile(),
	}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Error("Error writing payload",
			zap.String("function", "getActiveProfile"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func setActiveProfile(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name string `json:"name"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Error("Error decoding payload",
			zap.String("function", "setActiveProfile"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := config.SetActiveProfile(body.Name); err != nil {
		logger.Error("Error setting active profile",
			zap.String("profile", body.Name),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully set active profile",
		zap.String("profile", body.Name),
	)
}
