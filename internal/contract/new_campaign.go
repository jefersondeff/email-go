package contract

type NewCampaign struct {
	Name    string   `name:"required"`
	Content string   `content:"required"`
	Emails  []string `emails:"required"`
}
