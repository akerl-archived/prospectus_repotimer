package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/akerl/prospectus/v2/plugin"

	"gopkg.in/src-d/go-git.v4"
)

const (
	defaultDays = 7
)

type config struct {
	Days int `json:"days"`
}

type timerPlugin struct {
	Config config
}

func (p *timerPlugin) GetConfigPointer() interface{} {
	return &p.Config
}

func (p *timerPlugin) Load(_ plugin.LoadInput) plugin.AttributeSet {
	return plugin.AttributeSet{plugin.Attribute{Name: "repotimer"}}
}

func (p *timerPlugin) Check(input plugin.Attribute) plugin.Result {
	actualTime, err := p.getLastCommitTime(input.Dir)
	if err != nil {
		return plugin.NewErrorResult(err.Error())
	}
	currentTime := time.Now()
	actualDaysAgo := int(currentTime.Sub(actualTime).Hours() / 24)

	allowedDaysAgo := p.Config.Days
	if allowedDaysAgo == 0 {
		allowedDaysAgo = defaultDays
	}

	matches := true
	if actualDaysAgo > allowedDaysAgo {
		matches = false
	}

	var expected strings.Builder
	for i := 1; i < allowedDaysAgo; i++ {
		fmt.Fprintf(&expected, "%d, ", i)
	}
	fmt.Fprintf(&expected, "%d", allowedDaysAgo)

	return plugin.Result{
		Actual:   fmt.Sprintf("%d", actualDaysAgo),
		Expected: expected.String(),
		Matches:  matches,
	}
}

func (p *timerPlugin) Fix(input plugin.Result) plugin.Result {
	return input
}

func (p *timerPlugin) getLastCommitTime(dir string) (time.Time, error) {
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return time.Time{}, err
	}
	log, err := repo.Log(&git.LogOptions{Order: git.LogOrderCommitterTime})
	if err != nil {
		return time.Time{}, err
	}
	defer log.Close()
	head, err := log.Next()
	if err != nil {
		return time.Time{}, err
	}
	return head.Committer.When, nil
}

func main() {
	p := timerPlugin{}
	err := plugin.Start(&p)
	if err != nil {
		panic(err)
	}
}
