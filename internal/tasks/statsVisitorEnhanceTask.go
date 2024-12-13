package tasks

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"project/config"
	"strings"

	"github.com/gouniverse/statsstore"
	"github.com/gouniverse/taskstore"
	"github.com/spf13/cast"

	"github.com/mingrammer/cfmt"

	"github.com/mileusna/useragent"
)

// statsVisitorEnhanceTask enhances the visitor stats with the country
//
// =================================================================
// Example:
//
// go run main.go task stats-visitor-enhance
//
// =================================================================
type statsVisitorEnhanceTask struct {
	taskstore.TaskHandlerBase
}

// == CONSTRUCTOR =============================================================

func NewStatsVisitorEnhanceTask() *statsVisitorEnhanceTask {
	return &statsVisitorEnhanceTask{}
}

// == IMPLEMENTATION ==========================================================

// var _ jobsshared.TaskInterface = (*statsVisitorEnhanceTask)(nil) // verify it extends the task interface
var _ taskstore.TaskHandlerInterface = (*statsVisitorEnhanceTask)(nil) // verify it extends the task interface

// == PUBLIC METHODS ==========================================================

func (t *statsVisitorEnhanceTask) Enqueue() (taskstore.QueueInterface, error) {
	if config.TaskStore == nil {
		return nil, errors.New("task store is nil")
	}
	return config.TaskStore.TaskEnqueueByAlias(t.Alias(), map[string]interface{}{})
}

func (t *statsVisitorEnhanceTask) Alias() string {
	return "StatsVisitorEnhance"
}

func (t *statsVisitorEnhanceTask) Title() string {
	return "Stats Visitor Enhance"
}

func (t *statsVisitorEnhanceTask) Description() string {
	return "Enhances the visitor stats by adding the country"
}

func (t *statsVisitorEnhanceTask) Handle() bool {
	if config.StatsStore == nil {
		t.LogError("Task StatsVisitorEnhance. Store is nil")
		return false
	}

	ctx := context.Background()
	unprocessedEntries, err := config.StatsStore.VisitorList(ctx, statsstore.VisitorQueryOptions{
		Country: "empty",
		Limit:   10,
	})

	if err != nil {
		t.LogError("Task StatsVisitorEnhance. Error: " + err.Error())
		return false
	}

	if len(unprocessedEntries) < 1 {
		t.LogInfo("Task StatsVisitorEnhance. No entries to process")
		return true
	}

	t.LogInfo("Task StatsVisitorEnhance. Found: " + cast.ToString(len(unprocessedEntries)) + " entries to process")

	for i := 0; i < len(unprocessedEntries); i++ {
		entry := unprocessedEntries[i]
		t.processVisitor(ctx, entry)
	}

	// cfmt.Infoln("Visitors countries populate")
	// fingerprintList, errFingerprintList := models.FunnelStore.FingerprintList(funnelstore.FingerprintQueryOptions{
	// 	Country: "empty",
	// 	Limit:   1,
	// })

	// if errFingerprintList != nil {
	// 	t.LogError("Task FingerprintCountryPopulate. Error: " + errFingerprintList.Error())
	// 	return false
	// }

	// if len(fingerprintList) < 1 {
	// 	t.LogInfo("Task FingerprintCountryPopulate. No fingerprints to process")
	// 	return true
	// }

	// t.LogInfo("TaskFingerprintCountryPopulate. Found: " + utils.ToString(len(fingerprintList)) + " fingerprints without countries to process")
	// for i := 0; i < len(fingerprintList); i++ {
	// 	fingerprint := fingerprintList[i]
	// 	country := taskFingerprintCountryFindCountryByIp(fingerprint.IpAddress())
	// 	t.LogInfo("Fingerprint: " + fingerprint.ID() + " : " + fingerprint.IpAddress() + " : " + country)
	// 	errUpdated := models.FunnelStore.FingerprintUpdateByID(fingerprint.ID(), funnelstore.FingerprintUpdateByIdParameters{
	// 		Country: country,
	// 	})
	// 	if errUpdated != nil {
	// 		cfmt.Errorln(errUpdated.Error())
	// 	}
	// }

	return false
}

// == PRIVATE METHODS =========================================================

func (t *statsVisitorEnhanceTask) processVisitor(ctx context.Context, visitor statsstore.VisitorInterface) bool {
	if config.StatsStore == nil {
		t.LogError("Task StatsVisitorEnhance. Store is nil")
		return false
	}
	ua := useragent.Parse(visitor.UserAgent())
	userOs := ua.OS
	userOsVersion := ua.OSVersion
	userDevice := ua.Device
	userBrowser := ua.Name
	userBrowserVersion := ua.Version

	userDeviceType := ""

	if ua.Mobile {
		userDeviceType = "mobile"
	}
	if ua.Tablet {
		userDeviceType = "tablet"
	}
	if ua.Desktop {
		userDeviceType = "desktop"
	}
	if ua.Bot {
		userDeviceType = "bot"
	}

	country := t.findCountryByIp(visitor.IpAddress())

	visitor.SetCountry(country)
	visitor.SetUserBrowser(userBrowser)
	visitor.SetUserBrowserVersion(userBrowserVersion)
	visitor.SetUserDevice(userDevice)
	visitor.SetUserDeviceType(userDeviceType)
	visitor.SetUserOs(userOs)
	visitor.SetUserOsVersion(userOsVersion)

	errUpdated := config.StatsStore.VisitorUpdate(ctx, visitor)

	if errUpdated != nil {
		cfmt.Errorln(errUpdated.Error())
	}

	return false
}

func (t *statsVisitorEnhanceTask) findCountryByIp(ip string) string {
	if ip == "" || ip == "127.0.0.1" {
		return "UN"
	}

	resp, err := http.Get("https://ip2c.org/" + ip)

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return "ER" // error
	}

	if resp == nil {
		return "UN"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return "ER"
	}
	// Log the request body
	bodyString := string(body)
	cfmt.Infoln(bodyString)
	parts := strings.Split(bodyString, ";")
	if len(parts) > 2 {
		if parts[1] == "" {
			return "UN"
		}
		return parts[1]
	}

	return "UN"
}
