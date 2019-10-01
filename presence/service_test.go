package presence

import (
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/nym-directory/models"
	"github.com/nymtech/nym-directory/presence/mocks"
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("presence.Service", func() {
	var (
		mix1      models.MixHostInfo
		presence1 models.MixNodePresence
		coco1     models.CocoHostInfo
		presence2 models.CocoPresence
		mockDb    mocks.IDb

		serv service
	)
	BeforeEach(func() {
		mockDb = *new(mocks.IDb)
		serv = *newService(&mockDb)
		var now = time.Now()
		timemock.Freeze(now)

		// Set up fixtures
		mix1 = models.MixHostInfo{
			HostInfo: models.HostInfo{
				Host:   "foo.com:8000",
				PubKey: "pubkey1",
			},
			Layer: 1,
		}

		presence1 = models.MixNodePresence{
			MixHostInfo: mix1,
			LastSeen:    timemock.Now().UnixNano(),
		}

		coco1 = models.CocoHostInfo{
			HostInfo: models.HostInfo{
				Host:   "bar.com:8000",
				PubKey: "pubkey2",
			},
			Type: "foo",
		}

		presence2 = models.CocoPresence{
			CocoHostInfo: coco1,
			LastSeen:     timemock.Now().UnixNano(),
		}
	})

	Describe("Adding presence info", func() {
		Context("for a mixnode", func() {
			It("should add a presence to the db", func() {
				mockDb.On("AddMix", presence1)
				serv.AddMixNodePresence(mix1)
				mockDb.AssertCalled(GinkgoT(), "AddMix", presence1)
			})
		})
		Context("for a coconode", func() {
			It("should add a presence to the db", func() {
				mockDb.On("AddCoco", presence2)
				serv.AddCocoNodePresence(coco1)
				mockDb.AssertCalled(GinkgoT(), "AddCoco", presence2)
			})
		})
	})
	Describe("Listing presence info", func() {
		Context("when receiving a list request", func() {
			It("should call to the Db", func() {
				list := []models.MixNodePresence{
					presence1,
				}
				topology := models.Topology{
					MixNodes: list,
				}
				mockDb.On("Topology").Return(topology)
				result := serv.Topology()
				mockDb.AssertCalled(GinkgoT(), "Topology")
				assert.Equal(GinkgoT(), topology, result)
			})
		})
	})
})
