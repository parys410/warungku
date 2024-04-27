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

type WarungInterface interface {
	TambahBarang(barang *Barang)
	KurangiQty(barang *Barang, qty int)
	CetakStruk()
}

func (warung *Warung) TambahBarang(barang *Barang) {
	*warung.StrukBelanja = append(*warung.StrukBelanja, *barang)
}

func (warung *Warung) KurangiQty(barang *Barang, qty int) {
	foundIndex := -1
	for i, item := range *warung.StrukBelanja {
		if item.Nama == barang.Nama {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Printf("Barang %s tidak ditemukan\n", barang.Nama)
	} else {
		if (*warung.StrukBelanja)[foundIndex].Qty - qty < 0 {
			// Hapus dari slice
			sliceBefore := (*warung.StrukBelanja)[:foundIndex]
			sliceAfter := (*warung.StrukBelanja)[foundIndex+1:]
			*warung.StrukBelanja = append(sliceBefore, sliceAfter...)
			
			fmt.Println("Barang Dihapus")
		} else {
			(*warung.StrukBelanja)[foundIndex].Qty = (*warung.StrukBelanja)[foundIndex].Qty - qty
			fmt.Println("Barang Dikurangi")
		}
	}
}

func (warung *Warung) CetakStruk() {
	var totalBelanja int

	for _, item := range *warung.StrukBelanja {
		fmt.Printf("-------------------------\n")
		fmt.Printf("Barang : %s\n", item.Nama)
		fmt.Printf("Harga : %d\n", item.Harga)
		fmt.Printf("Qty : %d\n", item.Qty)
		fmt.Printf("Jumlah : %d\n", item.Harga * item.Qty)
		totalBelanja += item.Harga * item.Qty
	}
	fmt.Printf("-------------------------\n")
	fmt.Printf("Total Belanja : %d\n", totalBelanja)
}