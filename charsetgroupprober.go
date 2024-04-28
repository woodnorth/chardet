package chardet

type CharSetGroupProber struct {
	CommonCharsetProber
	activeNum       int
	probers         []CharsetProber
	bestGuessProber CharsetProber
}

func (c *CharSetGroupProber) init(filter LanguageFilter) {
	c.CommonCharsetProber.init(filter)
	c.activeNum = 0
	c.probers = make([]CharsetProber, 0)
	c.bestGuessProber = nil
}

func (c *CharSetGroupProber) reset() {
	c.CommonCharsetProber.reset()
	c.activeNum = 0
	for _, prober := range c.probers {
		prober.reset()
		prober.enActive()
		c.activeNum += 1
	}
	c.bestGuessProber = nil
}

func (c *CharSetGroupProber) charName() string {
	if c.bestGuessProber == nil {
		c.getConfidence()
		if c.bestGuessProber == nil {
			return ""
		}
	}
	return c.bestGuessProber.charName()
}

func (c *CharSetGroupProber) language() string {
	if c.bestGuessProber == nil {
		c.getConfidence()
		if c.bestGuessProber == nil {
			return ""
		}
	}
	return c.bestGuessProber.language()
}

func (c *CharSetGroupProber) feed(data []byte) ProbingState {
	for _, prober := range c.probers {
		if !prober.isActive() {
			continue
		}
		state := prober.feed(data)
		if state == FOUND_IT {
			c.bestGuessProber = prober
			c.state_ = FOUND_IT
			return c.state_
		}

		if state == NOT_ME {
			prober.disActive()
			c.activeNum -= 1
			if c.activeNum <= 0 {
				c.state_ = NOT_ME
				return c.state_
			}
		}
	}
	return c.state_
}

func (c *CharSetGroupProber) getConfidence() float64 {
	state := c.state_
	if state == FOUND_IT {
		return 0.99
	}
	if state == NOT_ME {
		return 0.01
	}

	bestConf := 0.0
	c.bestGuessProber = nil
	for _, prober := range c.probers {
		if !prober.isActive() {
			continue
		}
		conf := prober.getConfidence()
		if bestConf < conf {
			bestConf = conf
			c.bestGuessProber = prober
		}
	}

	if c.bestGuessProber == nil {
		return 0.0
	}
	return bestConf
}

func (c *CharSetGroupProber) getProbers() []CharsetProber {
	return c.probers
}
