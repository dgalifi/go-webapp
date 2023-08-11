package dummy_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgalifi/go-webapp/pkg/config"
	"github.com/dgalifi/go-webapp/pkg/server"
	"github.com/dgalifi/go-webapp/pkg/services/dummy"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDummy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dummy Suite")
}

var _ = Describe("On", func() {
	var eng *gin.Engine
	var cfg config.Config

	BeforeEach(func() {
		cfg = config.Config{
			GreetingsMessage: "Test",
		}

		ds := dummy.NewDummyService(cfg)

		s := server.NewServer(cfg, ds)
		eng = s.GetEngine()
	})

	Context("When", func() {
		It("should", func() {
			requestURL := "/dummy"

			req := httptest.NewRequest(http.MethodGet, requestURL, nil)
			res := httptest.NewRecorder()

			// make httpRequest
			eng.ServeHTTP(res, req)

			var resBody map[string]interface{}

			Expect(res.Code).To(Equal(http.StatusOK))
			json.Unmarshal(res.Body.Bytes(), &resBody)

			Expect(resBody["message"]).To(Equal(cfg.GreetingsMessage))
		})
	})
})
