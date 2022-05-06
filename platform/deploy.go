package platform

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	sdk "github.com/hashicorp/waypoint-plugin-sdk/proto/gen"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type DeployConfig struct {
}

type Platform struct {
	config DeployConfig
}

// Implement Configurable
func (p *Platform) Config() (interface{}, error) {
	return &p.config, nil
}

// Implement ConfigurableNotify
func (p *Platform) ConfigSet(config interface{}) error {
	_, ok := config.(*DeployConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *DeployConfig as parameter")
	}

	return nil
}

// Implement Builder
func (p *Platform) DeployFunc() interface{} {
	return p.deploy
}

func (p *Platform) StatusFunc() interface{} {
	return p.status
}

func (b *Platform) deploy(
	ctx context.Context,
	ui terminal.UI,
	log hclog.Logger,
	dcr *component.DeclaredResourcesResp,
) (*Deployment, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Step skipped")

	return &Deployment{}, nil
}

func (d *Platform) status(
	ctx context.Context,
	ji *component.JobInfo,
	ui terminal.UI,
	log hclog.Logger,
	deployment *Deployment,
) (*sdk.StatusReport, error) {
	sg := ui.StepGroup()
	s := sg.Add("Step skipped")
	s.Update("Step skipped")
	s.Done()

	return &sdk.StatusReport{}, nil
}
