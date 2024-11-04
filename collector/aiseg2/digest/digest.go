package digest

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Request represents a client for Digest authentication
type Request struct {
	user      string
	password  string
	realm     string
	nonce     string
	algorithm string
	nc        int
	cnonce    string
}

// New creates a new DigestClient
func New(user, password, algorithm string) *Request {
	return &Request{
		user:      user,
		password:  password,
		algorithm: algorithm,
		nc:        0,
	}
}

// md5Hash calculates MD5 hash
func md5Hash(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// getCnonce generates a client nonce
func getCnonce() string {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

// ParseChallenge parses the WWW-Authenticate header
func (r *Request) ParseChallenge(header string) {
	header = strings.TrimPrefix(header, "Digest ")
	fields := strings.Split(header, ",")
	for _, field := range fields {
		v := strings.TrimSpace(field)
		parts := strings.SplitN(v, "=", 2)
		key := strings.Trim(parts[0], `"`)
		value := strings.Trim(parts[1], `"`)
		switch key {
		case "realm":
			r.realm = value
		case "nonce":
			r.nonce = value
		}
	}
	r.cnonce = getCnonce()
}

// ComputeResponse computes the Digest response
func (r *Request) ComputeResponse(method, uri string) string {
	ha1 := md5Hash(fmt.Sprintf("%s:%s:%s", r.user, r.realm, r.password))
	ha2 := md5Hash(fmt.Sprintf("%s:%s", method, uri))
	r.nc++
	ncValue := fmt.Sprintf("%08x", r.nc)
	response := md5Hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s",
		ha1, r.nonce, ncValue, r.cnonce, "auth", ha2))
	return response
}

// Do perform an HTTP request with Digest authentication
func (r *Request) Do(ctx context.Context, method, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}

	// Initial request to get the WWW-Authenticate header
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		return resp, nil
	}

	authHeader := resp.Header.Get("WWW-Authenticate")
	r.ParseChallenge(authHeader)

	// Create the Digest authorization header
	uri := req.URL.RequestURI()
	response := r.ComputeResponse(method, uri)
	authValue := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", algorithm="%s", response="%s", qop="auth", nc="%08x", cnonce="%s"`,
		r.user, r.realm, r.nonce, uri, r.algorithm, response, r.nc, r.cnonce)
	req.Header.Set("Authorization", authValue)

	// Resend the request with the Authorization header
	return httpClient.Do(req)
}
