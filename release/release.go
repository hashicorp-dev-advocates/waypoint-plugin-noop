package release

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	sdk "github.com/hashicorp/waypoint-plugin-sdk/proto/gen"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type ReleaseConfig struct {
}

type ReleaseManager struct {
	config ReleaseConfig
}

// Implement Configurable
func (rm *ReleaseManager) Config() (interface{}, error) {
	return &rm.config, nil
}

// Implement ConfigurableNotify
func (rm *ReleaseManager) ConfigSet(config interface{}) error {
	_, ok := config.(*ReleaseConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *ReleaseConfig as parameter")
	}

	return nil
}

// Implement Builder
func (rm *ReleaseManager) ReleaseFunc() interface{} {
	// return a function which will be called by Waypoint
	return rm.release
}

func (rm *ReleaseManager) StatusFunc() interface{} {
	return rm.status
}

func (rm *ReleaseManager) release(
	ctx context.Context,
	log hclog.Logger,
	dcr *component.DeclaredResourcesResp,
	ui terminal.UI,
) (*Release, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Skipped step")

	return &Release{}, nil
}

func (rm *ReleaseManager) status(
	ctx context.Context,
	ji *component.JobInfo,
	log hclog.Logger,
	ui terminal.UI,
	release *Release,
) (*sdk.StatusReport, error) {
	sg := ui.StepGroup()
	s := sg.Add("Skipped step")

	s.Update("Skipped step")
	s.Done()

	return &sdk.StatusReport{}, nil
}
