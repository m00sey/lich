package sat

import (
	"context"
	"log"

	"net/http"

	"strings"

	"github.com/lich/swagger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {

	var (
		api *swagger.APIClient
	)

	BeforeEach(func() {
		api = swagger.NewAPIClient(swagger.NewConfiguration())
	})

	AfterEach(func() {

	})

	Context("", func() {
		BeforeEach(func() {

		})

		It("", func() {

			body := `{
				"User":[{

				}]
			}`

			c := http.Client{}
			req, err := http.NewRequest(http.MethodPost, "http://local.scoir.com:9001/integration/load",
				strings.NewReader(body))
			Expect(err).ToNot(HaveOccurred())

			resp, err := c.Do(req)
			Expect(err).ToNot(HaveOccurred())

			log.Println(resp)

			sat := swagger.Sat{

				TestDate:     "10/10/2010",
				MathScore:    11,
				ReadingScore: 12,
				WritingScore: 13,
			}

			ctx := context.TODO()
			ctx = context.WithValue(ctx, swagger.ContextAccessToken, "")

			sat, resp, err := api.DefaultApi.HighSchoolHighSchoolIdStudentStudentIdSATPost(ctx, "", "", sat)
			Expect(err).ShouldNot(HaveOccurred())

			log.Println(sat, resp)
		})
	})
})
