package widgets

// WidgetRegistry returns a list of all widgets
//
// Register all the new widgets here so that they can be used in the CMS
//
// Parameters:
//   - None
//
// Returns:
//   - []Widget - A list of all widgets
func WidgetRegistry() []Widget {
	return []Widget{
		NewAuthenticatedWidget(),
		NewContactFormWidget(),
		NewUnauthenticatedWidget(),
		NewVisibleWidget(),
	}
}
