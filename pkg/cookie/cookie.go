package cookie

import (
	"net/http"
	"strings"
)

// formatDomain formats the domain.
func formatDomain(domain string) string {
	if domain != "localhost" && !strings.HasPrefix(domain, ".") {
		return "." + domain
	}
	return domain
}

// Set sets cookies.
func Set(w http.ResponseWriter, accessToken, refreshToken, domain string) {
	formattedDomain := formatDomain(domain)
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		MaxAge:   60 * 60 * 24,
		Path:     "/",
		Domain:   formattedDomain,
		Secure:   true,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		MaxAge:   60 * 60 * 24 * 30,
		Path:     "/",
		Domain:   formattedDomain,
		Secure:   true,
		HttpOnly: true,
	})
}

// SetRegister sets registration cookies.
func SetRegister(w http.ResponseWriter, registerToken, domain string) {
	formattedDomain := formatDomain(domain)
	http.SetCookie(w, &http.Cookie{
		Name:     "register_token",
		Value:    registerToken,
		MaxAge:   60 * 60,
		Path:     "/",
		Domain:   formattedDomain,
		Secure:   true,
		HttpOnly: true,
	})
}

// Clear clears cookies.
func Clear(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "access_token",
		MaxAge: -1,
		Path:   "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "refresh_token",
		MaxAge: -1,
		Path:   "/",
	})
}

// ClearRegister clears registration cookies.
func ClearRegister(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "register_token",
		MaxAge: -1,
		Path:   "/",
	})
}

// ClearAll clears all cookies.
func ClearAll(w http.ResponseWriter) {
	Clear(w)
	ClearRegister(w)
}

// Get gets cookies.
func Get(r *http.Request, key string) (string, error) {
	cookie, err := r.Cookie(key)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// GetRegister gets registration cookies.
func GetRegister(r *http.Request, key string) (string, error) {
	cookie, err := r.Cookie(key)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
