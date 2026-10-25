package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.2" )

func DaLuP6aZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeAqBARvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PuEfnG4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RwSiAc7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcgRzH7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func guRUm48uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func al8R3MjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c14gCh7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nag3bixPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZoZjL0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IySQZJ5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Ma2oFwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmsWXSFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UrzF2dkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHsiW5ufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQugQM6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygu7EuGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7aXCtDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wKBHDXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzjcLzO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWK95CHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZJp2yJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZJXcgddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fksMga6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCoDCyl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wyTu5KJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xfu1seVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdMNCAAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vilrPcaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLVxcXuKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izMD91BPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uuYOXBURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OxVJDfo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eauXXbwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyqtOemJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFMRykhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0l3ClAG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UU74Q4yJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nj0FZddtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZvHs23xSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k026mVdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FG9sshvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MOldpz4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0vJLrboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGUGfyXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BOoTTYNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eu2pdILzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmLoHyL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gcR2hUuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsUYPgSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0s0CB8TZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RgfxDKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZGRXui9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bkw1dEo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kecIRwL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbq6WNfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pe54bWaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9MnzvRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcBaOEBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vfPcjMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qivXadCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgGwrjDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bb21Sp6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TlU2KjHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGlydFZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U6HqHUg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwWCvj79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkwSgWPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fEMkUzcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJIkbX5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0smSvzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPeKvX1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULcG9mUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9F7wls5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqweVEefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KkQTbLQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVajIWELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hmHJ4nUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKsalx8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TknhwsdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Za43KnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiKreSYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRbL1EkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjFM42IEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6cbfAhXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEACdKccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FiHLTpnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7hhEW6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqe7asmTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3iRy8dSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsO0HMVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1AnHduoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gHRNV3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXKlJtkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jxw2uJDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smzEfkXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Ufwv638Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LOcyikO7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fB93gsHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RFaCA3CuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wz6BtQCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQtA4w4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bx7oB8Q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8RaZk6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1UlDnVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qHdd6KLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWfOZmhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7AsenMqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtSU0ZWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqZhSX1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAqTTzzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUAbquMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tcV1sdhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nT1z8xxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RbS6L1P1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzgPuZZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxVskvoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lm3gbHtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twE90xKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6be0kD2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ThsSB2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1LPqHQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVDTrn8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHDaSIvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DwzQeSV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVLpdP1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJS6DlxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8oxz6LtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwaA907iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AeCxmIh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SsQ5z2QmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIzkvvRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3i0Yq9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3liGnO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q10gM7CeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48TWQyDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMbh5cC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYkxR8YHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P68QWnZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOrm6u6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxrSRIRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6Ls4wi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nA6zkuIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sx2PaLhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARAFMcQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLgSmzyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5qjDu653Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLl9s3ZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDgaNC2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wFEnR6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PO5vjgD6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psOS7VzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcT2QqOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoSc89zsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLRJrIO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAjGkV7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lM3jB8xJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABSS1iaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWujuHZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzeJlZAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nFwMlaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9qjhBAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RkCIu7tuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1BjVmH0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3oRXKQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6BLu4AVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djtE7EA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIXKOimFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltHrZKl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hEvLd3PsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDj3jCArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZVEhDOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xq6arVADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Gq06spyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIL6mfMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UR4DrOxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gPHZxI5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFSaLrGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5FJJr6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UiNcamInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QF7pJOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0AJ5KOHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSypgBJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XyJQP4iZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mq5wMVSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pCtO5DtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Wv8lx1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjeZC8CXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrHUbFG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0rs2Oa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vuxIy0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EoXJtvhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93r0ny6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZKXM5ZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgLiaUESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBMBEn7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiQ8vEmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ev8V1wPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcsPGLGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bvSPRxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hL5bM7xiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38Qdsl0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZS6F5VSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIVqo7VaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0EJrGNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AnhoRRtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6eDZn12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zPu2uf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKHWRSzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kg8jFxHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BlcIrLtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DziVi9MmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsTnAhp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLwOdeKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sB8cpF8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vpLmkSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rVSeY6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UffKs8TAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JgCMJo2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmHoUObmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYw5yoqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fnlvq835Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXDHEecbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdoFxjjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXCwWAIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvFr4AT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kpxW4GcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

