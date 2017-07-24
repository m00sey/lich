package sat

import (
	"context"
	"log"

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
