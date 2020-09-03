// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package victoriametrics provides facilities for working with VMAlert.
package victoriametrics

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/percona/pmm-managed/services/prometheus"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
)




// VMAlert is responsible for interactions with victoria metrics.
type VMAlert struct {
	db               *reform.DB
	baseURL          *url.URL
	client           *http.Client
	alertingRules *prometheus.AlertingRules

	cachedAlertingRules string

	l    *logrus.Entry
	sema chan struct{}
}

// NewVMAlert creates new Victoria Metrics service.
func NewVMAlert(alertRules *prometheus.AlertingRules, db *reform.DB, baseURL string) (*VMAlert, error) {
	if !Enabled() {
		return &VMAlert{}, nil
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &VMAlert{
		alertingRules: alertRules,
		db:               db,
		baseURL:          u,
		client:           new(http.Client),
		l:                logrus.WithField("component", "vmalert"),
		sema:             make(chan struct{}, 1),
	}, nil
}

// Run runs VMAlert configuration update loop until ctx is canceled.
func (svc *VMAlert) Run(ctx context.Context) {
	if !Enabled() {
		return
	}
	svc.l.Info("Starting...")
	defer svc.l.Info("Done.")
	alertingRules, err := svc.alertingRules.ReadRules()
	if err != nil {
		svc.l.Warnf("Cannot load alerting rules: %s", err)
	}
	svc.cachedAlertingRules = alertingRules
	for {
		select {
		case <-ctx.Done():
			return

		case <-svc.sema:
			// batch several update requests together by delaying the first one
			sleepCtx, sleepCancel := context.WithTimeout(ctx, updateBatchDelay)
			<-sleepCtx.Done()
			sleepCancel()

			if ctx.Err() != nil {
				return
			}

			if err := svc.updateConfiguration(); err != nil {
				svc.l.Errorf("Failed to update configuration, will retry: %+v.", err)
				svc.RequestConfigurationUpdate()
			}
		}
	}
}

// updateConfiguration updates Prometheus configuration.
func (svc *VMAlert) updateConfiguration() error {
	start := time.Now()
	defer func() {
		if dur := time.Since(start); dur > time.Second {
			svc.l.Warnf("updateConfiguration took %s.", dur)
		}
	}()

	return svc.saveConfigAndReload()
}

// RequestConfigurationUpdate requests Prometheus configuration update.
func (svc *VMAlert) RequestConfigurationUpdate() {
	if !Enabled() {
		return
	}
	select {
	case svc.sema <- struct{}{}:
		err := svc.saveConfigAndReload()
		if err != nil {
			svc.l.WithError(err).Errorf("cannot reload configuration")
		}
	default:
	}
}

// IsReady verifies that VMAlert works.
func (svc *VMAlert) IsReady(ctx context.Context) error {
	if !Enabled() {
		return nil
	}
	// check VMAlert /health API and log version
	u := *svc.baseURL
	u.Path = path.Join(u.Path, "health")
	resp, err := svc.client.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck
	b, err := ioutil.ReadAll(resp.Body)
	svc.l.Debugf("VMAlert: %s", b)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.Errorf("expected 200, got %d", resp.StatusCode)
	}

	svc.l.Debugf("%s", b)
	return nil
}

// reload asks VMAlert to reload configuration.
func (svc *VMAlert) reload() error {
	u := *svc.baseURL
	u.Path = path.Join(u.Path, "-", "reload")
	resp, err := svc.client.Post(u.String(), "", nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.Errorf("%d: %s", resp.StatusCode, b)
}

// saveConfigAndReload saves given VMAlert configuration to file and reloads VMAlert.
// If configuration can't be reloaded for some reason, old file is restored, and configuration is reloaded again.
func (svc *VMAlert) saveConfigAndReload() error {
	// read existing content
	oldCfg, err := svc.alertingRules.ReadRules()
	if err != nil {
		return errors.WithStack(err)
	}

	// compare with new config
	if  oldCfg == svc.cachedAlertingRules {
		svc.l.Infof("Configuration not changed, doing nothing.")
		return nil
	}
	err = svc.reload()
	if err != nil {
		return errors.WithStack(err)
	}
	svc.l.Infof("Configuration reloaded.")

	return nil
}
