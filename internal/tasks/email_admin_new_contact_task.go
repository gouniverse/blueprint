package tasks

import (
	"errors"
	"project/config"
	"project/internal/emails"

	"github.com/gouniverse/taskstore"
)

// NewEmailToAdminOnNewContactFormSubmittedTaskHandler sends a notification email to admin
// =================================================================
// Example:
//
// go run . task email-to-admin-on-new-contact-form-submitted --html=HTML
//
// =================================================================
func NewEmailToAdminOnNewContactFormSubmittedTaskHandler() *emailToAdminOnNewContactFormSubmittedTaskHandler {
	return &emailToAdminOnNewContactFormSubmittedTaskHandler{}
}

// emailToAdminOnNewContactFormSubmittedTaskHandler sends a notification email to admin
type emailToAdminOnNewContactFormSubmittedTaskHandler struct {
	taskstore.TaskHandlerBase
}

var _ taskstore.TaskHandlerInterface = (*emailToAdminOnNewContactFormSubmittedTaskHandler)(nil) // verify it extends the task interface

func (handler *emailToAdminOnNewContactFormSubmittedTaskHandler) Alias() string {
	return "email-to-admin-on-new-contact-form-submitted"
}

func (handler *emailToAdminOnNewContactFormSubmittedTaskHandler) Title() string {
	return "Email to Admin on New Contact"
}

func (handler *emailToAdminOnNewContactFormSubmittedTaskHandler) Description() string {
	return "Sends a notofication email to admin when a new contact form is submitted"
}

func (handler *emailToAdminOnNewContactFormSubmittedTaskHandler) Enqueue() (task taskstore.QueueInterface, err error) {
	if config.TaskStore == nil {
		return nil, errors.New("task store is nil")
	}
	return config.TaskStore.TaskEnqueueByAlias(handler.Alias(), map[string]any{})
}

func (handler *emailToAdminOnNewContactFormSubmittedTaskHandler) Handle() bool {
	if !handler.HasQueuedTask() && handler.GetParam("enqueue") == "yes" {
		_, err := handler.Enqueue()

		if err != nil {
			handler.LogError("Error enqueuing task: " + err.Error())
		} else {
			handler.LogSuccess("Task enqueued.")
		}

		return true
	}

	handler.LogInfo("Parameters ok ...")

	err := emails.NewEmailToAdminOnNewContactFormSubmitted().Send()

	if err != nil {
		handler.LogError("Sending email failed. Code: ")
		handler.LogError("Error: " + err.Error())
		return false
	}

	handler.LogSuccess("Sending email OK.")

	return true
}
