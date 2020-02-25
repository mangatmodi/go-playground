package benchmarks

import (
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"testing"
	"time"
)

/**
BenchmarkMapSetAndGet10-12         	20000000	        62.3 ns/op
BenchmarkMapSetAndGet100-12        	20000000	        64.8 ns/op
BenchmarkMapSetAndGet1000-12       	20000000	        69.6 ns/op
BenchmarkMapSetAndGet10000-12      	20000000	        90.6 ns/op
BenchmarkMapSetAndGet100000-12     	 5000000	       242 ns/op
BenchmarkMapSetAndGet1000000-12    	       1	1486821755 ns/op

Only Get
BenchmarkMapSetAndGet10-12         	100000000	        12.7 ns/op
BenchmarkMapSetAndGet100-12        	100000000	        11.5 ns/op
BenchmarkMapSetAndGet1000-12       	100000000	        11.3 ns/op
BenchmarkMapSetAndGet10000-12      	1000000000	         2.10 ns/op
BenchmarkMapSetAndGet100000-12     	500000000	         2.73 ns/op
BenchmarkMapSetAndGet1000000-12    	       1	1481274697 ns/op

*/
func randomStrings(size int) []string {
	m := make([]string, 0, size)
	//Leave one
	for i := 0; i < size; i++ {
		m = append(m, faker.UUIDDigit())
	}
	return m
}
func getRandomMap(size int) map[string]string {
	m := make(map[string]string, size)
	//Leave one
	for i := 0; i < size/2; i++ {
		k := faker.UUIDDigit()
		v := faker.UUIDDigit()
		m[k] = v
	}
	return m
}

func BenchmarkMapSetAndGet10(b *testing.B) {
	var map10 = getRandomMap(10)
	r := randomStrings(10)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		k := r[rand.Intn(len(r))]
		v := r[rand.Intn(len(r))]
		map10[k] = v
		v1 := map10[k]
		if v != v1 {
			panic("not same")
		}
	}
}
func BenchmarkMapSetAndGet100(b *testing.B) {
	var map10 = getRandomMap(100)
	r := randomStrings(100)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		k := r[rand.Intn(len(r))]
		v := r[rand.Intn(len(r))]
		map10[k] = v
		v1 := map10[k]
		if v != v1 {
			panic("not same")
		}
	}
}
func BenchmarkMapSetAndGet1000(b *testing.B) {
	var map10 = getRandomMap(1000)
	r := randomStrings(1000)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		k := r[rand.Intn(len(r))]
		v := r[rand.Intn(len(r))]
		map10[k] = v
		v1 := map10[k]
		if v != v1 {
			panic("not same")
		}
	}
}
func BenchmarkMapSetAndGet10000(b *testing.B) {
	var map10 = getRandomMap(10000)
	r := randomStrings(10000)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		k := r[rand.Intn(len(r))]
		v := r[rand.Intn(len(r))]
		map10[k] = v
		v1 := map10[k]
		if v != v1 {
			panic("not same")
		}
	}
}
func BenchmarkMapSetAndGet100000(b *testing.B) {
	var map10 = getRandomMap(100000)
	r := randomStrings(100000)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		k := r[rand.Intn(len(r))]
		v := r[rand.Intn(len(r))]
		map10[k] = v
		v1 := map10[k]
		if v != v1 {
			panic("not same")
		}
	}
}
func BenchmarkMapSetAndGet1000000(b *testing.B) {
	var map10 = getRandomMap(1000000)
	r := randomStrings(1000000)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		k := r[rand.Intn(len(r))]
		v := r[rand.Intn(len(r))]
		map10[k] = v
		v1 := map10[k]
		if v != v1 {
			panic("not same")
		}
	}
}
