package testimonials

import (
	"project/config"

	"github.com/gouniverse/entitystore"
)

func TestimonialList() ([]Testimonial, error) {
	result, err := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
		EntityType: "testimonial",
	})

	if err != nil {
		return nil, err
	}

	testimonials := []Testimonial{}

	for _, entry := range result {
		testimonial, err := NewTestimonialFromEntity(entry)

		if err != nil {
			continue
		}

		testimonials = append(testimonials, *testimonial)
	}

	return testimonials, nil
}
