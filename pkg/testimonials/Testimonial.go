package testimonials

import (
	"errors"

	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/entitystore"
)

const ENTITY_TYPE = "testimonial"

type Testimonial struct {
	dataobject.DataObject
}

func NewTestimonial() *Testimonial {
	return &Testimonial{}
}

func NewTestimonialFromEntity(entity entitystore.Entity) (*Testimonial, error) {
	if entity.Type() != ENTITY_TYPE {
		return nil, errors.New("invalid entity type")
	}

	attributes, err := entity.GetAttributes()

	if err != nil {
		return nil, err
	}

	testimonial := NewTestimonial()

	for _, attribute := range attributes {
		testimonial.Set(attribute.AttributeKey(), attribute.AttributeValue())
	}

	return testimonial, nil
}

// == SETTERS AND GETTERS =====================================================

func (t *Testimonial) Date() string {
	return t.Get("date")
}

func (t *Testimonial) SetDate(date string) {
	t.Set("date", date)
}

func (t *Testimonial) FirstName() string {
	return t.Get("first_name")
}

func (t *Testimonial) SetFirstName(firstName string) {
	t.Set("first_name", firstName)
}

func (t *Testimonial) ID() string {
	return t.Get("id")
}

func (t *Testimonial) SetID(id string) {
	t.Set("id", id)
}

func (t *Testimonial) ImageUrl() string {
	return t.Get("image_url")
}

func (t *Testimonial) SetImageUrl(imageUrl string) {
	t.Set("image_url", imageUrl)
}

func (t *Testimonial) JobTitle() string {
	return t.Get("job_title")
}

func (t *Testimonial) SetJobTitle(jobTitle string) {
	t.Set("job_title", jobTitle)
}

func (t *Testimonial) LastName() string {
	return t.Get("last_name")
}

func (t *Testimonial) SetLastName(lastName string) {
	t.Set("last_name", lastName)
}

func (t *Testimonial) Quote() string {
	return t.Get("quote")
}

func (t *Testimonial) SetQuote(quote string) {
	t.Set("quote", quote)
}

func (t *Testimonial) Status() string {
	return t.Get("status")
}

func (t *Testimonial) SetStatus(status string) {
	t.Set("status", status)
}

func (t *Testimonial) CreatedAt() string {
	return t.Get("created_at")
}

func (t *Testimonial) SetCreatedAt(createdAt string) {
	t.Set("created_at", createdAt)
}

func (t *Testimonial) UpdatedAt() string {
	return t.Get("updated_at")
}

func (t *Testimonial) SetUpdatedAt(updatedAt string) {
	t.Set("updated_at", updatedAt)
}
