package widgets

import (
	"net/http"
	"project/pkg/testimonials"

	"github.com/gouniverse/hb"
	"github.com/samber/lo"
)

var _ Widget = (*testimonialsWidget)(nil) // verify it extends the interface

// == CONSTUCTOR ==============================================================

// NewPrintWidget creates a new instance of the print struct.
//
// Parameters:
//   - None
//
// Returns:
//   - *print - A pointer to the print struct
func NewTestimonialsWidget() *testimonialsWidget {
	return &testimonialsWidget{}
}

// == WIDGET ================================================================

// print is the struct that will be used to render the print shortcode.
//
// This shortcode is used to evaluate the result of the provided content
// and return it.
//
// It uses Otto as the engine.
type testimonialsWidget struct{}

// == PUBLIC METHODS =========================================================

// Alias the shortcode alias to be used in the template.
func (t *testimonialsWidget) Alias() string {
	return "x-testimonials"
}

// Description a user-friendly description of the shortcode.
func (t *testimonialsWidget) Description() string {
	return "Renders the testimonials"
}

// Render implements the shortcode interface.
func (t *testimonialsWidget) Render(r *http.Request, content string, params map[string]string) string {
	testimonialList, err := testimonials.TestimonialList()

	if err != nil {
		return "Error: " + err.Error()
	}

	if len(testimonialList) < 1 {
		return "No testimonials found"
	}

	row := hb.Div().
		Class("row").
		Children(lo.Map(testimonialList, func(testimonial testimonials.Testimonial, index int) hb.TagInterface {
			// imageUrl := testimonial.ImageUrl()
			// imageUrl = links.NewWebsiteLinks().Thumbnail("jpg", "200", "200", "80", imageUrl)

			// image := hb.Img("").
			// 	Class("Image rounded-start").
			// 	Style("width: 100%; height: 100%; object-fit: cover;").
			// 	Style(`background-position: center; background-size: cover; background-repeat: no-repeat;`).
			// 	Style("background-image:url(" + imageUrl + ");")

			stars := hb.Wrap().
				Child(hb.I().Class("bi bi-star-fill")).
				Child(hb.I().Class("bi bi-star-fill")).
				Child(hb.I().Class("bi bi-star-fill")).
				Child(hb.I().Class("bi bi-star-fill")).
				Child(hb.I().Class("bi bi-star-fill"))

			name := testimonial.FirstName()
			if len(testimonial.LastName()) > 0 {
				name += " "
				name += testimonial.LastName()[0:1]
				name += "."
			}

			card := hb.Div().
				Class("card mb-3").
				Child(hb.Div().
					Class("card-body").
					Child(hb.H5().
						Class("card-title").
						Child(stars)).
					Child(hb.P().
						Class("card-text").
						Child(hb.Text(testimonial.Quote()))).
					Child(hb.P().
						Class("card-text").
						Text(name).
						Text(",")).
					Text(" ").
					Text(testimonial.JobTitle()))

			return hb.Div().
				Class("col-sm-6").
				Child(card)
		}))

	// <div class="row">
	// 		<div class="col-sm-6">
	// 			<div class="card mb-3">
	// 			  <div class="row g-0">
	// 				<div class="col-md-4">
	// 					<div class="Image rounded-start" style="background-image:url(/media/home-page/testimonials-user-1.png);"></div>
	// 				</div>
	// 				<div class="col-md-8">
	// 				  <div class="card-body">
	// 					<h5 class="card-title">
	// 						<i class="bi bi-star-fill"></i>
	// 						<i class="bi bi-star-fill"></i>
	// 						<i class="bi bi-star-fill"></i>
	// 						<i class="bi bi-star-fill"></i>
	// 						<i class="bi bi-star-fill"></i>
	// 					</h5>
	// 					<p class="card-text">
	// 						Roast My Contract saved me from a major headache!
	// 						I was about to sign a partnership agreement with some confusing clauses.
	// 						The AI review flagged them, and I was able to renegotiate for a much fairer deal.
	// 						Now I sleep soundly knowing everything is clear and above board!
	// 					</p>
	// 					<p class="card-text">
	// 						<small class="text-body-secondary">Sarah J., Entrepreneur</small>
	// 					  </p>
	// 				  </div>
	// 				</div>
	// 			  </div>
	// 			</div>
	// 		</div>

	return row.ToHTML()

	// path := r.URL.Path

	// vm := otto.New()

	// vm.Set("path", path)

	// result, err := vm.Run("result = " + content)

	// if err != nil {
	// 	cfmt.Errorln(err)
	// }

	// return result.String()
}
