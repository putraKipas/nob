// PINGPONG APPS
// 2 player => 2 goroutine

// kondisi kalah : saat flag/counter/random number, habis dibagi 11 (random % 11 == 0 THAN break)

// Contoh :
// Player A = Hit 1 // counter 8
// Player B = Hit 2 // counter 3
// Player A = Hit 3 // counter 24
// Player B = Hit 4 // counter 33

// Player B kalah, total hit : 4, kalah di nomor 33

// Player A = Hit 1 // counter 8
// Player B = Hit 2 // counter 9
// Player A = Hit 3 // counter 22

// Player A kalah, total hit : 3, kalah di nomor 22

// Requirement :
// - Struct Player {
// Name string
// Hit int
// }

// a := Player{Name: A, Hit:0}
// b := Player{Name: B, Hit:0}

// a.Hit++
// b.Hit++