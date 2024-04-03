package calculator

import "errors"

type Calculator struct {
	Base float64
}

func InitCalculator(base float64) *Calculator { // inisiasi kalkulator
	return &Calculator{
		Base: base,
	}
}

func (c *Calculator) Add(number float64) { // method untuk penjumlahan
	c.Base += number
}

func (c *Calculator) Subtract(number float64) { // method untuk pengurangan
	c.Base -= number
}

func (c *Calculator) Multiply(number float64) { // method untuk perkalian
	c.Base *= number
}

func (c *Calculator) Divide(number float64) error { // method untuk pembagian
	if number == 0 {
		return errors.New("cannot divide by zero")
	}

	c.Base /= number
	return nil
}

func (c *Calculator) Result() float64 { // method untuk menampilkan hasil
	return c.Base
}

func (c *Calculator) Reset() { // method untuk mereset hasil
	c.Base = 0
}
