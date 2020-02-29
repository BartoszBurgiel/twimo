package base

// Base2 is binary
var Base2 Base = []byte("01")

// Base8 is octal
var Base8 Base = []byte("01234567")

// Base16 is hexadecimal
var Base16 Base = []byte("01234567890abcdef")

// Base53 is a variant of Base62 excluding I, O, l, U, and u or I, 1, l, 0, and O
var Base53 Base = []byte("abcdefghjkmnpqrstvwxzyABCDEFGHJKLNMPQRSTVWXYZ23456789")

// Base60 = [a-z][A-Z][0-9]
var Base60 Base = []byte("abcdefghijklmnopqrstuvwxzyABCDEFGHIJKLNMOPQRSTUVWXYZ0123456789")

// Base85 = [a-z][A+Z][0-9] && all special characters exept {-, `, ", ', ´}
var Base85 Base = []byte("abcdefghijklmnoprstuvwxzyABCDEFGHIJKLNMOPRSTUVWXYZ0123456789!#$%^&*()_+~=[{}];:,./<>?")
