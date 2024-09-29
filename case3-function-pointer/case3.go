package main

import "fmt"

type Mahasiswa struct {
	Name    string
	IsAsdos bool
}

func updateStatus(mhs *[]Mahasiswa, name string) {
    for i, m := range *mhs {
        if m.Name == name {
            (*mhs)[i].IsAsdos = true
            return
        }
    }
}

func main() {
    m1 := Mahasiswa{Name: "Andi", IsAsdos: true}
    m2 := Mahasiswa{Name: "Budi", IsAsdos: true}
    m3 := Mahasiswa{Name: "Caca", IsAsdos: false}

    asdos := []Mahasiswa{m1, m2, m3}

    updateStatus(&asdos, "Caca")
    fmt.Println(m3)
}