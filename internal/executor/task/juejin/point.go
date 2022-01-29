package juejin

type Point struct{}

func (p *Point) Domain() crontab.Domain {
	return crontab.Domain_JueJin
}

func (p *Point) Kind() crontab.Kind {
	return crontab.Kind_JueJinPoint
}

func (p *Point) Execute(key string) (string, error) {
	var o ore
	return execute(key, pointURL, &o)
}
