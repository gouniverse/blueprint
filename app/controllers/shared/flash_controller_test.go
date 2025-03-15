package shared

import (
	"net/http"
	"net/url"
	"project/config"
	"project/internal/helpers"
	"project/internal/testutils"
	"strings"
	"testing"

	"github.com/gouniverse/utils"
)

func TestFlash(t *testing.T) {
	testutils.Setup()
	if config.CmsStore != nil {
		err := testutils.SeedTemplate(testutils.CMS_SITE_01, config.CmsUserTemplateID)

		if err != nil {
			t.Fatal(err)
		}
	}

	body, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewFlashController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"type":    {"success"},
			"message": {"Authentication Provider Error. Once is required field"},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`The message is no longer available`,
		`<a href="/">Click here to continue</a>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(body, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, body)
		}
	}
}

func TestFlashMessage_Info(t *testing.T) {
	testutils.Setup()
	if config.CmsStore != nil {
		err := testutils.SeedTemplate(testutils.CMS_SITE_01, config.CmsUserTemplateID)

		if err != nil {
			t.Fatal(err)
		}
	}

	infoUrl := helpers.ToFlashInfoURL("This is an info message", "/testbackendpoint", 5)

	flashMessageID := utils.StrRightFrom(infoUrl, `/flash?message_id=`)

	body, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewFlashController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"message_id": {flashMessageID},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`<div class="alert alert-info">`,
		`This is an info message`,
		`<a href="/testbackendpoint">Click here to continue</a>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(body, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, body)
		}
	}
}

func TestFlashMessage_Error(t *testing.T) {
	testutils.Setup()
	if config.CmsStore != nil {
		err := testutils.SeedTemplate(testutils.CMS_SITE_01, config.CmsUserTemplateID)

		if err != nil {
			t.Fatal(err)
		}
	}

	errorUrl := helpers.ToFlashErrorURL("This is an error message", "/testbackendpoint", 5)

	flashMessageID := utils.StrRightFrom(errorUrl, `/flash?message_id=`)

	body, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewFlashController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"message_id": {flashMessageID},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`<div class="alert alert-danger">`,
		`This is an error message`,
		`<a href="/testbackendpoint">Click here to continue</a>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(body, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, body)
		}
	}
}

func TestFlashMessage_Success(t *testing.T) {
	testutils.Setup()
	if config.CmsStore != nil {
		err := testutils.SeedTemplate(testutils.CMS_SITE_01, config.CmsUserTemplateID)

		if err != nil {
			t.Fatal(err)
		}
	}

	successUrl := helpers.ToFlashSuccessURL("This is a success message", "/testbackendpoint", 5)

	flashMessageID := utils.StrRightFrom(successUrl, `/flash?message_id=`)

	body, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewFlashController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"message_id": {flashMessageID},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`<div class="alert alert-success">`,
		`This is a success message`,
		`<a href="/testbackendpoint">Click here to continue</a>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(body, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, body)
		}
	}
}

func TestFlashMessage_Warning(t *testing.T) {
	testutils.Setup()
	if config.CmsStore != nil {
		err := testutils.SeedTemplate(testutils.CMS_SITE_01, config.CmsUserTemplateID)

		if err != nil {
			t.Fatal(err)
		}
	}

	warningUrl := helpers.ToFlashWarningURL("This is a warning message", "/testbackendpoint", 5)

	flashMessageID := utils.StrRightFrom(warningUrl, `/flash?message_id=`)

	body, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewFlashController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"message_id": {flashMessageID},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`<div class="alert alert-warning">`,
		`This is a warning message`,
		`<a href="/testbackendpoint">Click here to continue</a>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(body, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, body)
		}
	}
}

func TestFlashMessage_Get(t *testing.T) {
	testutils.Setup()
	if config.CmsStore != nil {
		err := testutils.SeedTemplate(testutils.CMS_SITE_01, config.CmsUserTemplateID)

		if err != nil {
			t.Fatal(err)
		}
	}

	infoUrl := helpers.ToFlashInfoURL("This is an info message", "/testbackendpoint", 5)

	flashMessageID := utils.StrRightFrom(infoUrl, `/flash?message_id=`)

	body, response, err := testutils.CallHtmlEndpoint(http.MethodGet, NewFlashController().Handler, testutils.NewRequestOptions{
		GetValues: url.Values{
			"message_id": {flashMessageID},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`<div class="alert alert-info">`,
		`This is an info message`,
		`<a href="/testbackendpoint">Click here to continue</a>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(body, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, body)
		}
	}
}

func TestFlashMessage_Delete(t *testing.T) {
	testutils.Setup()
	if config.CmsStore != nil {
		err := testutils.SeedTemplate(testutils.CMS_SITE_01, config.CmsUserTemplateID)

		if err != nil {
			t.Fatal(err)
		}
	}

	infoUrl := helpers.ToFlashInfoURL("This is an info message", "/testbackendpoint", 5)

	flashMessageID := utils.StrRightFrom(infoUrl, `/flash?message_id=`)

	body, response, err := testutils.CallHtmlEndpoint(http.MethodDelete, NewFlashController().Handler, testutils.NewRequestOptions{
		GetValues: url.Values{
			"message_id": {flashMessageID},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`<div class="alert alert-info">`,
		`This is an info message`,
		`<a href="/testbackendpoint">Click here to continue</a>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(body, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, body)
		}
	}
}

func TestFlashMessage_Post(t *testing.T) {
	testutils.Setup()
	if config.CmsStore != nil {
		err := testutils.SeedTemplate(testutils.CMS_SITE_01, config.CmsUserTemplateID)

		if err != nil {
			t.Fatal(err)
		}
	}

	infoUrl := helpers.ToFlashInfoURL("This is an info message", "/testbackendpoint", 5)

	flashMessageID := utils.StrRightFrom(infoUrl, `/flash?message_id=`)

	body, response, err := testutils.CallHtmlEndpoint(http.MethodPost, NewFlashController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"message_id": {flashMessageID},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`<div class="alert alert-info">`,
		`This is an info message`,
		`<a href="/testbackendpoint">Click here to continue</a>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(body, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, body)
		}
	}
}

func TestFlashMessage_Put(t *testing.T) {
	testutils.Setup()
	if config.CmsStore != nil {
		err := testutils.SeedTemplate(testutils.CMS_SITE_01, config.CmsUserTemplateID)

		if err != nil {
			t.Fatal(err)
		}
	}

	infoUrl := helpers.ToFlashInfoURL("This is an info message", "/testbackendpoint", 5)

	flashMessageID := utils.StrRightFrom(infoUrl, `/flash?message_id=`)

	body, response, err := testutils.CallHtmlEndpoint(http.MethodPut, NewFlashController().Handler, testutils.NewRequestOptions{
		PostValues: url.Values{
			"message_id": {flashMessageID},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal(`Response MUST be `, http.StatusOK, ` but was: `, response.StatusCode)
	}

	expecteds := []string{
		`<div class="alert alert-info">`,
		`This is an info message`,
		`<a href="/testbackendpoint">Click here to continue</a>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(body, expected) {
			t.Fatal(`Response MUST contain`, expected, ` but was `, body)
		}
	}
}
