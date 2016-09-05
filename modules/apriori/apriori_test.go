package apriori

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestApriori(t *testing.T) {

	Convey("Given a list of transactions", t, func() {
		data := [][]string{
			[]string{"iphone"},
			[]string{"pencil"},
			[]string{"motorolla"},
			[]string{"bike"},
			[]string{"pencil", "beats"},
			[]string{"iphone", "iphone-case"},
			[]string{"iphone-case", "carro", "bike", "helmet", "gloves", "jacket", "lamp"},
			[]string{"iphone", "iphone-case", "beats", "notebook", "pen", "desk", "laptop", "power bank"},
			[]string{"iphone", "iphone-case", "power bank"},
			[]string{"iphone", "iphone-case", "power bank"},
		}

		first := data[0]
		ap := Result(data)

		Convey("For validate that pencil es subset of [pencil beats]", func() {
			ss := Subset(data[1], data[4])
			Convey("The output should be true", func() {
				So(ss, ShouldEqual, true)
			})
		})

		Convey("When the support for the first item is calculated", func() {
			sup := Support(first, data)
			Convey("The support should be 0.5", func() {
				So(sup, ShouldEqual, 0.5)
			})
		})

		Convey("When the apriori algorithm is calculated for the input data", func() {
			Convey("The first combination with highest support should be contain iphone item", func() {
				So(ap[1], ShouldContain, "iphone")
			})
		})

	})
}
