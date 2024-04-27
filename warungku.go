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
	KurangiQty(namaBarang string, qty int)
	CetakStruk()
}

func (warung *Warung) TambahBarang(barang *Barang) {
	foundIndex := -1
	for i, item := range *warung.StrukBelanja {
		if item.Nama == barang.Nama {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		*warung.StrukBelanja = append(*warung.StrukBelanja, *barang)
	} else {
		(*warung.StrukBelanja)[foundIndex].Qty = (*warung.StrukBelanja)[foundIndex].Qty + barang.Qty
	}
}

func (warung *Warung) KurangiQty(namaBarang string, qty int) {
	foundIndex := -1
	for i, item := range *warung.StrukBelanja {
		if item.Nama == namaBarang {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Printf("Barang %s tidak ditemukan\n", namaBarang)
	} else {
		if (*warung.StrukBelanja)[foundIndex].Qty - qty <= 0 {
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