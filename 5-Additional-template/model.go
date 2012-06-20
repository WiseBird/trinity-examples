package main

type Model struct {
	Name string
	Address *AddressModel
}

type AddressModel struct {
	Street string
	City string
	State string
	Country string
}