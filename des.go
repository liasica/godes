package godes

func Des(key string, message string, encrypt int, mode int, iv string, padding int) string {

    var spfunction1 []int = []int{0x1010400, 0, 0x10000, 0x1010404, 0x1010004, 0x10404, 0x4, 0x10000, 0x400, 0x1010400, 0x1010404, 0x400, 0x1000404, 0x1010004, 0x1000000, 0x4, 0x404, 0x1000400, 0x1000400, 0x10400, 0x10400, 0x1010000, 0x1010000, 0x1000404, 0x10004, 0x1000004, 0x1000004, 0x10004, 0, 0x404, 0x10404, 0x1000000, 0x10000, 0x1010404, 0x4, 0x1010000, 0x1010400, 0x1000000, 0x1000000, 0x400, 0x1010004, 0x10000, 0x10400, 0x1000004, 0x400, 0x4, 0x1000404, 0x10404, 0x1010404, 0x10004, 0x1010000, 0x1000404, 0x1000004, 0x404, 0x10404, 0x1010400, 0x404, 0x1000400, 0x1000400, 0, 0x10004, 0x10400, 0, 0x1010004}
    var spfunction2 []int = []int{-0x7fef7fe0, -0x7fff8000, 0x8000, 0x108020, 0x100000, 0x20, -0x7fefffe0, -0x7fff7fe0, -0x7fffffe0, -0x7fef7fe0, -0x7fef8000, -0x80000000, -0x7fff8000, 0x100000, 0x20, -0x7fefffe0, 0x108000, 0x100020, -0x7fff7fe0, 0, -0x80000000, 0x8000, 0x108020, -0x7ff00000, 0x100020, -0x7fffffe0, 0, 0x108000, 0x8020, -0x7fef8000, -0x7ff00000, 0x8020, 0, 0x108020, -0x7fefffe0, 0x100000, -0x7fff7fe0, -0x7ff00000, -0x7fef8000, 0x8000, -0x7ff00000, -0x7fff8000, 0x20, -0x7fef7fe0, 0x108020, 0x20, 0x8000, -0x80000000, 0x8020, -0x7fef8000, 0x100000, -0x7fffffe0, 0x100020, -0x7fff7fe0, -0x7fffffe0, 0x100020, 0x108000, 0, -0x7fff8000, 0x8020, -0x80000000, -0x7fefffe0, -0x7fef7fe0, 0x108000}
    var spfunction3 []int = []int{0x208, 0x8020200, 0, 0x8020008, 0x8000200, 0, 0x20208, 0x8000200, 0x20008, 0x8000008, 0x8000008, 0x20000, 0x8020208, 0x20008, 0x8020000, 0x208, 0x8000000, 0x8, 0x8020200, 0x200, 0x20200, 0x8020000, 0x8020008, 0x20208, 0x8000208, 0x20200, 0x20000, 0x8000208, 0x8, 0x8020208, 0x200, 0x8000000, 0x8020200, 0x8000000, 0x20008, 0x208, 0x20000, 0x8020200, 0x8000200, 0, 0x200, 0x20008, 0x8020208, 0x8000200, 0x8000008, 0x200, 0, 0x8020008, 0x8000208, 0x20000, 0x8000000, 0x8020208, 0x8, 0x20208, 0x20200, 0x8000008, 0x8020000, 0x8000208, 0x208, 0x8020000, 0x20208, 0x8, 0x8020008, 0x20200}
    var spfunction4 []int = []int{0x802001, 0x2081, 0x2081, 0x80, 0x802080, 0x800081, 0x800001, 0x2001, 0, 0x802000, 0x802000, 0x802081, 0x81, 0, 0x800080, 0x800001, 0x1, 0x2000, 0x800000, 0x802001, 0x80, 0x800000, 0x2001, 0x2080, 0x800081, 0x1, 0x2080, 0x800080, 0x2000, 0x802080, 0x802081, 0x81, 0x800080, 0x800001, 0x802000, 0x802081, 0x81, 0, 0, 0x802000, 0x2080, 0x800080, 0x800081, 0x1, 0x802001, 0x2081, 0x2081, 0x80, 0x802081, 0x81, 0x1, 0x2000, 0x800001, 0x2001, 0x802080, 0x800081, 0x2001, 0x2080, 0x800000, 0x802001, 0x80, 0x800000, 0x2000, 0x802080}
    var spfunction5 []int = []int{0x100, 0x2080100, 0x2080000, 0x42000100, 0x80000, 0x100, 0x40000000, 0x2080000, 0x40080100, 0x80000, 0x2000100, 0x40080100, 0x42000100, 0x42080000, 0x80100, 0x40000000, 0x2000000, 0x40080000, 0x40080000, 0, 0x40000100, 0x42080100, 0x42080100, 0x2000100, 0x42080000, 0x40000100, 0, 0x42000000, 0x2080100, 0x2000000, 0x42000000, 0x80100, 0x80000, 0x42000100, 0x100, 0x2000000, 0x40000000, 0x2080000, 0x42000100, 0x40080100, 0x2000100, 0x40000000, 0x42080000, 0x2080100, 0x40080100, 0x100, 0x2000000, 0x42080000, 0x42080100, 0x80100, 0x42000000, 0x42080100, 0x2080000, 0, 0x40080000, 0x42000000, 0x80100, 0x2000100, 0x40000100, 0x80000, 0, 0x40080000, 0x2080100, 0x40000100}
    var spfunction6 []int = []int{0x20000010, 0x20400000, 0x4000, 0x20404010, 0x20400000, 0x10, 0x20404010, 0x400000, 0x20004000, 0x404010, 0x400000, 0x20000010, 0x400010, 0x20004000, 0x20000000, 0x4010, 0, 0x400010, 0x20004010, 0x4000, 0x404000, 0x20004010, 0x10, 0x20400010, 0x20400010, 0, 0x404010, 0x20404000, 0x4010, 0x404000, 0x20404000, 0x20000000, 0x20004000, 0x10, 0x20400010, 0x404000, 0x20404010, 0x400000, 0x4010, 0x20000010, 0x400000, 0x20004000, 0x20000000, 0x4010, 0x20000010, 0x20404010, 0x404000, 0x20400000, 0x404010, 0x20404000, 0, 0x20400010, 0x10, 0x4000, 0x20400000, 0x404010, 0x4000, 0x400010, 0x20004010, 0, 0x20404000, 0x20000000, 0x400010, 0x20004010}
    var spfunction7 []int = []int{0x200000, 0x4200002, 0x4000802, 0, 0x800, 0x4000802, 0x200802, 0x4200800, 0x4200802, 0x200000, 0, 0x4000002, 0x2, 0x4000000, 0x4200002, 0x802, 0x4000800, 0x200802, 0x200002, 0x4000800, 0x4000002, 0x4200000, 0x4200800, 0x200002, 0x4200000, 0x800, 0x802, 0x4200802, 0x200800, 0x2, 0x4000000, 0x200800, 0x4000000, 0x200800, 0x200000, 0x4000802, 0x4000802, 0x4200002, 0x4200002, 0x2, 0x200002, 0x4000000, 0x4000800, 0x200000, 0x4200800, 0x802, 0x200802, 0x4200800, 0x802, 0x4000002, 0x4200802, 0x4200000, 0x200800, 0, 0x2, 0x4200802, 0, 0x200802, 0x4200000, 0x800, 0x4000002, 0x4000800, 0x800, 0x200002}
    var spfunction8 []int = []int{0x10001040, 0x1000, 0x40000, 0x10041040, 0x10000000, 0x10001040, 0x40, 0x10000000, 0x40040, 0x10040000, 0x10041040, 0x41000, 0x10041000, 0x41040, 0x1000, 0x40, 0x10040000, 0x10000040, 0x10001000, 0x1040, 0x41000, 0x40040, 0x10040040, 0x10041000, 0x1040, 0, 0, 0x10040040, 0x10000040, 0x10001000, 0x41040, 0x40000, 0x41040, 0x40000, 0x10041000, 0x1000, 0x40, 0x10040040, 0x1000, 0x41040, 0x10001000, 0x40, 0x10000040, 0x10040000, 0x10040040, 0x10000000, 0x40000, 0x10001040, 0, 0x10041040, 0x40040, 0x10000040, 0x10040000, 0x10001000, 0x10001040, 0, 0x10041040, 0x41000, 0x41000, 0x1040, 0x1040, 0x40040, 0x10000000, 0x10041000}
    var masks []int = []int{4294967295, 2147483647, 1073741823, 536870911, 268435455, 134217727, 67108863, 33554431, 16777215, 8388607, 4194303, 2097151, 1048575, 524287, 262143, 131071, 65535, 32767, 16383, 8191, 4095, 2047, 1023, 511, 255, 127, 63, 31, 15, 7, 3, 1, 0}

    keys := Des_createKeys(key)
    m := 0
    messageLen := len(message)
    var temp int
    chunk := 0
    //set up the loops for single and triple des
    iterations := 9
    looping := []int{94, 62, -2, 32, 64, 2, 30, -2, -2}
    if encrypt != 0 {
        looping = []int{0, 32, 2, 62, 30, -2, 64, 96, 2}
    }
    if len(keys) == 32 {
        iterations = 3
        looping = []int{30, -2, -2}
        if encrypt != 0 {
            looping = []int{0, 32, 2}
        }
    }
    if padding == 2 {
        message += "        "
    } else {
        if padding == 1 {
            temp = 8 - (messageLen % 8)
            message += string(temp) + string(temp) + string(temp) + string(temp) + string(temp) + string(temp) + string(temp) + string(temp)
            if temp == 8 {
                messageLen += 8
            }
        } else if padding == 0 {
            message += string(0) + string(0) + string(0) + string(0) + string(0) + string(0) + string(0) + string(0)
        }
    }
    var result, tempresult string
    var cbcleft, cbcright int
    var cbcleft2, cbcright2 int
    if mode == 1 {
        cbcleft = (int(iv[0]) << 24) | (int(iv[1]) << 16) | (int(iv[2]) << 8) | int(iv[3])
        cbcright = (int(iv[4]) << 24) | (int(iv[5]) << 16) | (int(iv[6]) << 8) | int(iv[7])
        m = 0
    }
    var left, right, m_loop int

    for m < messageLen {
        left = (int(message[(0 + m_loop*8)]) << 24) | (int(message[1+m_loop*8]) << 16) | (int(message[2+m_loop*8]) << 8) | int(message[3+m_loop*8])
        right = (int(message[4+m_loop*8]) << 24) | (int(message[5+m_loop*8]) << 16) | (int(message[6+m_loop*8]) << 8) | int(message[7+m_loop*8])
        m = (m_loop + 1) * 8
        m_loop++
        if mode == 1 {
            if encrypt == 1 {
                left ^= cbcleft
                right ^= cbcright
            } else {
                cbcleft2 = cbcleft
                cbcright2 = cbcright
                cbcleft = left
                cbcright = right
            }
        }
        temp = ((left >> 4 & masks[4]) ^ right) & 0x0f0f0f0f
        right ^= temp
        left ^= temp << 4
        temp = ((left >> 16 & masks[16]) ^ right) & 0x0000ffff
        right ^= temp
        left ^= temp << 16
        temp = ((right >> 2 & masks[2]) ^ left) & 0x33333333
        left ^= temp
        right ^= temp << 2
        temp = ((right >> 8 & masks[8]) ^ left) & 0x00ff00ff
        left ^= temp
        right ^= temp << 8
        temp = ((left >> 1 & masks[1]) ^ right) & 0x55555555
        right ^= temp
        left ^= temp << 1
        left = (left << 1) | (left >> 31 & masks[31])
        right = (right << 1) | (right >> 31 & masks[31])
        for j := 0; j < iterations; j += 3 {
            endloop := looping[j+1]
            loopinc := looping[j+2]
            for i := looping[j]; i != endloop; i += loopinc {
                right1 := right ^ keys[i]
                right2 := ((right >> 4 & masks[4]) | (right << 28 & 0xffffffff)) ^ keys[i+1]
                //the result is attained by passing these bytes through the S selection functions
                temp = left
                left = right
                right = temp ^ (spfunction2[(right1>>24&masks[24])&0x3f] | spfunction4[(right1>>16&masks[16])&0x3f] | spfunction6[(right1>>8&masks[8])&0x3f] | spfunction8[right1&0x3f] | spfunction1[(right2>>24&masks[24])&0x3f] | spfunction3[(right2>>16&masks[16])&0x3f] | spfunction5[(right2>>8&masks[8])&0x3f] | spfunction7[right2&0x3f])
            }
            temp = left
            left = right
            right = temp
        }
        left = (left >> 1 & masks[1]) | (left << 31)
        right = (right >> 1 & masks[1]) | (right << 31)
        //now perform IP-1, which is IP in the opposite direction
        temp = ((left >> 1 & masks[1]) ^ right) & 0x55555555
        right ^= temp
        left ^= temp << 1
        temp = ((right >> 8 & masks[8]) ^ left) & 0x00ff00ff
        left ^= temp
        right ^= temp << 8
        temp = ((right >> 2 & masks[2]) ^ left) & 0x33333333
        left ^= temp
        right ^= temp << 2
        temp = ((left >> 16 & masks[16]) ^ right) & 0x0000ffff
        right ^= temp
        left ^= temp << 16
        temp = ((left >> 4 & masks[4]) ^ right) & 0x0f0f0f0f
        right ^= temp
        left ^= temp << 4
        if mode == 1 {
            if encrypt == 1 {
                cbcleft = left
                cbcright = right
            } else {
                left ^= cbcleft2
                right ^= cbcright2
            }
        }
        tempbyte := []byte{byte(left >> 24 & masks[24]), byte((left >> 16 & masks[16]) & 0xff), byte((left >> 8 & masks[8]) & 0xff), byte(left & 0xff), byte(right >> 24 & masks[24]), byte((right >> 16 & masks[16]) & 0xff), byte((right >> 8 & masks[8]) & 0xff), byte(right & 0xff)}
        tempresult += string(tempbyte)
        chunk += 8
        if chunk == 512 {
            result += tempresult
            tempresult = ""
            chunk = 0
        }
    }
    return result + tempresult
}

