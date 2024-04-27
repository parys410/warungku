package warungku

import "fmt"

type Barang struct {
	Nama string
	Harga int
	Qty int
}

type Warung struct {
	StrukBelanja *[]Barang
}

func (warung *Warung) TambahBarang(barang *Barang) {
	*warung.StrukBelanja = append(*warung.StrukBelanja, *barang)
}

func (warung *Warung) KurangiQty(barang *Barang, qty int) {
	for i, item := range *warung.StrukBelanja {
		if item.Nama == barang.Nama {
			(*warung.StrukBelanja)[i].Qty -= 1
		}
	}
}

func (warung *Warung) CetakStruk() {
	var totalBelanja int

	for _, item := range *warung.StrukBelanja {
		fmt.Printf("Barang : %s\n", item.Nama)
		fmt.Printf("Harga : %d\n", item.Harga)
		fmt.Printf("Qty : %d\n", item.Qty)
		fmt.Printf("Jumlah : %d\n", item.Harga * item.Qty)
		totalBelanja += item.Harga * item.Qty
	}
	fmt.Printf("Total Belanja : %d\n", totalBelanja)
}