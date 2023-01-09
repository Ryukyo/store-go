package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Ryukyo/store-go/internal/config"
	"github.com/Ryukyo/store-go/internal/models"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})

	testApp.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type testWriter struct{}

func (tw *testWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *testWriter) WriteHeader(i int) {}

func (tw *testWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
