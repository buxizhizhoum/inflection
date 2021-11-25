# Inflection

Inflection is a string transformation library. It transforms strings from CamelCase to underscored string.
This is an implement of [Inflection](https://github.com/jpvanhal/inflection) package in Python.

# examples

        import (
	        "fmt"
	        "github.com/buxizhizhoum/inflection"
        )

        func example () {
        // to convert a string to underscore
        res := inflection.Underscore("aA")
        // will return a_a
        fmt.Println(res)
        // to convert a string to camelize
        // will return AA
        fmt.Println(inflection.Camelize("a_a", true))
        }