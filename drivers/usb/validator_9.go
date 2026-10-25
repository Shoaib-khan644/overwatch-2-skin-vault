package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.4" )

func OAbi1VhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IynrmiFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTLNHAnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPEInqw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kz8HEkcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktNNgn8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMxUPKFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AFUoLZN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bH84ruEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdn22RL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9b0ClscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iu5rygyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gph71eNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvbrfM8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CeddP3IOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKV4OerNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSxxHBmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEH3hLcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCcYpjCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOVH6QJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKyDCgBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NRghNZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAxrzi1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZ0FEO5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2thpBpMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iyzEn3G5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8QfROmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mo9Sy6tTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqSdRAz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcK3nifyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyDBY4cfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXHbHBKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRjRL3VGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FugjM2qzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoLu5T4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3t8Im66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63mWhuKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0UlpCWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 849WjXcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txui0ThIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gL0hnnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9Ukza39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AGxdIUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24TIWZqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrL957RAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 44YklL31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4G5THdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LkgAlBLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deLdE680Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wiXpaSaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hTPFjlpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AwDwembWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZ757TkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBH8w8GhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17V1HG93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80A0TlyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZ1ls7DPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2C5TZf9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOhivnjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFfQ7ti3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RkKGmJTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zu0Qn75aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNENR1IHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3RlFNSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uK1nQ1c4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFqYjRgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5bXzGuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOaySu2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 60DcEXa7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJuZSFIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7jTL9kTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1U4RhjjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODs2jHGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TmfsRMWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPEnqsXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqW1ZtGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N84rQs4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ku9Rnii9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obd83PA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSHekFE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgLbGIrlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTjozfETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o6FTBrzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIK0RPk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZATyCwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X61XSYr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6gWLqMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTEtUyXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PtEzX6eNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5V2CVnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p6vou1b6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A9btDj6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6G8T10HJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIIXaM0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7sCJvZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92SE8OUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1VxOHKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJGrwHPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2r9XAWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qivFXlCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QbxDD2WcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aV4aY8RHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4TERxsjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YeM0AAGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NaWu8nHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PB8jczobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJ0qbhuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMMiAAm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NVieKqMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsUapMZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFPicLT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbICH0q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7y3dDDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHXXxbXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dB9yJchCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CqhQN3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dFYfmW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlKEdxORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6AmV9kPJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpHzwA2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2NJTF1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOT2VRMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2vrltXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j2UrJ6kcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0lupIqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dM0Und9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icLhCyJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMIEjJNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVvBPBJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TbgslADKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVnioqCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1dkddraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d03HuVNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HI5lY4VwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAteqtUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5qLiYIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ArGWdA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgxf11VxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jB3wD1vYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHSitn6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lF8S4mHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uvU7vWgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0edrtEHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99ZvbDcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7vMySnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KPnhJaqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHqV3R34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYWNoUBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQjUhUR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4kyEcjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upbyffPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBqQz4CxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUCNcdyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gd2kqJEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbBxoWr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VT7BxY4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZyeDEZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsLzbfpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkToRA55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFLb9xTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gF7ZvS3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdBQbKnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iTRDoguvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCKTomeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQeK7HuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SULvKPrvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9LwtZlm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rnMrJ17wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 841dSB8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJgurAMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9sACPPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGCeFgwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A0uiKhF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQRgxPpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MKHfTqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QniWLDJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8y7axB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWFtE9QfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6l3de7NnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5zyXMDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZD0tcIN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B50k9ElLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FcZqFlXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPNhBBTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func loM6W25hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJD6pANnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yW2r783QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8T7pxDiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LsLxdLw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YYgK8m2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gNTKEnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func liG6xwoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6j6vt6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func no6g17WMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twtbctJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ir5iVLgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJQc0Xs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8Jj7OMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YT5itOKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TeHQkwUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfyBuPUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQj5vQNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LE6lUP8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiAafXr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MN0xuBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2h6TmEYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func To0kAyYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEVMXCt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3LwOCeNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIZKMC2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BE0oPiErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gyCQII0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uk5DwhKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMj4ztePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wozLz1gKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNJergvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func moEbLgh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15GLcFYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxZW78H7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iw6rpux8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HcblNmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6e9fRcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lvjl9iaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezNxOv3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKT4mJl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Vr0XAqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frlCIXjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwIUgbSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQ5a9PzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jz5hetosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UvEaLRMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZUOaRTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNx9ISIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TFiBW9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCNofHkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUNM1BQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AaYrlJLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FxCn8N6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WnZYetXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QctOIoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwM0rrBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGn80FvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZLqXtTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWV4K7eMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QV1RMWpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdgCZxxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1eR6rCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5fR8ZrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uB2GcFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFehEqPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vl0YB9gRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZtOwkmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJVGhjIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amu7yxuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YotO4ltTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81URsmN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWLddyRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLMCIj5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B3zg3cB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqfImtgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4o4JrUOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3Cbj4RiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9vh5DfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G05TvEoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Se86NZYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pa7AxMVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2Du4G0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QclH9B3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Myjqeq1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f732PYhBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3D28MaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y18NWT13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdLVCfCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wncL15F6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNLFDXR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7TIKgv7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EEhnCEndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EsKkQRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HM3MKfy8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDhiKi0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aek0Al6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g85q3lu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfjBEuUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Hwkt8YhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clsscEOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMSAz3h9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2Xj2Gh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Qg6R3kJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUC05y8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNVivbHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KefLTl0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGTrO27TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upTLcHNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLyFsfkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4SgmQPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5IZ1kqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nFYbAyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NGSnsTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQffnFaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4h47KI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZrYX3u1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKqHkE4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cg4oSkEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHfHkRhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AwGGI1xfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3lPqm7OgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9x1wuZW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AJgmFtgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1sHUCt7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOStrlFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uc2MR1NDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxfCewusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjbruW21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dd5Odo8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgtLykFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thyTig1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6I3OzkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWSNczhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EP3JnNtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKora68BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQY8SCteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Ysby9YMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPGpngPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlMFPtt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqQFHNJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axfaFyziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhQYTJSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3b8gY1MVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qqwSfMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1eDahXyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7S337rIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDSmXJxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IK2LV7RUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7AVVRofGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qr0257aaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNHgmj1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzomM6fNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EPjd3C5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43vLhkm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHQxBnGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErmpPs2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xS39WZeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvSojx2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5X433gZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1lKTENVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8EbCkiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IeWhni5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mNRvhKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojr8jHQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VL3jIFyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjXIqc2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JxeBodrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFnQrfEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKxExAPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iKpqMKgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8fRN8HdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUhh6cNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EqzPjzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MC2IRoA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRxYr63PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spTqGV2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygOAWydpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JnaCZYUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxKDi3FaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IW1FQPFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQ8wKJCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pD5BZ5D2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5DSTKehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aeUE7ob5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5wgFBbnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMAFS1f2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djuuvJ2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WkkNNCzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIv9ds0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycujGuU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AhKdN5aCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JpgXKRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5npo9897Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cn6Mac9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHj01jMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5KUr1S3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9eQpAR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6r3H483FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCfj4BtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cysvbqccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PO5CrJ61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzHnHrqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Mj3D3LIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJkC3MCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func olyUpvIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func slODnILiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcOrRlSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMxpyBpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 09PMrDjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pPRNiuCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjHrbVb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xhJAXEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5u4vvBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zAN2gN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mWbAVtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z6R20hF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wxy5rHH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQjannx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mAHQdUsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyrE1uofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmnuqXDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRgR3n4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZU4oUWyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5j4i4LfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QRDQrLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BecskRdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhIhQ5kVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLwuHN32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XbCqefhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ec7zTuZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5TkFuhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

