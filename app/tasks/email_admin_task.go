package tasks

import (
	"errors"
	"project/config"
	"project/internal/emails"

	"github.com/gouniverse/taskstore"
)

// NewEmailToAdminTask sends a notification email to admin
// =================================================================
// Example:
//
// go run . task email-to-admin --html=HTML
//
// =================================================================
func NewEmailToAdminTaskHandler() *emailToAdminTaskHandler {
	return &emailToAdminTaskHandler{}
}

// emailToAdminTaskHandler sends a notification email to admin
type emailToAdminTaskHandler struct {
	taskstore.TaskHandlerBase
}

var _ taskstore.TaskHandlerInterface = (*emailToAdminTaskHandler)(nil) // verify it extends the task interface

func (handler *emailToAdminTaskHandler) Alias() string {
	return "email-to-admin"
}

func (handler *emailToAdminTaskHandler) Title() string {
	return "Email to Admin"
}

func (handler *emailToAdminTaskHandler) Description() string {
	return "Sends a notofication email to admin"
}

func (handler *emailToAdminTaskHandler) Enqueue(html string) (task taskstore.QueueInterface, err error) {
	if config.TaskStore == nil {
		return nil, errors.New("task store is nil")
	}

	return config.TaskStore.TaskEnqueueByAlias(handler.Alias(), map[string]any{
		"html": html,
	})
}

func (handler *emailToAdminTaskHandler) Handle() bool {
	html := handler.GetParam("html")

	if html == "" {
		handler.LogError("html is required parameter")
		return false
	}

	if !handler.HasQueuedTask() && handler.GetParam("enqueue") == "yes" {
		_, err := handler.Enqueue(html)

		if err != nil {
			handler.LogError("Error enqueuing task: " + err.Error())
		} else {
			handler.LogSuccess("Task enqueued.")
		}

		return true
	}

	handler.LogInfo("Parameters ok ...")

	err := emails.NewEmailNotifyAdmin().Send(html)

	if err != nil {
		handler.LogError("Sending email failed. Code: ")
		handler.LogError("Error: " + err.Error())
		return false
	}

	handler.LogSuccess("Sending email OK.")

	return true
}
