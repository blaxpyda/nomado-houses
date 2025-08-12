package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

// TravelPayoutsService interface defines methods for TravelPayouts API operations
type TravelPayoutsService interface {
	SearchHotels(destination, checkIn, checkOut string, adults int) (*HotelSearchResponse, error)
	GetPopularHotels(cityID string) (*PopularHotelsResponse, error)
	GetDestinations(query string) (*DestinationsResponse, error)
}

// travelPayoutsService implements TravelPayoutsService
type travelPayoutsService struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewTravelPayoutsService creates a new TravelPayouts service
func NewTravelPayoutsService() TravelPayoutsService {
	return &travelPayoutsService{
		apiKey:  os.Getenv("TRAVEL_PAYOUTS_API_KEY"),
		baseURL: "https://api.travelpayouts.com/v1/hotels",
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Hotel represents a hotel from TravelPayouts API
type Hotel struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	Stars       int      `json:"stars"`
	Rating      float64  `json:"rating"`
	PriceFrom   float64  `json:"priceFrom"`
	Currency    string   `json:"currency"`
	ImageURL    string   `json:"imageUrl"`
	Description string   `json:"description"`
	Amenities   []string `json:"amenities"`
	Location    struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"location"`
}

// HotelSearchResponse represents the response from hotel search API
type HotelSearchResponse struct {
	Hotels []Hotel `json:"hotels"`
	Status string  `json:"status"`
	Error  string  `json:"error,omitempty"`
}

// PopularHotelsResponse represents popular hotels in a destination
type PopularHotelsResponse struct {
	Hotels []Hotel `json:"hotels"`
	Status string  `json:"status"`
}

// Destination represents a destination from TravelPayouts
type Destination struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	IATA    string `json:"iata"`
}

// DestinationsResponse represents destinations search response
type DestinationsResponse struct {
	Destinations []Destination `json:"destinations"`
	Status       string        `json:"status"`
}

// SearchHotels searches for hotels in a destination
func (s *travelPayoutsService) SearchHotels(destination, checkIn, checkOut string, adults int) (*HotelSearchResponse, error) {
	if s.apiKey == "" {
		return nil, fmt.Errorf("TravelPayouts API key is not configured")
	}

	// Build URL with parameters
	params := url.Values{}
	params.Set("query", destination)
	params.Set("check_in", checkIn)
	params.Set("check_out", checkOut)
	params.Set("adults", fmt.Sprintf("%d", adults))
	params.Set("currency", "usd")
	params.Set("limit", "20")

	apiURL := fmt.Sprintf("%s/search?%s", s.baseURL, params.Encode())

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Access-Token", s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("invalid API key or authentication failed")
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("API rate limit exceeded, please try again later")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	// TravelPayouts API returns data in a different structure
	var apiResponse struct {
		Data []struct {
			ID      string  `json:"id"`
			Name    string  `json:"name"`
			Address string  `json:"address"`
			Stars   int     `json:"stars"`
			Rating  float64 `json:"rating"`
			Price   struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"price"`
			Image       string   `json:"main_image_url"`
			Description string   `json:"description"`
			Amenities   []string `json:"amenities"`
			Location    struct {
				Lat float64 `json:"latitude"`
				Lng float64 `json:"longitude"`
			} `json:"location"`
		} `json:"data"`
		Success bool   `json:"success"`
		Error   string `json:"error,omitempty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	if !apiResponse.Success && apiResponse.Error != "" {
		return nil, fmt.Errorf("API error: %s", apiResponse.Error)
	}

	// Convert to our format
	var hotels []Hotel
	for _, item := range apiResponse.Data {
		hotel := Hotel{
			ID:          item.ID,
			Name:        item.Name,
			Address:     item.Address,
			Stars:       item.Stars,
			Rating:      item.Rating,
			PriceFrom:   item.Price.Amount,
			Currency:    item.Price.Currency,
			ImageURL:    item.Image,
			Description: item.Description,
			Amenities:   item.Amenities,
			Location: struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			}{
				Lat: item.Location.Lat,
				Lng: item.Location.Lng,
			},
		}
		hotels = append(hotels, hotel)
	}

	return &HotelSearchResponse{
		Hotels: hotels,
		Status: "success",
	}, nil
}

// GetPopularHotels gets popular hotels in a city
func (s *travelPayoutsService) GetPopularHotels(cityID string) (*PopularHotelsResponse, error) {
	if s.apiKey == "" {
		return nil, fmt.Errorf("TravelPayouts API key is not configured")
	}

	apiURL := fmt.Sprintf("%s/popular?location=%s&limit=10", s.baseURL, cityID)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Access-Token", s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("invalid API key or authentication failed")
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("city not found with ID: %s", cityID)
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("API rate limit exceeded, please try again later")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	// Parse similar structure for popular hotels
	var apiResponse struct {
		Data []struct {
			ID      string  `json:"id"`
			Name    string  `json:"name"`
			Address string  `json:"address"`
			Stars   int     `json:"stars"`
			Rating  float64 `json:"rating"`
			Price   struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"price"`
			Image       string   `json:"main_image_url"`
			Description string   `json:"description"`
			Amenities   []string `json:"amenities"`
			Location    struct {
				Lat float64 `json:"latitude"`
				Lng float64 `json:"longitude"`
			} `json:"location"`
		} `json:"data"`
		Success bool `json:"success"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	// Convert to our format
	var hotels []Hotel
	for _, item := range apiResponse.Data {
		hotel := Hotel{
			ID:          item.ID,
			Name:        item.Name,
			Address:     item.Address,
			Stars:       item.Stars,
			Rating:      item.Rating,
			PriceFrom:   item.Price.Amount,
			Currency:    item.Price.Currency,
			ImageURL:    item.Image,
			Description: item.Description,
			Amenities:   item.Amenities,
			Location: struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			}{
				Lat: item.Location.Lat,
				Lng: item.Location.Lng,
			},
		}
		hotels = append(hotels, hotel)
	}

	return &PopularHotelsResponse{
		Hotels: hotels,
		Status: "success",
	}, nil
}

// GetDestinations searches for destinations
func (s *travelPayoutsService) GetDestinations(query string) (*DestinationsResponse, error) {
	if s.apiKey == "" {
		return nil, fmt.Errorf("TravelPayouts API key is not configured")
	}

	// Build URL with parameters
	params := url.Values{}
	params.Set("query", query)
	params.Set("limit", "10")

	apiURL := fmt.Sprintf("https://autocomplete.travelpayouts.com/places2?%s", params.Encode())

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Access-Token", s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("invalid API key or authentication failed")
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("API rate limit exceeded, please try again later")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	// TravelPayouts destinations API returns a different structure
	var apiResponse []struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Country string `json:"country_name"`
		IATA    string `json:"code"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	// Convert to our format
	var destinations []Destination
	for _, item := range apiResponse {
		destinations = append(destinations, Destination{
			ID:      item.ID,
			Name:    item.Name,
			Country: item.Country,
			IATA:    item.IATA,
		})
	}

	return &DestinationsResponse{
		Destinations: destinations,
		Status:       "success",
	}, nil
}