func Des_createKeys(key string) []int {
    var pc2bytes0 []int = []int{0, 0x4, 0x20000000, 0x20000004, 0x10000, 0x10004, 0x20010000, 0x20010004, 0x200, 0x204, 0x20000200, 0x20000204, 0x10200, 0x10204, 0x20010200, 0x20010204}
    var pc2bytes1 []int = []int{0, 0x1, 0x100000, 0x100001, 0x4000000, 0x4000001, 0x4100000, 0x4100001, 0x100, 0x101, 0x100100, 0x100101, 0x4000100, 0x4000101, 0x4100100, 0x4100101}
    var pc2bytes2 []int = []int{0, 0x8, 0x800, 0x808, 0x1000000, 0x1000008, 0x1000800, 0x1000808, 0, 0x8, 0x800, 0x808, 0x1000000, 0x1000008, 0x1000800, 0x1000808}
    var pc2bytes3 []int = []int{0, 0x200000, 0x8000000, 0x8200000, 0x2000, 0x202000, 0x8002000, 0x8202000, 0x20000, 0x220000, 0x8020000, 0x8220000, 0x22000, 0x222000, 0x8022000, 0x8222000}
    var pc2bytes4 []int = []int{0, 0x40000, 0x10, 0x40010, 0, 0x40000, 0x10, 0x40010, 0x1000, 0x41000, 0x1010, 0x41010, 0x1000, 0x41000, 0x1010, 0x41010}
    var pc2bytes5 []int = []int{0, 0x400, 0x20, 0x420, 0, 0x400, 0x20, 0x420, 0x2000000, 0x2000400, 0x2000020, 0x2000420, 0x2000000, 0x2000400, 0x2000020, 0x2000420}
    var pc2bytes6 []int = []int{0, 0x10000000, 0x80000, 0x10080000, 0x2, 0x10000002, 0x80002, 0x10080002, 0, 0x10000000, 0x80000, 0x10080000, 0x2, 0x10000002, 0x80002, 0x10080002}
    var pc2bytes7 []int = []int{0, 0x10000, 0x800, 0x10800, 0x20000000, 0x20010000, 0x20000800, 0x20010800, 0x20000, 0x30000, 0x20800, 0x30800, 0x20020000, 0x20030000, 0x20020800, 0x20030800}
    var pc2bytes8 []int = []int{0, 0x40000, 0, 0x40000, 0x2, 0x40002, 0x2, 0x40002, 0x2000000, 0x2040000, 0x2000000, 0x2040000, 0x2000002, 0x2040002, 0x2000002, 0x2040002}
    var pc2bytes9 []int = []int{0, 0x10000000, 0x8, 0x10000008, 0, 0x10000000, 0x8, 0x10000008, 0x400, 0x10000400, 0x408, 0x10000408, 0x400, 0x10000400, 0x408, 0x10000408}
    var pc2bytes10 []int = []int{0, 0x20, 0, 0x20, 0x100000, 0x100020, 0x100000, 0x100020, 0x2000, 0x2020, 0x2000, 0x2020, 0x102000, 0x102020, 0x102000, 0x102020}
    var pc2bytes11 []int = []int{0, 0x1000000, 0x200, 0x1000200, 0x200000, 0x1200000, 0x200200, 0x1200200, 0x4000000, 0x5000000, 0x4000200, 0x5000200, 0x4200000, 0x5200000, 0x4200200, 0x5200200}
    var pc2bytes12 []int = []int{0, 0x1000, 0x8000000, 0x8001000, 0x80000, 0x81000, 0x8080000, 0x8081000, 0x10, 0x1010, 0x8000010, 0x8001010, 0x80010, 0x81010, 0x8080010, 0x8081010}
    var pc2bytes13 []int = []int{0, 0x4, 0x100, 0x104, 0, 0x4, 0x100, 0x104, 0x1, 0x5, 0x101, 0x105, 0x1, 0x5, 0x101, 0x105}
    var masks []int = []int{4294967295, 2147483647, 1073741823, 536870911, 268435455, 134217727, 67108863, 33554431, 16777215, 8388607, 4194303, 2097151, 1048575, 524287, 262143, 131071, 65535, 32767, 16383, 8191, 4095, 2047, 1023, 511, 255, 127, 63, 31, 15, 7, 3, 1, 0}

    iterations := 1
    if len(key) > 8 {
        iterations = 3
    }
    keys := make([]int, 32*iterations)
    shifts := []int{0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0}
    n := 0
    for j := 0; j < iterations; j++ {
        left := (int(key[(0 + j*8)]) << 24) | (int(key[(1 + j*8)]) << 16) | (int(key[(2 + j*8)]) << 8) | int(key[(3 + j*8)])
        right := (int(key[(4 + j*8)]) << 24) | (int(key[(5 + j*8)]) << 16) | (int(key[(6 + j*8)]) << 8) | int(key[(7 + j*8)])
        temp := ((left >> 4 & masks[4]) ^ right) & 0x0f0f0f0f
        right ^= temp
        left ^= temp << 4
        temp = ((right >> 16 & masks[16]) ^ left) & 0x0000ffff
        left ^= temp
        right ^= temp << 16
        temp = ((left >> 2 & masks[2]) ^ right) & 0x33333333
        right ^= temp
        left ^= temp << 2
        temp = ((right >> 16 & masks[16]) ^ left) & 0x0000ffff
        left ^= temp
        right ^= temp << 16
        temp = ((left >> 1 & masks[1]) ^ right) & 0x55555555
        right ^= temp
        left ^= temp << 1
        temp = ((right >> 8 & masks[8]) ^ left) & 0x00ff00ff
        left ^= temp
        right ^= temp << 8
        temp = ((left >> 1 & masks[1]) ^ right) & 0x55555555
        right ^= temp
        left ^= temp << 1
        //the right side needs to be shifted and to get the last four bits of the left side
        temp = (left << 8) | ((right >> 20 & masks[20]) & 0x000000f0)
        //left needs to be put upside down
        left = (right << 24) | ((right << 8) & 0xff0000) | ((right >> 8 & masks[8]) & 0xff00) | ((right >> 24 & masks[24]) & 0xf0)
        right = temp
        //now go through and perform these shifts on the left and right keys
        for i := 0; i < len(shifts); i++ {
            if shifts[i] > 0 {
                left = (left << 2) | (left >> 26 & masks[26])
                right = (right << 2) | (right >> 26 & masks[26])
            } else {
                left = (left << 1) | (left >> 27 & masks[27])
                right = (right << 1) | (right >> 27 & masks[27])
            }
            left = left & -0xf
            right = right & -0xf
            lefttemp := pc2bytes0[left>>28&masks[28]] | pc2bytes1[(left>>24&masks[24])&0xf] | pc2bytes2[(left>>20&masks[20])&0xf] | pc2bytes3[(left>>16&masks[16])&0xf] | pc2bytes4[(left>>12&masks[12])&0xf] | pc2bytes5[(left>>8&masks[8])&0xf] | pc2bytes6[(left>>4&masks[4])&0xf]
            righttemp := pc2bytes7[right>>28&masks[28]] | pc2bytes8[(right>>24&masks[24])&0xf] | pc2bytes9[(right>>20&masks[20])&0xf] | pc2bytes10[(right>>16&masks[16])&0xf] | pc2bytes11[(right>>12&masks[12])&0xf] | pc2bytes12[(right>>8&masks[8])&0xf] | pc2bytes13[(right>>4&masks[4])&0xf]
            temp = ((righttemp >> 16 & masks[16]) ^ lefttemp) & 0x0000ffff
            keys[n] = lefttemp ^ temp
            n++
            keys[n] = righttemp ^ (temp << 16)
            n++
        }
    }
    return keys[:n]
}
