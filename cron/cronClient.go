/**
* @Author: xhzhang
* @Date: 2019/9/24 13:39
 */
package cron

import (
	"github.com/robfig/cron/v3"
)

type CronClient struct {
	Client *cron.Cron
}

func NewCronClient() CronClient {
	c := cron.New()
	return CronClient{Client: c}
}

func (c *CronClient) StartCron() {
	c.Client.Start()
}

func (c *CronClient) AddFunction(spec string, execFun func()) (cron.EntryID, error) {
	return c.Client.AddFunc(spec, execFun)
}

func (c *CronClient) AddJob(spec string, job cron.Job) (cron.EntryID, error) {
	return c.Client.AddJob(spec, job)
}

func (c *CronClient) RemoveJob(id cron.EntryID) {
	c.Client.Remove(id)
}

func (c *CronClient) GetEntrys() []cron.Entry {
	return c.Client.Entries()
}


func (c *CronClient) StopCron() {
	c.Client.Stop()
}


