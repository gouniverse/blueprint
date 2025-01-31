package config

import "github.com/gouniverse/cms"

func customEntityList() []cms.CustomEntityStructure {
	testimonial := cms.CustomEntityStructure{
		Type:      "testimonial",
		TypeLabel: "Testimonial",
		Group:     "Testimonials",
		AttributeList: []cms.CustomAttributeStructure{
			{
				Name:             "first_name",
				Type:             "string",
				FormControlLabel: "First name",
				FormControlType:  "input",
				FormControlHelp:  "The first name of the user",
			},
			{
				Name:             "last_name",
				Type:             "string",
				FormControlLabel: "Last name",
				FormControlType:  "input",
				FormControlHelp:  "The last name of the user",
			},
			{
				Name:             "job_title",
				Type:             "string",
				FormControlLabel: "Job title",
				FormControlType:  "input",
				FormControlHelp:  "The job title of the user",
			},
			{
				Name:             "quote",
				Type:             "string",
				FormControlLabel: "Quote",
				FormControlType:  "textarea",
				FormControlHelp:  "The testimonial quote",
			},
			{
				Name:             "image_url",
				Type:             "string",
				FormControlLabel: "Image (Base64)",
				FormControlType:  "textarea",
				FormControlHelp:  "The image of the person (base64 encoded)",
			},

			{
				Name:             "date",
				Type:             "string",
				FormControlLabel: "Date",
				FormControlType:  "string",
				FormControlHelp:  "The date of the testimonial",
			},
		},
	}

	list := []cms.CustomEntityStructure{}
	list = append(list, testimonial)

	return list
}
