package main

// Hamming weight pulled from wikipedia
// http://en.wikipedia.org/wiki/Hamming_weight
// this includes the crazy uint64s and popcount
const (
	m1  = 0x5555555555555555 //binary: 0101...
	m2  = 0x3333333333333333 //binary: 00110011..
	m4  = 0x0f0f0f0f0f0f0f0f //binary:  4 zeros,  4 ones ...
	m8  = 0x00ff00ff00ff00ff //binary:  8 zeros,  8 ones ...
	m16 = 0x0000ffff0000ffff //binary: 16 zeros, 16 ones ...
	m32 = 0x00000000ffffffff //binary: 32 zeros, 32 ones
	hff = 0xffffffffffffffff //binary: all ones
	h01 = 0x0101010101010101 //the sum of 256 to the power of 0,1,2,3...
)

//This uses fewer arithmetic operations than any other known
//implementation on machines with slow multiplication.
//It uses 17 arithmetic operations.
func popcount_2(x uint64) uint {
	x -= (x >> 1) & m1             //put count of each 2 bits into those 2 bits
	x = (x & m2) + ((x >> 2) & m2) //put count of each 4 bits into those 4 bits
	x = (x + (x >> 4)) & m4        //put count of each 8 bits into those 8 bits
	x += x >> 8                    //put count of each 16 bits into their lowest 8 bits
	x += x >> 16                   //put count of each 32 bits into their lowest 8 bits
	x += x >> 32                   //put count of each 64 bits into their lowest 8 bits
	return uint(x & 0x7f)
}

//This uses fewer arithmetic operations than any other known
//implementation on machines with fast multiplication.
//It uses 12 arithmetic operations, one of which is a multiply.
func popcount_3(x uint64) uint {
	x -= (x >> 1) & m1             //put count of each 2 bits into those 2 bits
	x = (x & m2) + ((x >> 2) & m2) //put count of each 4 bits into those 4 bits
	x = (x + (x >> 4)) & m4        //put count of each 8 bits into those 8 bits
	return uint((x * h01) >> 56)   //returns left 8 bits of x + (x<<8) + (x<<16) + (x<<24) + ...
}
