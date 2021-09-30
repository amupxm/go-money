package money

//Abs returns the absolute value of the amount.
func (c CAD) Abs() CAD {
	if c.dollar < 0 || c.cents < 0 {
		return Cents(-c.AsCent())
	}
	return c
}

//Add adds the given amount to the current amount.
func (c CAD) Add(o CAD) CAD {
	sum := Cents(c.AsCent() + o.AsCent())
	return sum
}

//Mul multiplies the amount by the given factor.
func (c CAD) Mul(scalar int64) CAD {
	return Cents(c.AsCent() * scalar)
}

//CanonicalForm returns the canonical form of the amount.
func (c CAD) CanonicalForm() (dollar int64, cent int64) {
	return c.dollar, c.cents
}

//Sub subtracts the given amount from the current amount.
func (c CAD) Sub(o CAD) CAD {
	return Cents(c.AsCent() - o.AsCent())
}
